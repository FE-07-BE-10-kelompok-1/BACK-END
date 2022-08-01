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
		err := c.Bind(&tmp)

		if err != nil {
			log.Println("Cannot parse data", err)
			c.JSON(http.StatusBadRequest, "error read input")
		}

		data, err := uh.userUsecase.AddUser(tmp.ToModel())

		if err != nil {
			log.Println("Cannot proces data", err)
			c.JSON(http.StatusInternalServerError, err)
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "success create data",
			"data":    data,
			"token":   common.GenerateToken(int(data.ID)),
		})
	}
}

func (uh *userHandler) LoginHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userLogin LoginFormat
		err := c.Bind(&userLogin)
		if err != nil {
			log.Println("Cannot parse data", err)
			return c.JSON(http.StatusBadRequest, "cannot read input")
		}
		row, data, e := uh.userUsecase.LoginUser(userLogin.LoginToModel())
		if e != nil {
			log.Println("Cannot proces data", err)
			return c.JSON(http.StatusBadRequest, "username or password incorrect")
		}
		if row == -1 {
			return c.JSON(http.StatusBadRequest, "username or password incorrect")
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success login",
			"token":   common.GenerateToken(int(data.ID)),
		})
	}
}

func (uh *userHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {

		var tmp InsertFormat
		result := c.Bind(&tmp)

		qry := map[string]interface{}{}
		idUpdate := common.ExtractData(c)

		if result != nil {
			log.Println(result, "Cannot parse input to object")
			return c.JSON(http.StatusInternalServerError, "Error read update")
		}

		if tmp.Fullname != "" {
			qry["fullname"] = tmp.Fullname
		}
		if tmp.Username != "" {
			qry["username"] = tmp.Username
		}
		if tmp.Phone != "" {
			qry["phone"] = tmp.Phone
		}
		if tmp.Password != "" {
			qry["password"] = tmp.Password
		}
		data, err := uh.userUsecase.UpdateUser(idUpdate, tmp.ToModel())

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "cannot update")
		}

		return c.JSON(http.StatusCreated, map[string]interface{}{
			"message": "Success update data",
			"data":    data,
		})
	}
}

func (uh *userHandler) GetProfile() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := common.ExtractData(c)

		data, err := uh.userUsecase.GetProfile(id)

		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				return c.JSON(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}
		return c.JSON(http.StatusFound, map[string]interface{}{
			"message": "data found",
			"data":    data,
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
			return c.JSON(http.StatusInternalServerError, "cannot delete user")
		}
		return c.JSON(http.StatusOK, "success to delete user")
	}
}
