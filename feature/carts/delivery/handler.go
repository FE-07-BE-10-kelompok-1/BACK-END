package delivery

import (
	"bookstore/config"
	"bookstore/domain"
	"bookstore/feature/common"
	"bookstore/feature/middlewares"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type cartHandler struct {
	cartUsecase domain.CartUsecase
}

func New(e *echo.Echo, cs domain.CartUsecase) {
	handler := &cartHandler{
		cartUsecase: cs,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/carts", handler.AddToCart(), useJWT)
	e.GET("/carts", handler.GetCarts(), useJWT)
	e.DELETE("/carts/:id", handler.DeleteFromCart(), useJWT)
}

func (ch *cartHandler) AddToCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := common.ExtractData(c)
		var addToCart AddToCart
		err := c.Bind(&addToCart)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		data := domain.Cart{Books_ID: addToCart.Books_ID, Users_ID: uint(userID)}

		err = ch.cartUsecase.AddToCart(data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success add to cart",
		})
	}
}

func (ch *cartHandler) GetCarts() echo.HandlerFunc {
	return func(c echo.Context) error {
		usersID := common.ExtractData(c)
		data, err := ch.cartUsecase.GetCarts(uint(usersID))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		cartResponse, total := ToCartsResponse(data)

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get you carts",
			"data":    cartResponse,
			"total":   total,
		})
	}
}

func (ch *cartHandler) DeleteFromCart() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		data := domain.Cart{ID: uint(id), Users_ID: uint(common.ExtractData(c))}
		err = ch.cartUsecase.DeleteFromCart(data)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success delete item from cart",
		})
	}
}
