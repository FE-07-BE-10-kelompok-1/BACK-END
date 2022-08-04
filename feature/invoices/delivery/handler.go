package delivery

import (
	"bookstore/config"
	"bookstore/domain"
	"bookstore/feature/common"
	"bookstore/feature/middlewares"
	"bookstore/infrastructure/payments/midtranspay"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

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
	e.GET("/orders", handler.GetOrders(), useJWT)
	e.POST("/orders", handler.MidtransCallback())
	e.POST("/orders/cancel", handler.CancelOrder())
}

func (ih *invoiceHandler) Checkout() echo.HandlerFunc {
	return func(c echo.Context) error {
		var reqBody CheckoutReq
		err := c.Bind(&reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}
		err = validator.New().Struct(reqBody)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    http.StatusBadRequest,
				"message": err.Error(),
			})
		}

		userData, err := ih.invoiceUsecase.GetUserData(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		data, err := ih.invoiceUsecase.CheckStocks(reqBody.Books_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		var total int
		for i := 0; i < len(data); i++ {
			total += data[i].Price
		}

		snapClient := c.Get("snapmidtrans").(snap.Client)
		orderToken := uuid.NewString()
		res, err := midtranspay.CreateTransactions(snapClient, orderToken, data, userData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		err = ih.invoiceUsecase.UpdateStock(reqBody.Books_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		invoice := domain.Invoice{
			ID:           orderToken,
			Users_ID:     userData.ID,
			Total:        total,
			Payment_Link: res.RedirectURL,
		}

		_, err = ih.invoiceUsecase.InsertInvoice(invoice, reqBody.Books_ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		err = ih.invoiceUsecase.DeleteCarts(userData.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":     http.StatusOK,
			"message":  "success checkout",
			"payments": res,
		})

	}
}

func (ih *invoiceHandler) GetOrders() echo.HandlerFunc {
	return func(c echo.Context) error {
		userData, err := ih.invoiceUsecase.GetUserData(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    http.StatusInternalServerError,
				"message": err.Error(),
			})
		}

		data, err := ih.invoiceUsecase.GetAllOrders(userData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get all orders",
			"data":    data,
		})
	}
}

func (ih *invoiceHandler) MidtransCallback() echo.HandlerFunc {
	return func(c echo.Context) error {
		var midtransReq MidtransCallbackRequest
		err := c.Bind(&midtransReq)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		var invoiceData domain.Invoice
		if midtransReq.Transaction_Status == "settlement" {
			invoiceData = domain.Invoice{Status: midtransReq.Transaction_Status, Payment_Method: midtransReq.Payment_Type, Paid_At: midtransReq.Settlement_Time}
			err = ih.invoiceUsecase.MidtransCallback(invoiceData, midtransReq.Order_ID)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		} else {
			invoiceData.Status = midtransReq.Transaction_Status
			err = ih.invoiceUsecase.MidtransCallback(invoiceData, midtransReq.Order_ID)
			if err != nil {
				log.Println(err)
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, "ok")
	}
}

func (ih *invoiceHandler) CancelOrder() echo.HandlerFunc {
	return func(c echo.Context) error {
		req, err := http.NewRequest(http.MethodGet, "https://server.athaprojects.me/books", nil)
		if err != nil {
			fmt.Printf("client: could not create request: %s\n", err)
			os.Exit(1)
		}

		req.Header.Add("Authorization", "Basic U0ItTWlkLXNlcnZlci1Zd0JIdzRZZTdNajEwVy1rNXVRTkJHTno6")

		res, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Printf("client: error making http request: %s\n", err)
			os.Exit(1)
		}

		fmt.Println(res.Header.Get("authorization"))

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("client: could not read response body: %s\n", err)
			os.Exit(1)
		}
		fmt.Printf("client: response body: %s\n", resBody)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":      200,
			"message":   "success cancel order",
			"serverKey": config.MIDTRANS_SERVER_KEY,
			"data":      resBody,
		})
	}
}
