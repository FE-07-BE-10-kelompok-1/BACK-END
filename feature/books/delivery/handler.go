package delivery

import (
	"bookstore/config"
	"bookstore/domain"
	"bookstore/feature/common"
	"bookstore/feature/middlewares"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type bookHandler struct {
	bookUsecase domain.BookUsecase
}

func New(e *echo.Echo, bs domain.BookUsecase) {
	handler := &bookHandler{
		bookUsecase: bs,
	}
	useJWT := middleware.JWTWithConfig(middlewares.UseJWT([]byte(config.SECRET)))
	e.POST("/books", handler.AddBook(), useJWT)
	e.GET("/books", handler.GetAllBooks())
	e.GET("/books/:id", handler.GetSpecificBook())
	e.PUT("/books/:id", handler.UpdateBook(), useJWT)
	e.DELETE("/books/:id", handler.DeleteBook(), useJWT)
}

func (bh *bookHandler) AddBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		session := c.Get("session").(*session.Session)
		bucket := c.Get("bucket").(string)

		err := bh.bookUsecase.GetUser(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": err.Error(),
			})
		}

		var newBook BookRequest
		err = c.Bind(&newBook)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		err = validator.New().Struct(newBook)
		if err != nil {
			log.Println(err)
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		image, err := c.FormFile("image")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}
		file, err := c.FormFile("file")
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		imageUrl, fileUrl, err := bh.bookUsecase.UploadFiles(session, bucket, image, file)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}
		newBook.Image = imageUrl
		newBook.File = fileUrl

		data, err := bh.bookUsecase.AddBook(newBook.ToDomain())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(201, map[string]interface{}{
			"code":    201,
			"message": "success create new book",
			"data":    data,
		})
	}
}

func (bh *bookHandler) GetAllBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		search := c.QueryParam("search")

		data, err := bh.bookUsecase.GetAllBooks()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		var filteredData []domain.Book
		for i := 0; i < len(data); i++ {
			if strings.Contains(data[i].Title, search) || strings.Contains(data[i].Author, search) {
				filteredData = append(filteredData, data[i])
			}
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get all books data",
			"data":    filteredData,
		})
	}
}

func (bh *bookHandler) GetSpecificBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		data, err := bh.bookUsecase.GetSpecificBook(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success get book " + param,
			"data":    data,
		})
	}
}

func (bh *bookHandler) UpdateBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := bh.bookUsecase.GetUser(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": err.Error(),
			})
		}

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		var updatedData BookUpdateRequest
		err = c.Bind(&updatedData)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		data, err := bh.bookUsecase.UpdateBook(uint(id), updatedData.ToDomain())
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success update book " + param,
			"data":    data,
		})
	}
}

func (bh *bookHandler) DeleteBook() echo.HandlerFunc {
	return func(c echo.Context) error {
		err := bh.bookUsecase.GetUser(uint(common.ExtractData(c)))
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]interface{}{
				"code":    401,
				"message": err.Error(),
			})
		}

		param := c.Param("id")
		id, err := strconv.Atoi(param)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"code":    400,
				"message": err.Error(),
			})
		}

		err = bh.bookUsecase.DeleteBook(uint(id))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"code":    500,
				"message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    200,
			"message": "success delete book",
		})
	}
}
