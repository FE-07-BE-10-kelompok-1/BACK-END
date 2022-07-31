package delivery

import (
	"bookstore/domain"
	"fmt"

	"github.com/labstack/echo/v4"
)

type bookHandler struct {
	bookUsecase domain.BookUsecase
}

func New(e *echo.Echo, bs domain.BookUsecase) {
	handler := &bookHandler{
		bookUsecase: bs,
	}
	fmt.Println(handler)
}
