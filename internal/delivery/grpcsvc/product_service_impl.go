package grpcsvc

import (
	"context"
	"github.com/binus-thesis-team/product-service/internal/config"
	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/binus-thesis-team/product-service/internal/usecase"
	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"github.com/binus-thesis-team/product-service/pkg/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// FindAllProductsByIDs :nodoc:
func (s *Service) FindAllProductsByIDs(ctx context.Context, in *pb.FindByIDsRequest) (out *pb.Products, err error) {
	products, err := s.productUsecase.FindByProductIDs(ctx, in.GetIds())
	switch err {
	case nil:
		protoProducts := pb.Products{}

		for _, item := range products {
			protoProducts.Products = append(protoProducts.Products, item.ToProto())
		}

		return &protoProducts, nil
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
func (s *Service) FindByProductID(ctx context.Context, in *pb.FindByIDRequest) (out *pb.Product, err error) {
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

func (s *Service) SearchAllProducts(ctx context.Context, req *pb.ProductSearchRequest) (out *pb.SearchResponse, err error) {
	size := utils.Int64WithLimit(req.GetSize(), config.MaxSizePerRequest())
	param := model.ProductSearchCriteria{
		Query: strings.ToLower(req.GetQuery()),
		Page:  req.GetPage(),
		Size:  size,
	}

	ids, count, err := s.productUsecase.SearchByPage(ctx, param)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.SearchResponse{
		Ids:   ids,
		Count: count,
	}, nil
}
