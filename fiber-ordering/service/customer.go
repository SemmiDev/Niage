package service

import (
	"context"
	"fiber-ordering/auth"
	"fiber-ordering/entity"
	"github.com/gofiber/fiber/v2"
	"log"
	"time"
)

type customerService struct {
	CustomerRepo entity.CustomerRepository
}

func (u *customerService) Login(email string, password string) string {
	// find customer by email
	data, err := u.CustomerRepo.GetByEmail(email)
	// kalau data tidak ada,  maka return Silahkan register terlebih dahulu
	if data == nil || err != nil {
		return "Silahkan register terlebih dahulu"
	}
	// kalau data tidak nil, maka compare password nya
	compare, err := auth.ComparePassword(data.Password, password)
	// kalau compare = true, maka return berhasil login
	if compare == true {
		return "anda berhasil login"
	}
	// kalau compare = false, maka return password salah
	log.Println("baris 30 check, dei")
	return "password yang anda masukkan salah"
}

// NewCustomerService will initiate new service
func NewCustomerService(cur entity.CustomerRepository) entity.CustomerService {
	return &customerService{
		CustomerRepo: cur,
	}
}

func (u *customerService) GetAll() (res []entity.Customer, err error) {
	res, err = u.CustomerRepo.GetAll()
	return
}

func (u *customerService) GetByID(CustomerID uint32) (res *entity.Customer, err error) {
	res, err = u.CustomerRepo.GetByID(CustomerID)
	return
}

func (u *customerService) DeleteByID(c *fiber.Ctx, CustomerID uint32) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.CustomerRepo.DeleteByID(ctx, CustomerID)
	return err
}

func (u *customerService) CreateCustomer(c *fiber.Ctx, cust *entity.Customer) (id uint32, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	passHash, _ := auth.GeneratePassword(cust.Password)
	cust.Password = passHash
	id, err = u.CustomerRepo.Store(ctx, cust)
	return
}

func (u *customerService) UpdateCustomerByID(c *fiber.Ctx, custID uint32, cust *entity.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := u.CustomerRepo.UpdateByID(ctx, custID, cust)
	if err != nil {
		return err
	}

	return nil
}
