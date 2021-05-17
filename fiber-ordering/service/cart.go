package service

import (
	"context"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type cartService struct {
	cartRepo			entity.CartRepository
	productService 		entity.ProductService
}

// NewCartService will create new cart service
func NewCartService(muc entity.CartRepository, menuRepo entity.ProductRepository, menuCategoryRepo entity.ProductCategoryRepository, productService entity.ProductService) entity.CartService {
	return &cartService{
		cartRepo: muc,
		productService: productService,
	}
}

func (c *cartService) GetAll() (res []entity.CartCompSingle, err error) {
	carts, err := c.cartRepo.GetAll()
	var products []entity.ProductComp
	var results []entity.CartCompSingle

	for _, val := range carts {
		product, _ := c.productService.GetByID(val.ProductID)
		products = append(products, *product)
	}
	for index, v := range carts {
		results = append(results, entity.CartCompSingle{
			CartID:     v.CartID,
			Quantity:   v.Quantity,
			TotalPrice: v.TotalPrice,
			Products:   products[index],
		})
	}
	res = results
	return
}

func (s *cartService) CreateCart(c *fiber.Ctx, cart *entity.CreateCartRequest) (cartID uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	product, _ := s.productService.GetByID(cart.ProductID)
	res := entity.Cart{
		CartID:     cartID,
		Quantity:   cart.Quantity,
		TotalPrice: product.Price * float32(cart.Quantity),
		ProductID:  cart.ProductID,
	}

	cartID, err = s.cartRepo.Store(ctx, &res)
	log.Println(err)
	return
}

func (s *cartService) DeleteCart(c *fiber.Ctx, id uint32) (res *entity.CartComp, err error) {
	panic("implement me")
}

func (s *cartService) UpdatePrice(c *fiber.Ctx, id uint32) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	// get cart by id
	m, err := s.cartRepo.GetByID(id)
	if err != nil {
		return err
	}
	productID := m.ProductID
	realProduct, _ := s.productService.GetByID(productID)
	productPrice := realProduct.Price
	updatePrice := productPrice * float32(m.Quantity)
	s.cartRepo.UpdateTotalPrice(ctx, id, updatePrice)
	return nil
}

func (s *cartService) GetByID(id uint32) (res *entity.CartCompSingle, err error) {
	// get cart by id
	m, err := s.cartRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	// get product id from cart
	productID := m.ProductID
	// get product from product service by id
	result, err := s.productService.GetByID(productID)
	res = &entity.CartCompSingle{
		CartID:     m.CartID,
		Quantity:   m.Quantity,
		TotalPrice: m.TotalPrice,
		Products: entity.ProductComp{
			ProductID:    result.ProductID,
			Code:         result.Code,
			Name:         result.Name,
			ProductImage: result.ProductImage,
			Price:        result.Price,
			ProductCategory: entity.ProductCategory{
				ProductCategoryID: result.ProductCategory.ProductCategoryID,
				CategoryName:      result.ProductCategory.CategoryName,
			},
			ProductStatus: result.ProductStatus,
		},
	}
	return
}