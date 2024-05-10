package client

import (
	"context"
	"github.com/binus-thesis-team/product-service/internal/model"

	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/binus-thesis-team/product-service/pkg/utils/grpcutils"
	grpcpool "github.com/processout/grpc-go-pool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcClient struct {
	pool *grpcpool.Pool
	opts []grpc.CallOption
}

func NewGRPCClient(target string, poolSetting *grpcutils.GRPCConnectionPoolSetting, dialOptions ...grpc.DialOption) (ProductServiceClient, error) {
	pool, err := grpcutils.NewGRPCConnection(target, poolSetting, dialOptions...)
	if err != nil {
		logrus.Errorf("Error creating gRPC connection: %v", err)
		return nil, err
	}

	return &grpcClient{pool: pool}, nil
}

func (g *grpcClient) SetCallOption(opts ...grpc.CallOption) {
	g.opts = opts
}

func (g *grpcClient) FindByProductID(ctx context.Context, id int64) (*model.Product, error) {
	conn, err := g.pool.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err := cli.FindByProductID(ctx, &pb.FindByIDRequest{Id: id}, g.opts...)
	if err != nil {
		return nil, err
	}
	return model.NewProductFromProto(out), nil
}

func (g *grpcClient) SearchAllProducts(ctx context.Context, query string) (ids []int64, count int64, err error) {
	conn, err := g.pool.Get(ctx)
	if err != nil {
		return nil, 0, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err := cli.SearchAllProducts(ctx, &pb.ProductSearchRequest{
		Query: query,
	})
	if err != nil {
		return nil, 0, err
	}

	return out.GetIds(), out.GetCount(), nil
}
