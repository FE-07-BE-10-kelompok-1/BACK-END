package delivery

import (
	"bookstore/domain"
	"fmt"

	"github.com/labstack/echo/v4"
)

type cartHandler struct {
	cartUsecase domain.CartUsecase
}

func New(e *echo.Echo, cs domain.CartUsecase) {
	handler := &cartHandler{
		cartUsecase: cs,
	}
	fmt.Println(handler)
}
