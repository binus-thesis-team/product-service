package httpsvc

import (
	"github.com/binus-thesis-team/iam-service/auth"
	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/labstack/echo/v4"
)

// service http service
type service struct {
	productUsecase model.ProductUsecase
	authMiddleware *auth.AuthenticationMiddleware
}

// RouteService ..
func RouteService(
	group *echo.Group,
	productUsecase model.ProductUsecase,
	authMiddleware *auth.AuthenticationMiddleware,
) {
	svc := &service{
		productUsecase: productUsecase,
		authMiddleware: authMiddleware,
	}

	svc.initInternalCommunicationRoutes(group.Group("/internal"))
	svc.initRoutes(group)
}

func (s *service) initRoutes(group *echo.Group) {
	productRoute := group.Group("/products", s.authMiddleware.MustAuthenticateAccessToken())
	{
		productRoute.POST("/", s.Create())
		productRoute.GET("/:product_id/", s.GetDetail())
		productRoute.GET("/", s.GetList())
		productRoute.PUT("/:product_id/", s.Update())
		productRoute.DELETE("/:product_id/", s.Delete())

		imageGroup := productRoute.Group("/images")
		{
			imageGroup.POST("/upload/", s.UploadImage())
			imageGroup.DELETE("/remove/", s.RemoveImage())
		}

		productRoute.POST("/file/upload/", s.UploadFile())
	}
}

func (s *service) initInternalCommunicationRoutes(group *echo.Group) {
	productRoute := group.Group("/products")
	{
		productRoute.GET("/:product_id/", s.GetDetail())
	}
}
