package console

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/binus-thesis-team/cacher"
	"github.com/binus-thesis-team/iam-service/auth"
	iamServiceClient "github.com/binus-thesis-team/iam-service/client"
	iam "github.com/binus-thesis-team/iam-service/pb/iam_service"
	iamGrpcUtils "github.com/binus-thesis-team/iam-service/utils/grpcutils"
	productServiceClient "github.com/binus-thesis-team/product-service/client"
	"github.com/binus-thesis-team/product-service/internal/config"
	"github.com/binus-thesis-team/product-service/internal/db"
	"github.com/binus-thesis-team/product-service/internal/delivery/httpsvc"
	"github.com/binus-thesis-team/product-service/internal/helper"
	"github.com/binus-thesis-team/product-service/internal/repository"
	"github.com/binus-thesis-team/product-service/internal/usecase"
	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/binus-thesis-team/product-service/pkg/utils"
	"github.com/binus-thesis-team/product-service/pkg/utils/grpcutils"
	productGrpcUtils "github.com/binus-thesis-team/product-service/pkg/utils/grpcutils"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var runServerCmd = &cobra.Command{
	Use:   "server",
	Short: "run server",
	Long:  `This subcommand start the server`,
	Run:   runServer,
}

func init() {
	RootCmd.AddCommand(runServerCmd)
}

func runServer(cmd *cobra.Command, args []string) {
	// Initiate all connection like db, redis, etc
	db.InitializePostgresConn()

	pgDB, err := db.PostgreSQL.DB()
	continueOrFatal(err)
	defer helper.WrapCloser(pgDB.Close)

	newIAMClient, err := getIAMServiceGRPCClient()
	continueOrFatal(err)

	newProductClient, err := getProductServiceGRPCClient()
	continueOrFatal(err)

	authenticationCacher := cacher.NewCacheManager()
	generalCacher := cacher.NewCacheManager()

	authRedisConn, err := db.NewRedigoRedisConnectionPool(config.RedisAuthCacheHost(), redisOpts)
	continueOrFatal(err)
	defer helper.WrapCloser(authRedisConn.Close)

	redisAuthLockConn, err := db.NewRedigoRedisConnectionPool(config.RedisAuthCacheLockHost(), redisOpts)
	continueOrFatal(err)

	authenticationCacher.SetConnectionPool(authRedisConn)
	authenticationCacher.SetLockConnectionPool(redisAuthLockConn)
	authenticationCacher.SetDisableCaching(config.DisableCaching())
	authenticationCacher.SetDefaultTTL(config.CacheTTL())

	generalCacher.SetDisableCaching(config.DisableCaching())

	if !config.DisableCaching() {
		redisConn, err := db.NewRedigoRedisConnectionPool(config.RedisCacheHost(), redisOpts)
		continueOrFatal(err)
		defer helper.WrapCloser(redisConn.Close)

		redisLockConn, err := db.NewRedigoRedisConnectionPool(config.RedisLockHost(), redisOpts)
		continueOrFatal(err)
		defer helper.WrapCloser(redisLockConn.Close)

		generalCacher.SetConnectionPool(redisConn)
		generalCacher.SetLockConnectionPool(redisLockConn)
		generalCacher.SetDefaultTTL(config.CacheTTL())
	}

	location, locErr := utils.SetTimeLocation("Asia/Jakarta")
	if locErr != nil {
		panic(locErr)
	}

	time.Local = location

	productRepository := repository.NewProductRepository(db.PostgreSQL, generalCacher)
	productUsecase := usecase.NewProductUsecase(productRepository, newProductClient)
	iamAuthAdapter := auth.NewIAMServiceAdapter(newIAMClient)
	authMiddleware := auth.NewAuthenticationMiddleware(iamAuthAdapter, authenticationCacher)

	httpServer := echo.New()
	httpServer.Pre(middleware.AddTrailingSlash())
	httpServer.Use(middleware.Logger())
	httpServer.Use(middleware.Recover())
	httpServer.Use(middleware.CORS())

	apiGroup := httpServer.Group("/api")
	httpsvc.RouteService(apiGroup, productUsecase, authMiddleware)

	sigCh := make(chan os.Signal, 1)
	errCh := make(chan error, 1)
	quitCh := make(chan bool, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		for {
			select {
			case <-sigCh:
				gracefulShutdown(nil, httpServer)
				quitCh <- true
			case e := <-errCh:
				log.Error(e)
				gracefulShutdown(nil, httpServer)
				quitCh <- true
			}
		}
	}()

	setupLogger()

	go func() {
		// Start HTTP server
		if err := httpServer.Start(fmt.Sprintf(":%s", config.HTTPPort())); err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	<-quitCh
	log.Info("exiting")
}

func gracefulShutdown(grpcSvr *grpc.Server, httpSvr *echo.Echo) {
	db.StopTickerCh <- true

	if grpcSvr != nil {
		grpcSvr.GracefulStop()
	}

	if httpSvr != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := httpSvr.Shutdown(ctx); err != nil {
			httpSvr.Logger.Fatal(err)
		}
	}
}

func clientInterceptor() grpc.UnaryClientInterceptor {
	return grpcutils.UnaryClientInterceptor(&grpcutils.GRPCUnaryInterceptorOptions{
		Timeout:           1*time.Second + 100*time.Millisecond,
		RetryCount:        0,
		UseCircuitBreaker: true,
	})
}

func getIAMServiceGRPCClient() (iam.IAMServiceClient, error) {
	grpcClient, err := iamServiceClient.NewGRPCClient("localhost:9000", newIAMGRPCPoolSetting(),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(clientInterceptor()))

	return grpcClient, err
}

func getProductServiceGRPCClient() (pb.ProductServiceClient, error) {
	grpcClient, err := productServiceClient.NewGRPCClient("localhost:9001", newProductGRPCPoolSetting(),
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithUnaryInterceptor(clientInterceptor()))

	return grpcClient, err
}

func newIAMGRPCPoolSetting() *iamGrpcUtils.GRPCConnectionPoolSetting {
	return &iamGrpcUtils.GRPCConnectionPoolSetting{
		MaxIdle:   100,
		MaxActive: 200,
	}
}

func newProductGRPCPoolSetting() *productGrpcUtils.GRPCConnectionPoolSetting {
	return &productGrpcUtils.GRPCConnectionPoolSetting{
		MaxIdle:   100,
		MaxActive: 200,
	}
}

func continueOrFatal(err error) {
	if err != nil {
		logrus.Fatal(err)
	}
}
