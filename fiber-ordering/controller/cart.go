package controller

import (
	"errors"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

type cartAPI struct {
	cartService entity.CartService
}

// NewCartAPI will initiate cart API
func NewCartAPI(app *fiber.App, mu entity.CartService) {
	cartAPI := &cartAPI{cartService: mu}

	api := app.Group("/api/v1")

	api.Post("/cart", cartAPI.createCart) //ok
	api.Get("/cart", cartAPI.getAll) //ok
	api.Get("/cart/:id_cart", cartAPI.getByID) //ok
	api.Delete("/cart/:id_cart", cartAPI.deleteCart)
	api.Put("/cart/:id_cart/update", cartAPI.updatePriceById)
}

func (a cartAPI) getAll(c *fiber.Ctx) error {
	result, err := a.cartService.GetAll()
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New(err.Error() + "failed to get order data"), c)
	}

	var data struct {
		Data []entity.CartCompSingle `json:"data"`
	}

	log.Println(data)

	data.Data = result
	return Response(data, fiber.StatusOK,nil, c)
}

func (a cartAPI) createCart(c *fiber.Ctx) error {
	body := entity.CreateCartRequest{}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	id, err := a.cartService.CreateCart(c, &body)
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New("cannot create product"), c)
	}

	var response struct {
		ID uint32 `json:"cart_id"`
	}
	response.ID = id

	return Response(response, fiber.StatusCreated, nil, c)

}

func (a cartAPI) deleteCart(c *fiber.Ctx) error {
	panic("")
}

func (a cartAPI) getByID(c *fiber.Ctx) error {
	cartID, err := strconv.Atoi(c.Params("id_cart"))
	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	data, err := a.cartService.GetByID(uint32(cartID))
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New(err.Error() + "failed to delete payment by id"), c)
	}

	var response struct {
		Data interface{} `json:"data"`
	}

	response.Data = data

	return Response(response, fiber.StatusOK, nil, c)
}

func (a cartAPI) updatePriceById(c *fiber.Ctx) error {
	cartID, err := strconv.Atoi(c.Params("id_cart"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	err = a.cartService.UpdatePrice(c, uint32(cartID))
	if err != nil {
		return Response(nil, fiber.StatusBadRequest, errors.New(err.Error() + "failed to update payment status"), c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}