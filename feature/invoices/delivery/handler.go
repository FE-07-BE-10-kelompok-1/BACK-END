package delivery

import (
	"bookstore/config"
	"bookstore/domain"
	"bookstore/feature/common"
	"bookstore/feature/middlewares"
	"bookstore/infrastructure/payments/midtranspay"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/midtrans/midtrans-go/snap"
)

type invoiceHandler struct {
	invoiceUsecase domain.InvoiceUsecase
}

func New(e *echo.Echo, is domain.InvoiceUsecase) {
	handler := &invoiceHandler{
		invoiceUsecase: is,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/checkout", handler.Checkout(), useJWT)
}

func (ih *invoiceHandler) Checkout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody CheckoutReq
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		err = validator.New().Struct(reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		userData, err := ih.invoiceUsecase.GetUserData(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		data, err := ih.invoiceUsecase.CheckStocks(reqBody.Books_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		var total int
		for i := 0; i < len(data); i++ {
			total += data[i].Price
		}

		snapClient := c.Get("snapmidtrans").(snap.Client)
		orderToken := uuid.NewString()
		res, err := midtranspay.CreateTransactions(snapClient, orderToken, data, userData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		err = ih.invoiceUsecase.UpdateStock(reqBody.Books_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		invoice := domain.Invoice{
			ID:           orderToken,
			Users_ID:     userData.ID,
			Total:        total,
			Payment_Link: res.RedirectURL,
		}

		_, err = ih.invoiceUsecase.InsertInvoice(invoice, reqBody.Books_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, res)

	}
}
