package controller

import (
	"errors"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
)

type productAPI struct {
	productService entity.ProductService
}

// NewProductAPI will initiate product API
func NewProductAPI(app *fiber.App, mu entity.ProductService) {
	productAPI := &productAPI{
		productService: mu,
	}

	api := app.Group("/api/v1")

	api.Get("/product", productAPI.getAll)
	api.Post("/product", productAPI.createProduct)
	api.Get("/product/category", productAPI.getAllType)
	api.Post("/product/category", productAPI.createType)
	api.Delete("/product/category/:id", productAPI.deleteType)
	api.Delete("/product/:id", productAPI.deleteProduct)
	api.Get("/product/:id", productAPI.getByID)
}

func (t *productAPI) getAll(c *fiber.Ctx) error {
	result, err := t.productService.GetAll()
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError, err, c)
	}

	var data struct {
		Data []entity.ProductComp `json:"data"`
	}

	data.Data = result
	return Response(data,fiber.StatusOK, nil, c)
}

func (t *productAPI) getByID(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	res, err := t.productService.GetByID(uint32(id))
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New("failed to get product by id"), c)
	}

	var data struct {
		Data *entity.ProductComp `json:"data"`
	}
	data.Data = res
	return Response(data, fiber.StatusOK, nil, c)
}

func (t *productAPI) getAllType(c *fiber.Ctx) error {
	result, err := t.productService.GetAllCategory()
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New("failed to get product data"), c)
	}

	var data struct {
		Data []entity.ProductCategory `json:"data"`
	}

	data.Data = result
	return Response(data, fiber.StatusOK, nil, c)
}

func (t *productAPI) createProduct(c *fiber.Ctx) error {
	body := entity.Product{}
	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	if body.Name == "" {
		return Response(nil,fiber.StatusBadRequest, errors.New("please provide name"), c)
	}

	id, err := t.productService.CreateProduct(c, &body)
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New("cannot create product"), c)
	}

	var response struct {
		ID uint32 `json:"product_id"`
	}
	response.ID = id

	return Response(response, fiber.StatusCreated, nil, c)
}

func (t *productAPI) deleteProduct(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, err, c)
	}

	err = t.productService.DeleteProduct(c, uint32(id))
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError, errors.New("failed to delete product"), c)
	}

	return Response(nil, fiber.StatusOK,nil, c)
}

func (t *productAPI) createType(c *fiber.Ctx) error {
	body := entity.ProductCategory{}

	if err := c.BodyParser(&body); err != nil {
		return Response(body, http.StatusUnprocessableEntity, err, c)
	}

	if body.CategoryName == "" {
		return Response(nil, fiber.StatusBadRequest, errors.New("StatusBadRequest"), c)
	}

	id, err := t.productService.CreateCategory(c, &body)
	if err != nil {
		return Response(nil,fiber.StatusInternalServerError,errors.New("failed to create product type"), c)
	}

	var response struct {
		ID uint32 `json:"product_type_id"`
	}

	response.ID = id
	return Response(response, fiber.StatusCreated, nil, c)
}

func (t *productAPI) deleteType(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		return Response(nil,fiber.StatusBadRequest, errors.New("id not integer"), c)
	}

	_, err = t.productService.DeleteCategory(c, uint32(id))
	if err != nil {
		return Response(nil, fiber.StatusInternalServerError,err, c)
	}

	return Response(nil, fiber.StatusOK, nil, c)
}
