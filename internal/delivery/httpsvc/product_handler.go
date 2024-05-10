package httpsvc

import (
	"github.com/binus-thesis-team/iam-service/utils"
	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/binus-thesis-team/product-service/internal/usecase"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (s *service) Create() echo.HandlerFunc {
	type request struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Stock       int64   `json:"stock"`
		Description string  `json:"description"`
		ImageUrl    string  `json:"image_url"`
	}

	return func(c echo.Context) error {
		ctx := c.Request().Context()

		req := request{}
		if err := c.Bind(&req); err != nil {
			logrus.Error(err)
			return ErrInvalidArgument
		}

		createdProduct, err := s.productUsecase.Create(ctx, model.GetUserFromCtx(ctx), model.CreateProductRequest{
			Name:        req.Name,
			Price:       req.Price,
			Stock:       req.Stock,
			Description: req.Description,
			ImageUrl:    req.ImageUrl,
		})
		switch err {
		case nil:
			break
		case usecase.ErrNotFound:
			return ErrNotFound
		case usecase.ErrDuplicateProduct:
			return ErrProductAlreadyExist
		case usecase.ErrPermissionDenied:
			return ErrPermissionDenied
		default:
			logrus.Error(err)
			return ErrInternal
		}

		return c.JSON(http.StatusCreated, setSuccessResponse(createdProduct))
	}
}

func (s *service) GetDetail() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		productID := utils.StringToInt64(c.Param("product_id"))

		product, err := s.productUsecase.FindByID(ctx, productID)
		if err != nil {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"product_id": productID,
			}).Error(err)

			return ErrInternal
		}

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"product_id": product.ID,
		}).Info("success delete product from db")

		return c.JSON(http.StatusOK, setSuccessResponse(product))
	}
}

func (s *service) GetList() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		pageStr := c.QueryParam("page")
		if pageStr == "" {
			pageStr = "1"
		}
		page, err := strconv.Atoi(pageStr)
		if err != nil {
			logrus.WithError(err).Error("failed to parse page")
			return c.JSON(http.StatusBadRequest, err)
		}

		limitStr := c.QueryParam("limit")
		if limitStr == "" {
			limitStr = "10"
		}
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			logrus.WithError(err).Error("failed to parse limit")
			return c.JSON(http.StatusBadRequest, err)
		}

		query := c.QueryParam("query")
		dir := c.QueryParam("dir")
		sort := c.QueryParam("sort")

		products, count, err := s.productUsecase.SearchByCriteria(ctx, model.GetUserFromCtx(ctx), model.ProductSearchCriteria{
			Query:   query,
			Page:    int64(page),
			Size:    int64(limit),
			SortBy:  sort,
			SortDir: dir,
		})
		if err != nil {
			logrus.WithError(err).Error("failed to get products")
			return c.JSON(http.StatusBadRequest, err)
		}

		logrus.WithFields(logrus.Fields{
			"page":  page,
			"limit": limit,
		}).Info("success get products")

		return c.JSON(http.StatusOK, toResourcePaginationResponse(page, limit, count, products))
	}
}

func (s *service) handleFindProductIDsByQuery() echo.HandlerFunc {
	type searchResponse struct {
		Count int64   `json:"count"`
		Ids   []int64 `json:"ids"`
	}

	return func(c echo.Context) error {
		ctx := c.Request().Context()

		query := c.QueryParam("query")

		productIDs, count, err := s.productUsecase.FindIDsByQuery(ctx, query)
		if err != nil {
			logrus.WithError(err).Error("failed to get products")
			return c.JSON(http.StatusBadRequest, err)
		}

		return c.JSON(http.StatusOK, setSuccessResponse(searchResponse{
			Count: count,
			Ids:   productIDs,
		}))
	}
}

