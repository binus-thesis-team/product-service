package client

import (
	"context"

	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/binus-thesis-team/product-service/pkg/utils/grpcutils"
	grpcpool "github.com/processout/grpc-go-pool"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type grpcClient struct {
	Conn *grpcpool.Pool
	Opts []grpc.CallOption
}

func NewGRPCClient(target string, poolSetting *grpcutils.GRPCConnectionPoolSetting, dialOptions ...grpc.DialOption) (pb.ProductServiceClient, error) {
	c, err := grpcutils.NewGRPCConnection(target, poolSetting, dialOptions...)
	if err != nil {
		logrus.Fatal(err)
		return nil, err
	}

	return &grpcClient{Conn: c}, nil
}

func (c *grpcClient) SetCallOption(opts ...grpc.CallOption) {
	c.Opts = opts
}

func (c *grpcClient) FindAllProductsByIDs(ctx context.Context, in *pb.FindByIDsRequest, opts ...grpc.CallOption) (out *pb.Products, err error) {
	conn, err := c.Conn.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err = cli.FindAllProductsByIDs(ctx, in, opts...)
	return
}

func (c *grpcClient) FindByProductID(ctx context.Context, in *pb.FindByIDRequest, opts ...grpc.CallOption) (out *pb.Product, err error) {
	conn, err := c.Conn.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err = cli.FindByProductID(ctx, in, opts...)
	return
}

func (c *grpcClient) SearchAllProducts(ctx context.Context, in *pb.ProductSearchRequest, opts ...grpc.CallOption) (out *pb.SearchResponse, err error) {
	conn, err := c.Conn.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err = cli.SearchAllProducts(ctx, in, opts...)
	return
}

func (c *grpcClient) UploadProducts(ctx context.Context, in *pb.UploadProductsRequest, opts ...grpc.CallOption) (out *pb.UploadProductsResponse, err error) {
	conn, err := c.Conn.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err = cli.UploadProducts(ctx, in, opts...)
	return
}
