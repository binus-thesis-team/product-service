package http

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/irvankadhafi/go-boilerplate/internal/dto"
	"github.com/irvankadhafi/go-boilerplate/internal/service"
	httpresponse "github.com/irvankadhafi/go-boilerplate/pkg/http_response"
	"github.com/sirupsen/logrus"
)

type ProductController struct {
	productService service.Product
}

func NewProductController(productService service.Product) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// Create inserts product to db
func (pc *ProductController) Create(ctx *gin.Context) {
	dtoProduct := &dto.CreateProductRequest{}

	if err := ctx.ShouldBindJSON(&dtoProduct); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	productID, err := pc.productService.Create(ctx, dtoProduct)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"product_id": productID,
	}).Info("success inserts product to db")

	httpresponse.Write(ctx, http.StatusOK, productID, nil)
}

// Get detail product to db
func (pc *ProductController) GetDetail(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	product, err := pc.productService.GetDetail(ctx, id)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"product": product,
	}).Info("success get detail product to db")

	httpresponse.Write(ctx, http.StatusOK, product, nil)
}

// Delete will delete product from db by productID
func (pc *ProductController) Delete(ctx *gin.Context) {
	paramID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"id": ctx.Param("id"),
		}).Error(fmt.Errorf("fail convert param id string to int: %v", err))

		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	productID := int64(paramID)
	if err := pc.productService.Delete(ctx, productID); err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"product_id": productID,
		}).Error(err)

		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"product_id": productID,
	}).Info("success delete product from db")

	httpresponse.Write(ctx, http.StatusOK, productID, nil)
}

// Update will update product by id for every field that is not default value
func (pc *ProductController) Update(ctx *gin.Context) {
	dtoProduct := &dto.UpdateProductRequest{}
	if err := ctx.ShouldBindJSON(dtoProduct); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	paramID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logrus.WithContext(ctx).WithFields(logrus.Fields{
			"id": ctx.Param("id"),
		}).Error(fmt.Errorf("fail convert param id string to int: %v", err))

		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	dtoProduct.ID = int64(paramID)

	productID, err := pc.productService.Update(ctx, dtoProduct)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"product_id": productID,
	}).Info("success update product")

	httpresponse.Write(ctx, http.StatusOK, productID, nil)
}

func (pc *ProductController) GetList(ctx *gin.Context) {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	query := ctx.Query("query")
	dir := ctx.Query("dir")
	sort := ctx.Query("sort")

	products, err := pc.productService.GetList(ctx, page, limit, query, sort, dir)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"page":  page,
		"limit": limit,
	}).Info("success get products")

	httpresponse.Write(ctx, http.StatusOK, products, nil)
}

func (pc *ProductController) UploadImageProduct(ctx *gin.Context) {
	dtoUploadImage := &dto.UploadImageProductRequest{}

	if err := ctx.ShouldBind(&dtoUploadImage); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	workDir, _ := os.Getwd()
	filepath := filepath.Join(workDir, "", "/assets/products/", dtoUploadImage.ProductImage.Filename)

	dtoUploadImage.Path = filepath

	if err := pc.productService.UploadImage(ctx, dtoUploadImage); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	res := &dto.UploadImageProductResponse{
		FileName: dtoUploadImage.ProductImage.Filename,
		ImageUrl: dtoUploadImage.Path,
	}

	logrus.WithContext(ctx).WithField("Upload Product", res).Info("success upload image products")

	httpresponse.Write(ctx, http.StatusOK, res, nil)
}

func (pc *ProductController) RemoveImageProduct(ctx *gin.Context) {
	dtoRemoveImage := &dto.RemoveImageProductRequest{}

	if err := ctx.ShouldBindJSON(&dtoRemoveImage); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	if err := pc.productService.RemoveImage(ctx, dtoRemoveImage); err != nil {
		logrus.WithContext(ctx).Error(err)
		httpresponse.Write(ctx, http.StatusBadRequest, nil, err)
		return
	}

	filename := strings.Split(dtoRemoveImage.ImageUrl, "assets/products/")[1]
	logrus.WithContext(ctx).WithFields(logrus.Fields{
		"file_name": filename,
	}).Info("success remove image product")

	httpresponse.Write(ctx, http.StatusOK, filename, nil)
}
