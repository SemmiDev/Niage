package service

import (
	"context"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

type paymentTypeService struct {
	PaymentTypeRepo entity.PaymentTypeRepository
}

// NewPaymentTypeService will initiate service
func NewPaymentTypeService(ptr entity.PaymentTypeRepository) entity.PaymentTypeService {
	return &paymentTypeService{
		PaymentTypeRepo: ptr,
	}
}

func (u *paymentTypeService) GetAll() (res []entity.PaymentType, err error) {
	res, err = u.PaymentTypeRepo.GetAll()
	return
}

func (u *paymentTypeService) GetByID(paymentTypeID uint32) (res *entity.PaymentType, err error) {
	res, err = u.PaymentTypeRepo.GetByID(paymentTypeID)
	return
}

func (u *paymentTypeService) DeleteByID(c *fiber.Ctx, paymentTypeID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.PaymentTypeRepo.DeleteByID(ctx, paymentTypeID)
	return err
}

func (u *paymentTypeService) CreatePaymentType(c *fiber.Ctx, paymentType *entity.PaymentType) (res uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res, err = u.PaymentTypeRepo.Store(ctx, paymentType)
	return
}