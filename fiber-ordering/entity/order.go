package entity

import (
	"context"
	"github.com/gofiber/fiber/v2"
)

// Order represent orders data
type Order struct {
	OrderID     uint32  `json:"order_id"` // auto increment
	CustomerID  uint32  `json:"customer_id"`
	TotalPrice  float32 `json:"total_price"`
}

type OrderResponse struct {
	TotalPrice float32 `json:"total_price"`
	Orders []Order
}

// OrderDetails represent order_details data
type OrderDetails struct {
	OrderDetailsID uint32  `json:"order_details_id"`
	OrderID        uint32  `json:"order_id"`
	ProductID      uint32  `json:"product_id"`
	Quantity       uint32  `json:"quantity"`
	TotalPrice     float32 `json:"total_price"`
}

// OrderWithDetails for aggregating
type OrderWithDetails struct {
	OrderInfo Order          `json:"order_info"`
	Details   []OrderDetails `json:"details"`
}

// OrderRepository represents repo object for order
type OrderRepository interface {
	GetByCustID(custID uint32) ([]Order, error)
	GetOrderDetailsByOrderID(orderID uint32) ([]OrderDetails, error)
	GetAll() ([]Order, error)
	GetByID(OrderID uint32) (*Order, error)
	UpdateByID(c context.Context, orderID uint32, order *Order) error
	UpdateTotalPrice(c context.Context, orderID uint32, totalPrice float32) error
	DeleteByID(c context.Context, OrderID uint32) error
	Store(c context.Context, ord *Order) (uint32, error)
}

// OrderDetailsRepository represents repo object for order
type OrderDetailsRepository interface {
	GetOrderDetailsByOrderID(orderID uint32) ([]OrderDetails, error)
	GetAll() ([]OrderDetails, error)
	GetByID(orderDetailsID uint32) (*OrderDetails, error)
	UpdateByID(c context.Context, orderDetailsID uint32, order *OrderDetails) error
	UpdateTotalPrice(c context.Context, orderDetailsID uint32) error
	DeleteByID(c context.Context, orderDetailsID uint32) error
	Store(c context.Context, ord *OrderDetails) (uint32, error)
}

// OrderService to create service for order
type OrderService interface {
	GetOrderWithDetails(orderID uint32) (OrderWithDetails, error)
	GetAllOrders() ([]Order, error)
	GetOrderByID(orderID uint32) (*Order, error)
	GetOrderDetailByID(orderID uint32) (*OrderDetails, error)
	GetOrdersByCustID(customerID uint32) (OrderResponse, error)
	CreateOrder(c *fiber.Ctx, order *Order) (uint32, error)
	UpdateOrder(c *fiber.Ctx, orderID uint32, order *Order) error
	UpdateOrderPrice(c *fiber.Ctx, orderID uint32) error
	UpdateOrderDetailPrice(c *fiber.Ctx, orderDetailID uint32) error
	CreateOrderDetail(c *fiber.Ctx, orderD *OrderDetails) (uint32, error)
	DeleteOrder(c *fiber.Ctx, orderID uint32) error
	DeleteOrderDetail(c *fiber.Ctx, orderDetailID uint32) error
}