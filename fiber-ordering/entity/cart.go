package entity

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Cart for cart
type Cart struct {
	CartID      	uint32  `json:"cart_id"`
	Quantity  		uint32	`json:"quantity"`
	TotalPrice      float32  `json:"total_price"`
	ProductID     	uint32  `json:"product_id"`
}

// CreateCartRequest for createCartRequest body request
type CreateCartRequest struct {
	CartID      	uint32  `json:"cart_id"`
	Quantity  		uint32	`json:"quantity"`
	ProductID     	uint32  `json:"product_id"`
}

type CartCompSingle struct {
	CartID      	uint32  		`json:"cart_id"`
	Quantity  		uint32			`json:"quantity"`
	TotalPrice      float32  		`json:"total_price"`
	Products       	ProductComp   	`json:"product"`
}

// CartComp :nodoc:
type CartComp struct {
	CartID      	uint32  		`json:"cart_id"`
	Quantity  		uint32			`json:"quantity"`
	TotalPrice      float32  		`json:"total_price"`
	Products       	[]ProductComp   `json:"products"`
}

// CartRepository nn
type CartRepository interface {
	GetAll() ([]Cart, error)
	GetByID(cartID uint32) (*Cart, error)
	UpdateByID(c context.Context, cartID uint32, order *Cart) error
	DeleteByID(c context.Context, cartID uint32) error
	Store(c context.Context, ord *Cart) (cartID uint32, err error)
	UpdateTotalPrice(ctx context.Context, cartID uint32, updatePrice float32) error
}

// CartService :nododc:
type CartService interface {
	GetAll() (res []CartCompSingle, err error)
	CreateCart(c *fiber.Ctx, order *CreateCartRequest) (id uint32, err error)
	DeleteCart(c *fiber.Ctx, id uint32) (res *CartComp, err error)
	GetByID(id uint32) (res *CartCompSingle, err error)
	UpdatePrice(c *fiber.Ctx, u uint32) error
}