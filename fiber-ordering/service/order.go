package service

import (
	"context"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type orderService struct {
	orderRepo        	entity.OrderRepository
	orderDetailsRepo 	entity.OrderDetailsRepository
	productService 		entity.ProductService
}

// NewOrderService will create new order service
func NewOrderService(
	ouc entity.OrderRepository,
	odc entity.OrderDetailsRepository,
	productService entity.ProductService,
) entity.OrderService {
	return &orderService{
		orderRepo:        ouc,
		orderDetailsRepo: odc,
		productService: productService,
	}
}

func (u *orderService) GetOrderWithDetails(orderID uint32) (res entity.OrderWithDetails, err error) {
	order, err := u.orderRepo.GetByID(orderID)
	if err != nil {
		return
	}

	orderDetails, err := u.orderRepo.GetOrderDetailsByOrderID(orderID)
	if err != nil {
		return
	}

	res.Details = orderDetails
	res.OrderInfo = *order

	return
}

func (u *orderService) GetAllOrders() (res []entity.Order, err error) {
	res, err = u.orderRepo.GetAll()
	return
}

func (u *orderService) GetOrderByID(orderID uint32) (result *entity.Order, err error) {
	res, err := u.orderRepo.GetByID(orderID)
	// get customer id from order
	// prices
	ordersByCustomerID,_ := u.GetOrderWithDetails(orderID)
	details := ordersByCustomerID.Details

	var prices float32
	for _, v := range details {
		prices += v.TotalPrice
		log.Println(prices)
	}

	res.TotalPrice = prices
	result = res
	return
}

func (u *orderService) GetOrderDetailByID(orderDetailID uint32) (res *entity.OrderDetails, err error) {
	res, err = u.orderDetailsRepo.GetByID(orderDetailID)
	return
}

func (u *orderService) GetOrdersByCustID(customerID uint32) (res entity.OrderResponse, err error) {
	listOrders, err := u.orderRepo.GetByCustID(customerID)

	var ordersWithDetails []entity.OrderWithDetails

	for _, v := range listOrders {
		ordersDetailsByID, _ := u.GetOrderWithDetails(v.OrderID)
		ordersWithDetails = append(ordersWithDetails, ordersDetailsByID)
	}

	var totalPrice float32

	for _, v := range ordersWithDetails {
		log.Println(v.OrderInfo.TotalPrice)
		totalPrice += v.OrderInfo.TotalPrice
	}

	res.TotalPrice = totalPrice
	res.Orders = listOrders

	return
}

func (u *orderService) CreateOrder(c *fiber.Ctx, order *entity.Order) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	id, err = u.orderRepo.Store(ctx, order)
	return
}

func (u *orderService) UpdateOrder(c *fiber.Ctx, orderID uint32, order *entity.Order) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.orderRepo.UpdateByID(
		ctx,
		orderID,
		order,
	)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderService) CreateOrderDetail(c *fiber.Ctx, order *entity.OrderDetails) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	product, _ := u.productService.GetByID(order.ProductID)

	res := entity.OrderDetails{
		OrderID:        order.OrderID,
		ProductID:      order.ProductID,
		Quantity:       order.Quantity,
		TotalPrice:     float32(order.Quantity) * product.Price,
	}
	log.Println(res)

	id, err = u.orderDetailsRepo.Store(
		ctx,
		&res,
	)

	// update order price
	err = u.UpdateOrderPrice(c, res.OrderID)
	return
}

func (u *orderService) UpdateOrderPrice(c *fiber.Ctx, orderID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	orderDetails, err := u.orderDetailsRepo.GetOrderDetailsByOrderID(orderID)
	if err != nil {
		return err
	}

	var totalPrice float32

	for _, ord := range orderDetails {
		totalPrice += ord.TotalPrice
	}

	err = u.orderRepo.UpdateTotalPrice(ctx, orderID, totalPrice)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderService) UpdateOrderDetailPrice(c *fiber.Ctx, orderDetailID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.orderDetailsRepo.UpdateTotalPrice(ctx, orderDetailID)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderService) DeleteOrder(c *fiber.Ctx, orderID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.orderRepo.DeleteByID(ctx, orderID)
	if err != nil {
		return err
	}

	return nil
}

func (u *orderService) DeleteOrderDetail(c *fiber.Ctx, orderDetailID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.orderDetailsRepo.DeleteByID(ctx, orderDetailID)
	if err != nil {
		return err
	}

	return nil
}