package service

import (
	"context"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"time"
)

type paymentService struct {
	PaymentRepo entity.PaymentRepository
}

// NewPaymentService will initiate new usecase
func NewPaymentService(pyr entity.PaymentRepository) entity.PaymentService {
	return &paymentService{
		PaymentRepo: pyr,
	}
}

func (u *paymentService) UpdateStatus(c *fiber.Ctx, paymentID uint32, status int32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.PaymentRepo.UpdateStatus(ctx, paymentID, status)
	return err
}

func (u *paymentService) GetAll() (res []entity.Payment, err error) {
	res, err = u.PaymentRepo.GetAll()
	return
}

func (u *paymentService) GetByID(paymentID uint32) (res *entity.Payment, err error) {
	res, err = u.PaymentRepo.GetByID(paymentID)
	return
}

func (u *paymentService) DeleteByID(c *fiber.Ctx, paymentID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.PaymentRepo.DeleteByID(ctx, paymentID)
	return err
}

func (u *paymentService) CreatePayment(c *fiber.Ctx, payment *entity.Payment) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	id, err = u.PaymentRepo.Store(ctx, payment)
	return
}

func (u *paymentService) GetListOfPaymentsByCustomerID(c *fiber.Ctx, customerID uint32) (res []entity.Payment, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	res, err = u.PaymentRepo.GetListOfPaymentsByCustomerID(ctx,customerID)
	return res, err
}