package client

import (
	"context"

	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductServiceClient defines the interface for interacting with the Product service.
type ProductServiceClient interface {
	FindByProductIDs(ctx context.Context, req *pb.FindByProductIDsRequest, co ...grpc.CallOption) (*pb.FindByProductIDsResponse, error)
	FindByProductID(ctx context.Context, req *pb.FindByProductIDRequest, co ...grpc.CallOption) (*pb.FindByProductIDResponse, error)
}

// convertErrorGRPCToErrorGeneral mapping error from GRPC Error to general Error
func convertErrorGRPCToErrorGeneral(err error) error {
	if err == nil {
		return nil
	}

	grpcError, ok := status.FromError(err)
	if !ok {
		return err
	}

	switch grpcError.Code() {
	case codes.NotFound:
		return ErrNotFound
	default:
		logrus.Error(err)
		return ErrInternalServerError
	}
}
