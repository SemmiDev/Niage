package entity

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// PaymentType for payment type
type PaymentType struct {
	PaymentTypeID uint32 `json:"payment_type_id"`
	Method        string `json:"payment_method"`
	Company       string `json:"payment_company"`
}

// PaymentTypeRepository for repo
type PaymentTypeRepository interface {
	GetAll() ([]PaymentType, error)
	GetByID(paymentTypeID uint32) (*PaymentType, error)
	DeleteByID(c context.Context, paymentTypeID uint32) error
	Store(c context.Context, paymentType *PaymentType) (uint32, error)
}

// PaymentTypeService for service
type PaymentTypeService interface {
	GetAll() ([]PaymentType, error)
	GetByID(paymentTypeID uint32) (*PaymentType, error)
	DeleteByID(c *fiber.Ctx, paymentTypeID uint32) error
	CreatePaymentType(c *fiber.Ctx, paymentType *PaymentType) (uint32, error)
}
