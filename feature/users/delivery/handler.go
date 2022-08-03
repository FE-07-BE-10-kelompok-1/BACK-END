package delivery

import (
	"bookstore/domain"
	"bookstore/feature/common"
	"log"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userUsecase domain.UserUsecase
}

func New(us domain.UserUsecase) domain.UserHandler {
	return &userHandler{
		userUsecase: us,
	}
}

func (uh *userHandler) InsertUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var tmp InsertFormat
		errBind := c.Bind(&tmp)

		if errBind != nil {
			log.Println("Cannot parse data", errBind)
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}

		_, err := uh.userUsecase.AddUser(tmp.ToModel())
		if err != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"code":    201,
			"message": "success operation",
		})
	}
}

func (uh *userHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		err := c.Bind(&userLogin)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "cannot read input",
			})
		}
		row, data, e := uh.userUsecase.LoginUser(userLogin.LoginToModel())
		if e != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "username or password incorrect",
			})
		}
		if row == -1 {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": "cannot read input",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"token":   common.GenerateToken(int(data.ID)),
			"role":    data.Role,
			"message": "success login",
		})
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var tmp UpdateFormat
		result := c.Bind(&tmp)
		idUpdate := common.ExtractData(c)
		if result != nil {
			log.Println(result, "Cannot parse input to object")
			return c.JSON(http.StatusInternalServerError, "Error read update")
		}

		_, err := uh.userUsecase.UpdateUser(idUpdate, tmp.UpdateToModel())

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success operation",
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, map[string]interface{}{
					"code":    400,
					"message": "data not found",
				})
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"code":     200,
			"id":       data.ID,
			"fullname": data.Fullname,
			"username": data.Username,
			"phone":    data.Phone,
			"password": data.Password,
			"message":  "Success Operation",
		})
	}
}

func (uh *userHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)
		if id == 0 {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}
		_, errDel := uh.userUsecase.DeleteUser(id)
		if errDel != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": "internal server error",
			})
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "Success Operation",
		})
	}
}
