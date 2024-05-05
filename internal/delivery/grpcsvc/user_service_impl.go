package grpcsvc

import (
	"context"
	"time"

	"github.com/binus-thesis-team/product-service/internal/usecase"
	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/binus-thesis-team/product-service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// FindByProductIds :nodoc:
func (s *Service) FindByProductIDs(ctx context.Context, in []int64) (out *pb.FindByProductIDsResponse, err error) {
	products, err := s.productUsecase.FindByProductIDs(ctx, in)
	switch err {
	case nil:

		newProducts := make([]*pb.Product, 0, len(products))
		for _, v := range products {

			var createdAt, updatedAt, deletedAt string
			if v.CreatedAt != nil {
				createdAt = v.CreatedAt.Format(time.DateTime)
			}
			if v.UpdatedAt != nil {
				updatedAt = v.UpdatedAt.Format(time.DateTime)
			}
			if v.DeletedAt.Valid {
				deletedAt = v.DeletedAt.Time.Format(time.DateTime)
			}
			data := &pb.Product{
				Id:          v.ID,
				Name:        v.Name,
				Price:       v.Price,
				Stock:       v.Stock,
				Description: v.Description,
				ImageUrl:    v.ImageUrl,
				CreatedAt:   createdAt,
				UpdatedAt:   updatedAt,
				DeletedAt:   deletedAt,
			}

			newProducts = append(newProducts, data)
		}

		out.Product = newProducts
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
func (s *Service) FindByProductID(ctx context.Context, in int64) (out *pb.FindByProductIDResponse, err error) {
	product, err := s.productUsecase.FindByID(ctx, in)
	switch err {
	case nil:
		var createdAt, updatedAt, deletedAt string
		if product.CreatedAt != nil {
			createdAt = product.CreatedAt.Format(time.DateTime)
		}
		if product.UpdatedAt != nil {
			updatedAt = product.UpdatedAt.Format(time.DateTime)
		}
		if product.DeletedAt.Valid {
			deletedAt = product.DeletedAt.Time.Format(time.DateTime)
		}

		out.Product = &pb.Product{
			Id:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Stock:       product.Stock,
			Description: product.Description,
			ImageUrl:    product.ImageUrl,
			CreatedAt:   createdAt,
			UpdatedAt:   updatedAt,
			DeletedAt:   deletedAt,
		}
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