func (s *service) Update() echo.HandlerFunc {
	type request struct {
		Name        string  `json:"name"`
		Price       float64 `json:"price"`
		Stock       int64   `json:"stock"`
		Description string  `json:"description"`
		ImageUrl    string  `json:"image_url"`
	}

	return func(c echo.Context) error {
		ctx := c.Request().Context()

		req := request{}
		if err := c.Bind(&req); err != nil {
			logrus.Error(err)
			return ErrInvalidArgument
		}
		productID := utils.StringToInt64(c.Param("product_id"))

		createdProduct, err := s.productUsecase.Update(ctx, model.GetUserFromCtx(ctx), model.UpdateProductRequest{
			ID:          productID,
			Name:        req.Name,
			Price:       req.Price,
			Stock:       req.Stock,
			Description: req.Description,
			ImageUrl:    req.ImageUrl,
		})
		switch err {
		case nil:
			break
		case usecase.ErrNotFound:
			return ErrNotFound
		case usecase.ErrDuplicateProduct:
			return ErrProductAlreadyExist
		case usecase.ErrPermissionDenied:
			return ErrPermissionDenied
		default:
			logrus.Error(err)
			return ErrInternal
		}

		return c.JSON(http.StatusCreated, setSuccessResponse(createdProduct))
	}
}

func (s *service) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		productID := utils.StringToInt64(c.Param("product_id"))

		if err := s.productUsecase.DeleteByProductID(ctx, model.GetUserFromCtx(ctx), productID); err != nil {
			logrus.WithContext(ctx).WithFields(logrus.Fields{
				"product_id": productID,
			}).Error(err)

			return ErrInternal
		}

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"product_id": productID,
		}).Info("success delete product from db")

		return c.JSON(http.StatusOK, setSuccessResponse(productID))
	}
}

func (s *service) UploadImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		file, err := c.FormFile("product_image")
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to get image file")
			return ErrInvalidArgument
		}

		workDir, _ := os.Getwd()
		path := filepath.Join(workDir, "", "/assets/products/", file.Filename)

		err = s.productUsecase.UploadImage(ctx, model.GetUserFromCtx(ctx), model.UploadImageProductRequest{
			ProductImage: file,
			Path:         path,
		})
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to upload image")
			return ErrInternal
		}

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"file_name": file.Filename,
			"file_path": path,
		}).Info("success upload image")

		return c.JSON(http.StatusOK, setSuccessResponse(model.UploadImageProductResponse{
			FileName: file.Filename,
			ImageUrl: path,
		}))
	}
}

func (s *service) RemoveImage() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		input := model.RemoveImageProductRequest{}

		err := c.Bind(&input)
		if err != nil {
			logrus.Error(err)
			return ErrInternal
		}

		err = s.productUsecase.RemoveImage(ctx, model.GetUserFromCtx(ctx), input)
		switch err {
		case nil:
			break
		case usecase.ErrNotFound:
			return ErrNotFound
		case usecase.ErrPermissionDenied:
			return ErrPermissionDenied
		default:
			logrus.Error(err)
			return ErrInternal
		}

		return c.JSON(http.StatusOK, setSuccessResponse(input.ImageUrl))
	}
}

func (s *service) UploadFile() echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()

		file, err := c.FormFile("product_file")
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to get product file")
			return ErrInvalidArgument
		}

		src, err := file.Open()
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to open product file")
			return err
		}
		defer src.Close()

		productFile, err := io.ReadAll(src)
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to open read product file")
			return err
		}

		err = s.productUsecase.UploadFile(ctx, model.GetUserFromCtx(ctx), model.UploadFileProductRequest{
			ProductFile: productFile,
		})
		if err != nil {
			logrus.WithContext(ctx).WithError(err).Error("failed to upload file")
			return ErrInternal
		}

		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"file_name": file.Filename,
		}).Info("success upload image")

		return c.JSON(http.StatusOK, setSuccessResponse(model.UploadImageProductResponse{
			FileName: file.Filename,
		}))
	}
}
