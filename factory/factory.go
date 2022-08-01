package factory

import (
	usersData "bookstore/feature/users/data"
	usersDelivery "bookstore/feature/users/delivery"
	usersUsecase "bookstore/feature/users/usecase"

	booksData "bookstore/feature/books/data"
	booksDelivery "bookstore/feature/books/delivery"
	booksUsecase "bookstore/feature/books/usecase"

	cartsData "bookstore/feature/carts/data"
	cartsDelivery "bookstore/feature/carts/delivery"
	cartsUsecase "bookstore/feature/carts/usecase"

	invoicesData "bookstore/feature/invoices/data"
	invoicesDelivery "bookstore/feature/invoices/delivery"
	invoicesUsecase "bookstore/feature/invoices/usecase"

	ordersData "bookstore/feature/orders/data"
	ordersDelivery "bookstore/feature/orders/delivery"
	ordersUsecase "bookstore/feature/orders/usecase"

	route "bookstore/route"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitFactory(e *echo.Echo, db *gorm.DB) {
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	e.Use(middleware.CORS())

	userData := usersData.New(db)
	validator := validator.New()
	useCase := usersUsecase.New(userData, validator)
	userHandler := usersDelivery.New(useCase)
	route.RouteUser(e, userHandler)

	bookData := booksData.New(db)
	BookuseCase := booksUsecase.New(bookData)
	booksDelivery.New(e, BookuseCase)

	cartData := cartsData.New(db)
	CartuseCase := cartsUsecase.New(cartData)
	cartsDelivery.New(e, CartuseCase)

	invoiceData := invoicesData.New(db)
	InvoiceuseCase := invoicesUsecase.New(invoiceData)
	invoicesDelivery.New(e, InvoiceuseCase)

	orderData := ordersData.New(db)
	OrderuseCase := ordersUsecase.New(orderData)
	ordersDelivery.New(e, OrderuseCase)
}
