package entity

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Customer for customer table
type Customer struct {
	CustomerID    uint32 `json:"customer_id"`
	FullName      string `json:"customer_full_name"`
	Email         string `json:"customer_email"`
	Username      string `json:"customer_username"`
	Password      string `json:"customer_password"`
	PhoneNumber   string `json:"customer_phone_number"`
	AccountStatus bool   `json:"account_status"`
}

// CustomerRepository for repo
type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByID(CustomerID uint32) (*Customer, error)
	DeleteByID(c context.Context, CustomerID uint32) error
	Store(c context.Context, cust *Customer) (uint32, error)
	UpdateByID(c context.Context, custID uint32, cust *Customer) error
	GetByEmail(email string) (res *Customer, err error)
}

// CustomerService for service
type CustomerService interface {
	GetAll() ([]Customer, error)
	GetByID(CustomerID uint32) (*Customer, error)
	DeleteByID(c *fiber.Ctx, CustomerID uint32) error
	CreateCustomer(c *fiber.Ctx, cust *Customer) (uint32, error)
	UpdateCustomerByID(c *fiber.Ctx, custID uint32, cust *Customer) error
	Login(email string, password string) string
}