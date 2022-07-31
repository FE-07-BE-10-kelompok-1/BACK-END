package delivery

import (
	"bookstore/domain"
	"fmt"

	"github.com/labstack/echo/v4"
)

type invoiceHandler struct {
	invoiceUsecase domain.InvoiceUsecase
}

func New(e *echo.Echo, is domain.InvoiceUsecase) {
	handler := &invoiceHandler{
		invoiceUsecase: is,
	}
	fmt.Println(handler)
}
