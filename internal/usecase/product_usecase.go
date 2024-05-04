package usecase

import (
	"context"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/binus-thesis-team/iam-service/rbac"
	"github.com/binus-thesis-team/iam-service/utils"
	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/sirupsen/logrus"
)

type productUsecase struct {
	productRepository model.ProductRepository
	//productServiceClient          productClient.ProductServiceClient
}

func NewProductUsecase(productRepository model.ProductRepository) model.ProductUsecase {
	return &productUsecase{
		productRepository: productRepository,
		//productServiceClient:          productServiceClient,
	}
}

// Create inserts product to db, return productID and error
func (u *productUsecase) Create(ctx context.Context, user model.SessionUser, input model.CreateProductRequest) (product *model.Product, err error) {
	if !user.HasAccess(rbac.ResourceProduct, rbac.ActionCreateAny) {
		return nil, ErrPermissionDenied
	}

	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"input": utils.Dump(input),
	})

	if err := input.Validate(); err != nil {
		logger.Error(err)
		return nil, err
	}

	if err := input.ValidateDTOCreateProductRequest(); err != nil {
		logger.Error(err)
		return nil, err
	}

	product = &model.Product{
		Name:        input.Name,
		Price:       input.Price,
		Stock:       input.Stock,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
	}

	if err := u.productRepository.Create(ctx, user.GetUserID(), product); err != nil {
		logger.Error(err)
		return nil, err
	}

	return u.FindByID(ctx, product.ID)
}

func (u *productUsecase) FindByID(ctx context.Context, id int64) (product *model.Product, err error) {
	product, err = u.productRepository.FindByID(ctx, id)
	if err != nil {
		logrus.WithField("id", id).Error(err)
		return nil, err
	}

	if product == nil {
		return nil, ErrNotFound
	}

	return product, nil
}

func (u *productUsecase) FindByProductIDs(ctx context.Context, productIDs []int64) (products []*model.Product, err error) {
	products, err = u.productRepository.FindByProductIDs(ctx, productIDs)
	if err != nil {
		logrus.WithField("productID", productIDs).Error(err)
		return nil, err
	}

	if products == nil {
		return nil, ErrNotFound
	}

	return products, nil
}

func (u *productUsecase) Update(ctx context.Context, user model.SessionUser, input model.UpdateProductRequest) (product *model.Product, err error) {
	if !user.HasAccess(rbac.ResourceProduct, rbac.ActionCreateAny) {
		return nil, ErrPermissionDenied
	}

	logger := logrus.WithFields(logrus.Fields{
		"ctx":   utils.DumpIncomingContext(ctx),
		"input": utils.Dump(input),
	})

	if err := input.Validate(); err != nil {
		logger.Error(err)
		return nil, err
	}

	if err := input.ValidateDTOUpdateProductRequest(); err != nil {
		logger.Error(err)
		return nil, err
	}

	product, err = u.productRepository.FindByID(ctx, input.ID)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	product = &model.Product{
		ID:          product.ID,
		Name:        input.Name,
		Price:       input.Price,
		Stock:       input.Stock,
		Description: input.Description,
		ImageUrl:    input.ImageUrl,
	}

	if err := u.productRepository.UpdateByID(ctx, user.GetUserID(), product); err != nil {
		logger.Error(err)
		return nil, err
	}

	return product, nil
}

func (u *productUsecase) DeleteByProductID(ctx context.Context, user model.SessionUser, productID int64) (err error) {
	if !user.HasAccess(rbac.ResourceProduct, rbac.ActionDeleteAny) {
		return ErrPermissionDenied
	}

	logger := logrus.WithFields(logrus.Fields{
		"ctx":       utils.DumpIncomingContext(ctx),
		"user":      utils.Dump(user),
		"productID": productID,
	})

	product, err := u.FindByID(ctx, productID)
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := u.productRepository.DeleteByID(ctx, product.ID); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (u *productUsecase) Search(ctx context.Context, user model.SessionUser, searchCriteria model.ProductSearchCriteria) (products []*model.Product, count int64, err error) {
	if !user.HasAccess(rbac.ResourceProduct, rbac.ActionViewAny) {
		err = ErrPermissionDenied
		return
	}

	logger := logrus.WithFields(logrus.Fields{
		"ctx":            utils.DumpIncomingContext(ctx),
		"currentUser":    utils.Dump(user),
		"searchCriteria": utils.Dump(searchCriteria),
	})

	// TODO: filtering searchCriteria.Query to Product Service with GRPC/REST
	// get the productIDs to filtering on products

	// example
	// productIDs := []int64{1, 2, 3}

	// searchCriteria.ProductIDs = productIDs
	ids, count, err := u.productRepository.SearchByPage(ctx, searchCriteria)
	if err != nil {
		return nil, 0, err
	}

	fmt.Println("ids", ids)
	fmt.Println("count", count)

	if len(ids) == 0 || count == 0 {
		return nil, 0, nil
	}

	products = u.FindAllByIDs(ctx, ids)
	if len(products) <= 0 {
		logger.Error(ErrNotFound)
		return
	}

	return products, count, nil
}

func (u *productUsecase) FindAllByIDs(ctx context.Context, ids []int64) (products []*model.Product) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"ids": ids,
	})

	fmt.Println("masuk")
	var wg sync.WaitGroup
	c := make(chan *model.Product, len(ids))
	for _, id := range ids {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()

			product, err := u.FindByID(ctx, id)
			if err != nil {
				logger.Error(err)
				return
			}
			c <- product
		}(id)
	}
	wg.Wait()
	close(c)

	if len(c) <= 0 {
		return
	}

	rs := map[int64]*model.Product{}
	for product := range c {
		if product != nil {
			rs[product.ID] = product
		}
	}

	for _, id := range ids {
		if user, ok := rs[id]; ok {
			products = append(products, user)
		}
	}

	fmt.Println("products", products)

	return
}

func (ps *productUsecase) UploadImage(ctx context.Context, user model.SessionUser, input model.UploadImageProductRequest) error {
	if !user.HasAccess(rbac.ResourceUser, rbac.ActionCreateAny) {
		return ErrPermissionDenied
	}

	logger := logrus.WithFields(logrus.Fields{
		"ctx":         utils.DumpIncomingContext(ctx),
		"currentUser": utils.Dump(user),
	})

	if err := input.ValidateDTOUploadImageProductRequest(); err != nil {
		logger.Error(err)
		return err
	}

	src, err := input.ProductImage.Open()
	if err != nil {
		logger.Error(err)
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create(input.Path)
	if err != nil {
		logger.Error(err)
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (ps *productUsecase) RemoveImage(ctx context.Context, user model.SessionUser, input model.RemoveImageProductRequest) error {
	if !user.HasAccess(rbac.ResourceUser, rbac.ActionCreateAny) {
		return ErrPermissionDenied
	}

	logger := logrus.WithFields(logrus.Fields{
		"ctx":         utils.DumpIncomingContext(ctx),
		"currentUser": utils.Dump(user),
	})

	if err := input.Validate(); err != nil {
		logger.Error(err)
		return err
	}

	// Check if file exists
	if _, err := os.Stat(input.ImageUrl); os.IsNotExist(err) {
		logger.Error(err)
		return err
	}

	// Delete the file
	if err := os.Remove(input.ImageUrl); err != nil {
		logger.Error(err)
		return err
	}

	return nil
}