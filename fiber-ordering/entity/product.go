package entity

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Product for product
type Product struct {
	ProductID      	uint32  `json:"product_id"`
	Code  			string	`json:"product_code"`
	Name        	string  `json:"product_name"`
	ProductImage  	string  `json:"product_image"`
	Price       	float32 `json:"product_price"`
	CategoryID  	uint32  `json:"product_category_id"`
	ProductStatus  	bool    `json:"product_status"`
}

// ProductCategory for product type
type ProductCategory struct {
	ProductCategoryID  	uint32 `json:"product_category_id"`
	CategoryName    	string `json:"product_category_name"`
}

// ProductComp :nodoc:
type ProductComp struct {
	ProductID      	uint32   `json:"product_id"`
	Code        	string   `json:"product_code"`
	Name        	string   `json:"product_name"`
	ProductImage  	string   `json:"product_image"`
	Price       	float32  `json:"price"`
	ProductCategory ProductCategory `json:"product_category"`
	ProductStatus  	bool     `json:"product_status"`
}

// ProductRepository nn
type ProductRepository interface {
	GetAll() ([]Product, error)
	GetByID(productID uint32) (*Product, error)
	UpdateByID(c context.Context, productID uint32, order *Product) error
	DeleteByID(c context.Context, productID uint32) error
	Store(c context.Context, ord *Product) (productID uint32, err error)
}

// ProductCategoryRepository repo
type ProductCategoryRepository interface {
	GetAll() ([]ProductCategory, error)
	GetByID(productCategoryID uint32) (*ProductCategory, error)
	UpdateByID(c context.Context, productCategoryID uint32, order *ProductCategory) error
	DeleteByID(c context.Context, mtypeID uint32) error
	Store(c context.Context, ord *ProductCategory) (productCategoryID uint32, err error)
}

// ProductService :nododc:
type ProductService interface {
	GetAll() (res []ProductComp, err error)
	CreateProduct(c *fiber.Ctx, order *Product) (id uint32, err error)
	DeleteProduct(c *fiber.Ctx, id uint32) (err error)
	GetByID(id uint32) (res *ProductComp, err error)
	GetAllCategory() (res []ProductCategory, err error)
	CreateCategory(c *fiber.Ctx, m *ProductCategory) (id uint32, err error)
	DeleteCategory(c *fiber.Ctx, id uint32) (res *ProductCategory, err error)
}