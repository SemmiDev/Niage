package entity

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Payment represent payment data
type Payment struct {
	PaymentID     uint32  `json:"payment_id"`
	OrderID       uint32  `json:"order_id"`
	Amount        float32 `json:"amount"`
	PaymentTypeID uint32  `json:"payment_type_id"`
	PaymentDate   string  `json:"payment_date"`
	PaymentStatus int32   `json:"payment_status"` // 0:pending, 1:completed, 1:cancelled
}

// PaymentRepository payment repo
type PaymentRepository interface {
	GetAll() ([]Payment, error)
	GetByID(paymentID uint32) (*Payment, error)
	DeleteByID(c context.Context, paymentID uint32) error
	Store(c context.Context, payment *Payment) (uint32, error)
	UpdateStatus(c context.Context, paymentID uint32, status int32) error
	GetListOfPaymentsByCustomerID(c context.Context, customerID uint32) ([]Payment, error)
}

// PaymentService for payment
type PaymentService interface {
	GetAll() ([]Payment, error)
	GetByID(paymentID uint32) (*Payment, error)
	DeleteByID(c *fiber.Ctx, paymentID uint32) error
	CreatePayment(c *fiber.Ctx, payment *Payment) (uint32, error)
	UpdateStatus(c *fiber.Ctx, paymentID uint32, status int32) error
	GetListOfPaymentsByCustomerID(c *fiber.Ctx, customerID uint32) ([]Payment, error)
}