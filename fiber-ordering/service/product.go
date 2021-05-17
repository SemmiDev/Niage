package service

import (
	"context"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

type productService struct {
	menuRepo     entity.ProductRepository
	menuCategoryRepo entity.ProductCategoryRepository
}

// NewProductService will create new order service
func NewProductService(
	muc entity.ProductRepository,
	mtc entity.ProductCategoryRepository,
) entity.ProductService {
	return &productService{
		menuRepo:     muc,
		menuCategoryRepo: mtc,
	}
}

func (t *productService) GetAll() (res []entity.ProductComp, err error) {
	menuRAW, err := t.menuRepo.GetAll()
	if err != nil {
		return
	}

	res = make([]entity.ProductComp, len(menuRAW))
	for i, m := range menuRAW {

		res[i].ProductID = m.ProductID
		res[i].Code = m.Code
		res[i].Name = m.Name
		res[i].ProductImage = m.ProductImage
		res[i].Price = m.Price

		var typeTemp *entity.ProductCategory
		typeTemp, err = t.menuCategoryRepo.GetByID(m.CategoryID)
		if err != nil {
			return
		}

		res[i].ProductCategory = *typeTemp
		res[i].ProductStatus = m.ProductStatus
	}

	return
}

func (t *productService) CreateProduct(c *fiber.Ctx, order *entity.Product) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	id, err = t.menuRepo.Store(ctx, order)
	return
}

func (t *productService) DeleteProduct(c *fiber.Ctx, id uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := t.menuRepo.DeleteByID(ctx, id)

	if err != nil {
		return err
	}
	return nil
}

func (t *productService) GetByID(id uint32) (res *entity.ProductComp, err error) {
	m, err := t.menuRepo.GetByID(id)
	if err != nil {
		return
	}

	if m == nil {
		return
	}

	typeTemp, err := t.menuCategoryRepo.GetByID(m.CategoryID)
	if err != nil {
		return
	}

	menuComp := entity.ProductComp{
		ProductID:      	m.ProductID,
		Code:        		m.Code,
		Name:        		m.Name,
		ProductImage:  		m.ProductImage,
		Price:       		m.Price,
		ProductCategory:    *typeTemp,
		ProductStatus:  	m.ProductStatus,
	}

	res = &menuComp

	return
}

func (t *productService) GetAllCategory() (res []entity.ProductCategory, err error) {
	res, err = t.menuCategoryRepo.GetAll()
	return
}

func (t *productService) CreateCategory(c *fiber.Ctx, m *entity.ProductCategory) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	id, err = t.menuCategoryRepo.Store(ctx, m)
	return
}

func (t *productService) DeleteCategory(c *fiber.Ctx, id uint32) (res *entity.ProductCategory, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res, err = t.menuCategoryRepo.GetByID(id)
	if err != nil {
		return
	}
	err = t.menuCategoryRepo.DeleteByID(ctx, id)
	return
}