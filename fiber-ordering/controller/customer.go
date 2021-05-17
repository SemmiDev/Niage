package controller

import (
	"errors"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type customerAPI struct {
	customerService entity.CustomerService
}

// NewCustomerAPI will initiate new api
func NewCustomerAPI(app *fiber.App, cuc entity.CustomerService) {
	customerAPI := &customerAPI{
		customerService: cuc,
	}

	api := app.Group("/api/v1")

	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).
			JSON(map[string]interface{}{
				"health": "ok",
				"status": http.StatusOK,
			})
	})

	api.Get("/customer", customerAPI.GetAll)
	api.Post("/customer", customerAPI.Create)
	api.Get("/customer/:id_customer", customerAPI.GetByID)
	api.Put("/customer/:id_customer", customerAPI.UpdateByID)
	api.Delete("/customer/:id_customer", customerAPI.DeleteByID)

	// sign in & sign up
	api.Post("/login", customerAPI.Login)
}

func (s *customerAPI) GetAll(c *fiber.Ctx) error {
	result, err := s.customerService.GetAll()
	return Response(result, fiber.StatusOK, err, c)
}

func (s *customerAPI) GetByID(c *fiber.Ctx) error {
	id := c.Params("id_customer")

	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}

	result, err := s.customerService.GetByID(uint32(customerID))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, errors.New("failed to get user customer by id"), c)
	}

	var data struct {
		Data *entity.Customer `json:"data"`
	}

	data.Data = result
	return Response(data, http.StatusOK, err, c)
}

func (s *customerAPI) DeleteByID(c *fiber.Ctx) error {
	id := c.Params("id_customer")

	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}

	err = s.customerService.DeleteByID(c, uint32(customerID))
	if err != nil {
		if err != nil {
			return Response(nil, fiber.StatusInternalServerError, errors.New("failed to delete customer by id"), c)
		}
	}

	return Response(nil,fiber.StatusOK, err, c)
}

func (s *customerAPI) UpdateByID(c *fiber.Ctx) error {
	id := c.Params("id_customer")

	var body struct {
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	customer := entity.Customer{
		FullName:    body.FullName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
	}

	customerID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return Response(nil, fiber.StatusUnprocessableEntity, errors.New("id is not defined"), c)
	}

	err = s.customerService.UpdateCustomerByID(c, uint32(customerID), &customer)
	if err != nil {
		return Response(body, http.StatusInternalServerError, err, c)
	}
	return Response(fiber.Map{"status": "updated"}, fiber.StatusOK, nil, c)
}

func (s *customerAPI) Create(c *fiber.Ctx) error {
	var body struct {
		FullName    string `json:"full_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Username    string `json:"username"`
		Password    string `json:"password"`
	}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	cust := entity.Customer{
		FullName:    body.FullName,
		Email:       body.Email,
		PhoneNumber: body.PhoneNumber,
		Username:    body.Username,
		Password:    body.Password,
	}

	id, err := s.customerService.CreateCustomer(c, &cust)
	if err != nil {
		return Response(body, http.StatusInternalServerError, err, c)
	}

	var response struct {
		CustID uint32 `json:"customer_id"`
	}
	response.CustID = id

	return Response(response, fiber.StatusCreated, nil, c)
}

func (s *customerAPI) Login(c *fiber.Ctx) error {
	email := c.FormValue("email")
	password := c.FormValue("password")

	status := s.customerService.Login(email, password)
	return Response(fiber.Map{"login" : status}, fiber.StatusOK, nil, c)
}