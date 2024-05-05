package grpcsvc

import (
	"github.com/binus-thesis-team/cacher"
	"github.com/binus-thesis-team/product-service/internal/model"
	pb "github.com/binus-thesis-team/product-service/pb/product_service"
)

// Service :nodoc:
type Service struct {
	pb.UnimplementedProductServiceServer
	cacheManager   cacher.CacheManager
	productUsecase model.ProductUsecase
}

// NewService :nodoc:
func NewService() *Service {
	return new(Service)
}

// RegisterCacheManager :nodoc:
func (s *Service) RegisterCacheManager(k cacher.CacheManager) {
	s.cacheManager = k
}

// RegisterProductUsecase :nodoc:
func (s *Service) RegisterProductUsecase(pc model.ProductUsecase) {
	s.productUsecase = pc
}
