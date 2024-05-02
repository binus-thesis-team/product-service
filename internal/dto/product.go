package dto

import "mime/multipart"

type CreateProductRequest struct {
	Name        string  `json:"name,omitempty" binding:"required"`
	Price       float64 `json:"price,omitempty" binding:"required"`
	Stock       int64   `json:"stock,omitempty" binding:"required"`
	Description string  `json:"description,omitempty" binding:"required"`
	ImageUrl    string  `json:"image_url,omitempty" binding:"required"`
}

type GetProductListRequest struct {
	Dir       string
	Limit     int64
	Offset    int64
	Page      int64
	ProcessId string
	Query     string
	Sort      string
}

type Product struct {
	Id          int64
	Name        string
	Price       float64
	Stock       int64
	Description string
	ImageUrl    string
	CreatedAt   string
	UpdatedAt   string
	DeletedAt   string
}

type UpdateProductRequest struct {
	ID          int64   `json:"-"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Stock       int64   `json:"stock"`
	Description string  `json:"description"`
	ImageUrl    string  `json:"image_url"`
}

type UploadImageProductRequest struct {
	ProductImage *multipart.FileHeader `form:"product_image" binding:"required"`
}

type UploadImageProductResponse struct {
	FileName string `json:"file_name"`
	ImageUrl string `json:"image_url"`
}

type RemoveImageProductRequest struct {
	FileName string `json:"file_name" binding:"required"`
}
