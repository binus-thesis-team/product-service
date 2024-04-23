package model

type Product struct {
	Id          int64   `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Stock       int64   `json:"stock,omitempty"`
	Description string  `json:"description,omitempty"`
	ImageUrl    string  `json:"image_url,omitempty"`
	CreatedAt   string  `json:"created_at,omitempty"`
	UpdatedAt   string  `json:"updated_at,omitempty"`
	DeletedAt   string  `json:"deleted_at,omitempty"`
}
