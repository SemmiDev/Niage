package controller

import (
	"errors"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type orderAPI struct {
	orderService entity.OrderService
}

// NewOrderAPI will initiate order API
func NewOrderAPI(app *fiber.App, ouc entity.OrderService) {
	orderAPI := &orderAPI{
		orderService: ouc,
	}

	api := app.Group("/api/v1")

	api.Get("/order", orderAPI.getAllOrders) //ok
	api.Post("/order",  orderAPI.createOrder) //ok
	api.Get("/order/:id_order", orderAPI.getOrderByID) //ok
	api.Get("/order/:id_order/details", orderAPI.getOrderWithDetails) //ok
	api.Post("/order/:id_order/details",  orderAPI.createOrderDetail) //ok
	api.Get("/order/detail/:id_order_detail", orderAPI.getOrderDetailByID) //ok

	// key
	api.Get("/order/customer/:id_customer", orderAPI.getOrderByCustomer)

	api.Put("/order/:id_order/price",  orderAPI.updateOrderPrice)
	api.Put("/order/detail/:id_order_detail/price",  orderAPI.updateOrderDetailPrice)

	api.Delete("/order/:id_order",  orderAPI.deleteOrderByID)
	api.Delete("/order/detail/:id_order_detail",  orderAPI.deleteOrderDetailByID)
}

type OrdersResponse struct {
	OrderID     uint32  `json:"order_id"`
	CustomerID  uint32  `json:"customer_id"`
}

func (t *orderAPI) getAllOrders(c *fiber.Ctx) error {
	result, err := t.orderService.GetAllOrders()
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get order data"), c)
	}

	var results []OrdersResponse
	var resultTemp OrdersResponse

	for _, v := range result {
		resultTemp.OrderID = v.OrderID
		resultTemp.CustomerID = v.CustomerID

		results = append(results, resultTemp)
	}

	var data struct {
		Data []OrdersResponse `json:"data"`
	}

	data.Data = results
	return Response(data, fiber.StatusOK,nil, c)
}

func (t *orderAPI) getOrderByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_order"))
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order") + "is not integer"), c)
	}

	result, err := t.orderService.GetOrderByID(uint32(id))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + " | failed to get user order by id"), c)
	}

	var RES struct{
		OrderID     uint32  `json:"order_id"`
		CustomerID  uint32  `json:"customer_id"`
	}

	var data struct {
		Data interface{} `json:"data"`
	}

	RES.OrderID = result.OrderID
	RES.CustomerID = result.CustomerID

	data.Data = RES
	return Response(data, fiber.StatusOK, nil,c)
}

func (t *orderAPI) getOrderDetailByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id_order_detail"))
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order_detail") + "is not integer"), c)
	}

	result, err := t.orderService.GetOrderDetailByID(uint32(id))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "| failed to get user order by id"), c)
	}

	var data struct {
		Data *entity.OrderDetails `json:"data"`
	}

	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}

func (t *orderAPI) getOrderWithDetails(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("id_order"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order") + " | is not integer"), c)
	}

	result, err := t.orderService.GetOrderWithDetails(uint32(orderID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(c.Params("id_order") + " | failed to get user order by id"), c)
	}

	var data struct {
		Data entity.OrderWithDetails `json:"data"`
	}

	data.Data = result

	return Response(data, fiber.StatusOK, nil, c)
}

func (t *orderAPI) getOrderByCustomer(c *fiber.Ctx) error {
	cutID, err := strconv.ParseInt(c.Params("id_customer"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_customer") + " | is not integer"), c)
	}

	result, err := t.orderService.GetOrdersByCustID(uint32(cutID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get order by customer"), c)
	}

	var data struct {
		Data interface{} `json:"data"`
	}

	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}

func (t *orderAPI) createOrder(c *fiber.Ctx) error {
	var body struct {
		CustomerID  uint32  `json:"customer_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusBadRequest, errors.New(err.Error() + errors.New(" | bad request body").Error()), c)
	}

	order := entity.Order{
		CustomerID:  body.CustomerID,
	}

	id, err := t.orderService.CreateOrder(c, &order)
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + errors.New(" | failed to create order").Error()), c)
	}

	var response struct {
		OrderID uint32 `json:"order_id"`
	}
	response.OrderID = id
	return Response(response, fiber.StatusCreated, nil, c)
}

func (t *orderAPI) createOrderDetail(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("id_order"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order") + " | is not integer"), c)
	}

	var body struct {
		MenuID     uint32  `json:"product_id"`
		Quantity   uint32  `json:"quantity"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	ordDetail := entity.OrderDetails{
		OrderID:    uint32(orderID),
		ProductID:  body.MenuID,
		Quantity:   body.Quantity,
	}

	id, err := t.orderService.CreateOrderDetail(c,&ordDetail)

	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to create order detail"), c)
	}

	var response struct {
		OrdDetailID uint32 `json:"order_details_id"`
	}
	response.OrdDetailID = id
	return Response(response, fiber.StatusCreated, nil ,c)
}

func (t *orderAPI) updateOrderPrice(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("id_order"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order") + " | is not integer"), c)
	}

	err = t.orderService.UpdateOrderPrice(c, uint32(orderID))

	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to update order total price"), c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}

func (t *orderAPI) updateOrderDetailPrice(c *fiber.Ctx) error {
	orderDetailID, err := strconv.ParseInt(c.Params("id_order_detail"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order_detail") + " | is not integer"), c)
	}

	err = t.orderService.UpdateOrderDetailPrice(c, uint32(orderDetailID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to update order detail total price"), c)
	}
	return Response(nil, fiber.StatusOK, nil, c)
}

func (t *orderAPI) deleteOrderByID(c *fiber.Ctx) error {
	orderID, err := strconv.ParseInt(c.Params("id_order"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order")+" | is not integer"), c)
	}

	err = t.orderService.DeleteOrder(c, uint32(orderID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to delete order by id"), c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}

func (t *orderAPI) deleteOrderDetailByID(c *fiber.Ctx) error {
	orderDetailID, err := strconv.ParseInt(c.Params("id_order_detail"), 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(c.Params("id_order_detail")+" | is not integer"), c)
	}

	err = t.orderService.DeleteOrderDetail(c, uint32(orderDetailID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to delete order detail by id"), c)
	}
	return Response(nil, fiber.StatusOK, nil, c)
}