package controller

import (
	"errors"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type paymentAPI struct {
	paymentService     entity.PaymentService
	paymentTypeService entity.PaymentTypeService
}

// NewPaymentAPI will initiate api
func NewPaymentAPI(app *fiber.App, pac entity.PaymentService, ptc entity.PaymentTypeService) {
	paymentAPI := &paymentAPI{
		paymentService:     pac,
		paymentTypeService: ptc,
	}

	api := app.Group("/api/v1")

	api.Get("/payment", paymentAPI.GetAll) //ok
	api.Get("/payment/type", paymentAPI.GetAllType)
	api.Post("/payment", paymentAPI.Create) //ok
	api.Post("/payment/type", paymentAPI.CreateType) //ok
	api.Get("/payment/:id_payment", paymentAPI.GetByID) //ok
	api.Get("/payment/type/:id_payment_type", paymentAPI.GetTypeByID) //ok
	api.Delete("/payment/:id_payment", paymentAPI.DeleteByID)
	api.Delete("/payment/type/:id_payment_type", paymentAPI.DeleteTypeByID)
	api.Put("/payment/:id_payment/status", paymentAPI.UpdateStatus) // ok
	api.Get("/payment/list/:id_customer", paymentAPI.GetListOfPaymentsByCustomerID)
}

func (p *paymentAPI) GetAll(c *fiber.Ctx) error {
	result, err := p.paymentService.GetAll()
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get payment data"), c)
	}
	var data struct {
		Data []entity.Payment `json:"data"`
	}
	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) GetByID(c *fiber.Ctx) error {
	paymentID, err := strconv.Atoi(c.Params("id_payment"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	result, err := p.paymentService.GetByID(uint32(paymentID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error()+"failed to get payment data by id"), c)
	}

	var data struct {
		Data *entity.Payment `json:"data"`
	}

	data.Data = result

	return Response(data, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) GetAllType(c *fiber.Ctx) error {
	result, err := p.paymentTypeService.GetAll()
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get payment type data"), c)
	}

	var data struct {
		Data []entity.PaymentType `json:"data"`
	}

	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) GetTypeByID(c *fiber.Ctx) error {
	paymentTypeID, err := strconv.Atoi(c.Params("id_payment_type"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	result, err := p.paymentTypeService.GetByID(uint32(paymentTypeID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get payment type data by id"), c)
	}

	var data struct {
		Data *entity.PaymentType `json:"data"`
	}

	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) UpdateStatus(c *fiber.Ctx) error {
	paymentID, err := strconv.Atoi(c.Params("id_payment"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	var body struct {
		Status int32 `json:"status"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	err = p.paymentService.UpdateStatus(c, uint32(paymentID), body.Status)
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(err.Error() + "failed to update payment status"), c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) DeleteTypeByID(c *fiber.Ctx) error {
	paymentTypeID, err := strconv.Atoi(c.Params("id_payment_type"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	err = p.paymentTypeService.DeleteByID(c, uint32(paymentTypeID))

	if err != nil {
		return Response(nil, fiber.StatusInternalServerError,errors.New(err.Error()+ "failed to delete payment type by id"), c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) CreateType(c *fiber.Ctx) error {
	var body struct {
		Method  string `json:"method"`
		Company string `json:"company"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	paymentType := entity.PaymentType{
		Method:  body.Method,
		Company: body.Company,
	}

	id, err := p.paymentTypeService.CreatePaymentType(c, &paymentType)
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError,errors.New(err.Error() + "failed to create menu tyoe"), c)
	}

	var response struct {
		ID uint32 `json:"payment_type_id"`
	}
	response.ID = id
	return Response(response, fiber.StatusCreated, nil, c)
}

func (p *paymentAPI) DeleteByID(c *fiber.Ctx) error {
	paymentID, err := strconv.Atoi(c.Params("id_payment"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	err = p.paymentService.DeleteByID(c, uint32(paymentID))
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New(err.Error() + "failed to delete payment by id"), c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}

func (p *paymentAPI) Create(c *fiber.Ctx) error {
	var body struct {
		OrderID       uint32  `json:"order_id"`
		Amount        float32 `json:"amount"`
		PaymentTypeID uint32  `json:"payment_type_id"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	payment := entity.Payment{
		OrderID:       body.OrderID,
		Amount:        body.Amount,
		PaymentTypeID: body.PaymentTypeID,
	}

	id, err := p.paymentService.CreatePayment(c, &payment)
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New(err.Error() + "failed to create payment"), nil)
	}

	var response struct {
		PaymentID uint32 `json:"payment_id"`
	}
	response.PaymentID = id

	return Response(response,fiber.StatusCreated, nil, c)
}

func (p *paymentAPI) GetListOfPaymentsByCustomerID(c *fiber.Ctx) error {
	customerID, err := strconv.Atoi(c.Params("id_customer"))
	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	result, err := p.paymentService.GetListOfPaymentsByCustomerID(c, uint32(customerID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get list of payment by customer id"), c)
	}

	var data struct {
		Data []entity.Payment `json:"data"`
	}

	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}