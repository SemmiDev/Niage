package main

import (
	"database/sql"
	"fiber-ordering/config"
	"fiber-ordering/controller"
	"fiber-ordering/repository"
	"fiber-ordering/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	_ "github.com/joho/godotenv/autoload"
	"log"
	"os"
	"strconv"
	"time"
)

// Server represents server
type Server struct {
	Reader      *sql.DB
	Writer      *sql.DB
	Port        string
}

func configureMySQL() (*sql.DB, *sql.DB) {

	readerConfig := config.Option{
		Host:     config.MySQL.ReaderHost,
		Port:     config.MySQL.ReaderPort,
		Database: config.MySQL.Database,
		User:     config.MySQL.ReaderUser,
		Password: config.MySQL.ReaderPassword,
	}

	writerConfig := config.Option{
		Host:     config.MySQL.WriterHost,
		Port:     config.MySQL.WriterPort,
		Database: config.MySQL.Database,
		User:     config.MySQL.WriterUser,
		Password: config.MySQL.WriterPassword,
	}

	reader, writer, err := config.SetupDatabase(readerConfig, writerConfig)
	if err != nil {
		log.Fatalf("%s: %s", "Failed to connect mysql", err)
	}

	log.Println("MySQL connection is successfully established!")

	return reader, writer
}

//func init() {
//	os.Setenv("PORT", "9090")
//	os.Setenv("READER_HOST", "localhost")
//	os.Setenv("READER_PORT", "3306")
//	os.Setenv("READER_USER", "sammidev")
//	os.Setenv("READER_PASSWORD", "sammidev")
//	os.Setenv("WRITER_HOST", "localhost")
//	os.Setenv("WRITER_PORT", "3306")
//	os.Setenv("WRITER_USER", "sammidev")
//	os.Setenv("WRITER_PASSWORD", "sammidev")
//	os.Setenv("MYSQL_DATABASE_NAME", "shop")
//	os.Setenv("TIMEOUT_ON_SECONDS", "120")
//	os.Setenv("OPERATION_ON_EACH_CONTEXT", "500")
//	os.Setenv("SERVER_READ_TIMEOUT", "60")
//	os.Setenv("DB_MAX_CONNECTIONS", "100")
//	os.Setenv("JWT_SECRET_KEY", "secretsekaliy")
//	os.Setenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT", "30")
//}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	reader, writer := configureMySQL()
	server := Server{
		Reader:      reader,
		Writer:      writer,
		Port:        config.Server.Port,
	}

	server.Run()
}

func (s *Server) Run() {
	port := config.Server.Port
	if port == "" {
		port = ":9090"
	}

	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

	// fiber initialize
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
	})

	// customer
	customerRepo := repository.NewCustomerRepo(s.Reader, s.Writer)
	customerService := service.NewCustomerService(customerRepo)

	// menu
	menuRepo := repository.NewProductRepo(s.Reader, s.Writer)
	menuCategoryRepo := repository.NewProductCategoryRepo(s.Reader, s.Writer)
	menuService := service.NewProductService(menuRepo, menuCategoryRepo)

	// order
	orderRepo := repository.NewOrderRepo(s.Reader, s.Writer)
	orderDetailsRepo := repository.NewOrderDetailsRepo(s.Reader, s.Writer)
	orderService := service.NewOrderService(orderRepo, orderDetailsRepo, menuService)

	// car
	cartRepo := repository.NewCartRepo(s.Reader, s.Writer)
	cartService := service.NewCartService(cartRepo, menuRepo, menuCategoryRepo, menuService)

	// payment
	paymentRepo := repository.NewPaymentRepo(s.Reader, s.Writer)
	paymentTypeRepo := repository.NewPaymentTypeRepo(s.Reader, s.Writer)
	paymentService := service.NewPaymentService(paymentRepo)
	paymentTypeService := service.NewPaymentTypeService(paymentTypeRepo)

	controller.NewCustomerAPI(app, customerService)
	controller.NewProductAPI(app, menuService)
	controller.NewOrderAPI(app, orderService)
	controller.NewCartAPI(app, cartService)
	controller.NewPaymentAPI(app, paymentService, paymentTypeService)

	app.Use(logger.New())
	app.Use(cors.New())

	app.Listen(fmt.Sprintf(":%s", port))
}