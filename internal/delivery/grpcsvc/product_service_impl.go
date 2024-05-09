package grpcsvc

import (
	"context"
	"sync"

	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/binus-thesis-team/product-service/internal/usecase"
	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/binus-thesis-team/product-service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FindByProductIds :nodoc:
func (s *Service) FindByProductIDs(ctx context.Context, in *pb.FindByProductIDsRequest) (out *pb.FindByProductIDsResponse, err error) {
	products, err := s.productUsecase.FindByProductIDs(ctx, in.GetIds())
	switch err {
	case nil:

		out.Products = make([]*pb.Product, 0, len(products))
		var wg sync.WaitGroup
		for index, v := range products {
			wg.Add(1)
			go func(idx int, product *model.Product) {
				defer wg.Done()
				out.Products[idx] = product.ToProto()
			}(index, v)
		}
		wg.Wait()

		return out, nil
	case usecase.ErrNotFound:
		return nil, status.Error(codes.NotFound, "not found")
	default:
		logrus.WithFields(logrus.Fields{
			"ctx": utils.DumpIncomingContext(ctx),
			"req": utils.Dump(in),
		}).Error(err)
		return nil, status.Error(codes.Internal, "something wrong")
	}
}

// FindByProductId :nodoc:
func (s *Service) FindByProductID(ctx context.Context, in *pb.FindByProductIDRequest) (out *pb.Product, err error) {
	product, err := s.productUsecase.FindByID(ctx, in.GetId())
	switch err {
	case nil:
		out = product.ToProto()

		return out, nil
	case usecase.ErrNotFound:
		return nil, status.Error(codes.NotFound, "not found")
	default:
		logrus.WithFields(logrus.Fields{
			"ctx": utils.DumpIncomingContext(ctx),
			"req": utils.Dump(in),
		}).Error(err)
		return nil, status.Error(codes.Internal, "something wrong")
	}
}
