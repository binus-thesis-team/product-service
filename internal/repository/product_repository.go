package repository

import (
	"context"
	"fmt"

	"github.com/binus-thesis-team/cacher"
	"github.com/binus-thesis-team/iam-service/utils"
	"github.com/binus-thesis-team/product-service/internal/config"
	"github.com/binus-thesis-team/product-service/internal/model"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type productRepository struct {
	db           *gorm.DB
	cacheManager cacher.CacheManager
}

func NewProductRepository(db *gorm.DB, cacheManager cacher.CacheManager) model.ProductRepository {
	return &productRepository{
		db:           db,
		cacheManager: cacheManager,
	}
}

func (u *productRepository) Create(ctx context.Context, requesterID int64, product *model.Product) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         utils.DumpIncomingContext(ctx),
		"requesterID": requesterID,
		"product":     utils.Dump(product),
	})

	err := u.db.WithContext(ctx).Create(product).Error
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := u.cacheManager.DeleteByKeys([]string{
		u.newCacheKeyByID(product.ID),
	}); err != nil {
		logger.Error(err)
	}

	return nil
}

func (u *productRepository) FindByID(ctx context.Context, id int64) (*model.Product, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	cacheKey := u.newCacheKeyByID(id)
	if !config.DisableCaching() {
		reply, mu, err := findFromCacheByKey[*model.Product](u.cacheManager, cacheKey)
		defer cacher.SafeUnlock(mu)
		if err != nil {
			logger.Error(err)
			return nil, err
		}

		if mu == nil {
			return reply, nil
		}
	}

	product := &model.Product{}
	err := u.db.WithContext(ctx).Take(product, "id = ?", id).Error
	switch err {
	case nil:
	case gorm.ErrRecordNotFound:
		storeNil(u.cacheManager, cacheKey)
		return nil, nil
	default:
		logger.Error(err)
		return nil, err
	}

	err = u.cacheManager.StoreWithoutBlocking(cacher.NewItem(cacheKey, utils.Dump(product)))
	if err != nil {
		logger.Error(err)
	}

	return product, nil
}

func (u *productRepository) FindByProductIDs(ctx context.Context, productIDs []int64) ([]*model.Product, error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
	})

	var products []*model.Product

	// TODO: Implement caching ?
	// cacheKey := u.newCacheKeyByProductID(productIDs)
	// if !config.DisableCaching() {
	// 	_, mu, err := findFromCacheByKey[int64](u.cacheManager, cacheKey)
	// 	defer cacher.SafeUnlock(mu)
	// 	if err != nil {
	// 		logger.Error(err)
	// 		return nil, err
	// 	}
	// }

	var scopes []func(*gorm.DB) *gorm.DB
	scopes = append(scopes, u.scopeByProductIDs(productIDs))

	err := u.db.Model(model.Product{}).
				Scopes(scopes...).
				Find(&products).Error
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return products, nil
}

func (u *productRepository) UpdateByID(ctx context.Context, requesterID int64, product *model.Product) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":         utils.DumpIncomingContext(ctx),
		"requesterID": requesterID,
		"product":     utils.Dump(product),
	})

	err := u.db.WithContext(ctx).Updates(product).Debug().Error
	if err != nil {
		logger.Error(err)
		return err
	}

	if err := u.cacheManager.DeleteByKeys([]string{
		u.newCacheKeyByID(product.ID),
	}); err != nil {
		logger.Error(err)
	}

	return nil
}

func (u *productRepository) DeleteByID(ctx context.Context, id int64) error {
	logger := logrus.WithFields(logrus.Fields{
		"ctx": utils.DumpIncomingContext(ctx),
		"id":  id,
	})

	space := &model.Product{ID: id}

	if err := u.db.WithContext(ctx).Delete(space).Error; err != nil {
		logger.Error(err)
		return err
	}

	return nil
}

func (u *productRepository) SearchByPage(ctx context.Context, searchCriteria model.ProductSearchCriteria) (ids []int64, count int64, err error) {
	logger := logrus.WithFields(logrus.Fields{
		"ctx":            utils.DumpIncomingContext(ctx),
		"searchCriteria": utils.Dump(searchCriteria),
	})

	var scopes []func(*gorm.DB) *gorm.DB

	if len(searchCriteria.ProductIDs) > 0 {
		scopes = append(scopes, u.scopeByProductIDs(searchCriteria.ProductIDs))
	}

	count, err = u.countAll(ctx, searchCriteria)
	if err != nil {
		logger.Error(err)
		return nil, 0, err
	}

	if count <= 0 {
		return nil, 0, nil
	}

	ids, err = u.findAllIDsByCriteria(ctx, searchCriteria)
	switch err {
	case nil:
		return ids, count, nil
	case gorm.ErrRecordNotFound:
		return nil, 0, nil
	default:
		logger.Error(err)
		return nil, 0, err
	}
}

func (u *productRepository) findAllIDsByCriteria(ctx context.Context, criteria model.ProductSearchCriteria) ([]int64, error) {
	var scopes []func(*gorm.DB) *gorm.DB
	scopes = append(scopes, scopeByPageAndLimit(criteria.Page, criteria.Size))

	if len(criteria.ProductIDs) > 0 {
		scopes = append(scopes, u.scopeByProductIDs(criteria.ProductIDs))
	}

	var ids []int64
	err := u.db.WithContext(ctx).
		Model(model.Product{}).
		Scopes(scopes...).
		Order(fmt.Sprintf("%s %s", criteria.SortBy, criteria.SortDir)).
		Pluck("id", &ids).Debug().Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":      utils.DumpIncomingContext(ctx),
			"criteria": utils.Dump(criteria),
		}).Error(err)
		return nil, err
	}

	return ids, nil
}

func (u *productRepository) scopeByProductIDs(productIDs []int64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("product_id IN (?)", productIDs)
	}
}

func (u *productRepository) countAll(ctx context.Context, criteria model.ProductSearchCriteria) (int64, error) {
	var scopes []func(*gorm.DB) *gorm.DB

	if len(criteria.ProductIDs) > 0 {
		scopes = append(scopes, u.scopeByProductIDs(criteria.ProductIDs))
	}

	var count int64
	err := u.db.WithContext(ctx).Model(model.Product{}).
		Scopes(scopes...).
		Count(&count).
		Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"ctx":      utils.DumpIncomingContext(ctx),
			"criteria": utils.Dump(criteria),
		}).Error(err)
		return 0, err
	}

	return count, nil
}

func (u *productRepository) newCacheKeyByProductID(productID int64) string {
	return fmt.Sprintf("cache:id:product:%d", productID)
}

func (u *productRepository) newCacheKeyByID(id int64) string {
	return fmt.Sprintf("cache:object:product:id:%d", id)
}
