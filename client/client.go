package client

import (
	"context"
	"github.com/binus-thesis-team/product-service/internal/model"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ProductServiceClient defines the interface for interacting with the Product service.
type ProductServiceClient interface {
	FindByProductID(ctx context.Context, id int64) (*model.Product, error)
	SearchAllProducts(ctx context.Context, query string) (ids []int64, count int64, err error)
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
