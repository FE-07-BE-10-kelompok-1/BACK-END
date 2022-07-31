package delivery

import (
	"bookstore/domain"
	"fmt"

	"github.com/labstack/echo/v4"
)

type orderHandler struct {
	orderUsecase domain.OrderUsecase
}

func New(e *echo.Echo, os domain.OrderUsecase) {
	handler := &orderHandler{
		orderUsecase: os,
	}
	fmt.Println(handler)
}
