package model

import (
	"context"
	"errors"
	"mime/multipart"
	"time"

	pb "github.com/binus-thesis-team/product-service/pb/product_service"
	"google.golang.org/protobuf/types/known/timestamppb"

	"gorm.io/gorm"
)

type ProductUsecase interface {
	Create(ctx context.Context, user SessionUser, input CreateProductRequest) (product *Product, err error)
	FindByID(ctx context.Context, id int64) (product *Product, err error)
	FindByProductIDs(ctx context.Context, productIDs []int64) (product []*Product, err error)
	Update(ctx context.Context, user SessionUser, input UpdateProductRequest) (product *Product, err error)
	DeleteByProductID(ctx context.Context, user SessionUser, productID int64) (err error)
	SearchByPage(ctx context.Context, searchCriteria ProductSearchCriteria) (ids []int64, count int64, err error)
	SearchByCriteria(ctx context.Context, user SessionUser, searchCriteria ProductSearchCriteria) (products []*Product, count int64, err error)
	FindIDsByQuery(ctx context.Context, query string) (ids []int64, count int64, err error)
	FindAllByIDs(ctx context.Context, ids []int64) (products []*Product)
	UploadImage(ctx context.Context, user SessionUser, input UploadImageProductRequest) error
	RemoveImage(ctx context.Context, user SessionUser, input RemoveImageProductRequest) error
	UploadFile(ctx context.Context, user SessionUser, input UploadFileProductRequest) error
	UploadFileWithoutSession(ctx context.Context, input UploadFileProductRequest) error
}

type ProductRepository interface {
	Create(ctx context.Context, requesterID int64, product *Product) error
	FindByID(ctx context.Context, id int64) (*Product, error)
	UpdateByID(ctx context.Context, requesterID int64, product *Product) (err error)
	DeleteByID(ctx context.Context, id int64) error
	SearchByPage(ctx context.Context, searchCriteria ProductSearchCriteria) (ids []int64, count int64, err error)
	FindAllByQuery(ctx context.Context, query string, size, cursorAfter int64) (ids []int64, err error)
}

type Product struct {
	ID          int64          `json:"id,omitempty" gorm:"<-:create; primary_key;AUTO_INCREMENT"`
	Name        string         `json:"name,omitempty"`
	Price       float64        `json:"price,omitempty"`
	Stock       int64          `json:"stock,omitempty"`
	Description string         `json:"description,omitempty"`
	ImageUrl    string         `json:"image_url,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty" gorm:"->;<-:create"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (p *Product) ToProto() *pb.Product {
	product := &pb.Product{
		Id:          p.ID,
		Name:        p.Name,
		Price:       p.Price,
		Stock:       p.Stock,
		Description: p.Description,
		ImageUrl:    p.ImageUrl,
	}

	if product.CreatedAt.IsValid() {
		product.CreatedAt = timestamppb.New(*p.CreatedAt)
	}
	if product.UpdatedAt.IsValid() {
		product.UpdatedAt = timestamppb.New(*p.CreatedAt)
	}
	if product.DeletedAt.IsValid() {
		product.DeletedAt = timestamppb.New(p.DeletedAt.Time)
	}

	return product
}

func NewProductFromProto(p *pb.Product) *Product {
	if p == nil {
		return nil
	}

	product := &Product{
		ID:          p.GetId(),
		Name:        p.GetName(),
		Price:       p.GetPrice(),
		Stock:       p.GetStock(),
		Description: p.GetDescription(),
		ImageUrl:    p.GetImageUrl(),
	}

	createdAt := p.GetCreatedAt().AsTime()
	product.CreatedAt = &createdAt

	updatedAt := p.GetUpdatedAt().AsTime()
	product.UpdatedAt = &updatedAt

	deletedAt := p.GetDeletedAt().AsTime()
	_ = product.DeletedAt.Scan(deletedAt)

	return product
}

type CreateProductRequest struct {
	Name        string  `json:"name,omitempty" binding:"required"`
	Price       float64 `json:"price,omitempty" binding:"required"`
	Stock       int64   `json:"stock,omitempty" binding:"required"`
	Description string  `json:"description,omitempty" binding:"required"`
	ImageUrl    string  `json:"image_url,omitempty" binding:"required"`
}

func (c *CreateProductRequest) Validate() error {
	return validate.Struct(c)
}

func (c *CreateProductRequest) ValidateDTOCreateProductRequest() error {
	if c.Name == "" {
		return errors.New("Name is required")
	}

	if c.Price <= 0 {
		return errors.New("Price must be greater than 0")
	}

	if c.Stock <= 0 {
		return errors.New("Stock must be greater than 0")
	}

	if c.Description == "" {
		return errors.New("Description is required")
	}

	if c.ImageUrl == "" {
		return errors.New("Image URL is required")
	}

	return nil
}

type UpdateProductRequest struct {
	ID          int64   `json:"-"`
	Name        string  `json:"name,omitempty" binding:"required"`
	Price       float64 `json:"price,omitempty" binding:"required"`
	Stock       int64   `json:"stock,omitempty" binding:"required"`
	Description string  `json:"description,omitempty" binding:"required"`
	ImageUrl    string  `json:"image_url,omitempty" binding:"required"`
}

func (c *UpdateProductRequest) Validate() error {
	return validate.Struct(c)
}

func (c *UpdateProductRequest) ValidateDTOUpdateProductRequest() error {
	if c.ID <= 0 {
		return errors.New("ID is required")
	}

	if c.Name == "" {
		return errors.New("Name is required")
	}

	if c.Price <= 0 {
		return errors.New("Price must be greater than 0")
	}

	if c.Stock <= 0 {
		return errors.New("Stock must be greater than 0")
	}

	if c.Description == "" {
		return errors.New("Description is required")
	}

	if c.ImageUrl == "" {
		return errors.New("Image URL is required")
	}

	return nil
}

// ProductSearchCriteria :nodoc:
type ProductSearchCriteria struct {
	Query   string `json:"query"`
	Page    int64  `json:"page"`
	Size    int64  `json:"size"`
	SortBy  string `json:"sort_by"`
	SortDir string `json:"sort_dir"`
}

// SetDefaultValue will set default value for page and size if zero
func (c *ProductSearchCriteria) SetDefaultValue() {
	if c.Page == 0 {
		c.Page = 1
	}
	if c.Size == 0 {
		c.Size = 10
	}
	if c.SortBy == "" {
		c.SortBy = "created_at"
	}
	if c.SortDir == "" {
		c.SortDir = "desc"
	}
}

type UploadImageProductRequest struct {
	ProductImage *multipart.FileHeader `form:"product_image" binding:"required"`
	Path         string
}

func (ps *UploadImageProductRequest) ValidateDTOUploadImageProductRequest() error {
	const maxFileSize = 5 << 20 // 5 MB in bytes
	if ps.ProductImage.Size > maxFileSize {
		return errors.New("The file size exceeds the maximum limit of 5 MB")
	}

	return nil
}

type UploadImageProductResponse struct {
	FileName string `json:"file_name"`
	ImageUrl string `json:"image_url"`
}

type RemoveImageProductRequest struct {
	ImageUrl string `json:"image_url" binding:"required"`
}

func (r *RemoveImageProductRequest) Validate() error {
	return validate.Struct(r)
}

type UploadFileProductRequest struct {
	ProductFile []byte `form:"product_file" binding:"required"`
}
