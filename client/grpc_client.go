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

func (c *grpcClient) FindByProductIDs(ctx context.Context, in *pb.FindByProductIDsRequest, opts ...grpc.CallOption) (out *pb.FindByProductIDsResponse, err error) {
	conn, err := c.Conn.Get(ctx)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = conn.Close()
	}()
	cli := pb.NewProductServiceClient(conn.ClientConn)
	out, err = cli.FindByProductIDs(ctx, in, opts...)
	return
}

func (c *grpcClient) FindByProductID(ctx context.Context, in *pb.FindByProductIDRequest, opts ...grpc.CallOption) (out *pb.FindByProductIDResponse, err error) {
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
