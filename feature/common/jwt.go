package common

import (
	"bookstore/config"
	"log"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func GenerateToken(ID int) string {
	info := jwt.MapClaims{}
	info["ID"] = ID
	auth := jwt.NewWithClaims(jwt.SigningMethodHS256, info)
	token, err := auth.SignedString([]byte(config.SECRET))
	if err != nil {
		log.Fatal("cannot generate key")
		return ""
	}

	return token
}

func ExtractData(c echo.Context) int {
	head := c.Request().Header
	token := strings.Split(head.Get("Authorization"), " ")

	res, _ := jwt.Parse(token[len(token)-1], func(t *jwt.Token) (interface{}, error) {
		return []byte(config.SECRET), nil
	})

	if res.Valid {
		resClaim := res.Claims.(jwt.MapClaims)
		parseID := resClaim["ID"].(float64)
		return int(parseID)
	}

	return -1
}
