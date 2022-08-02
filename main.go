package main

import (
	"bookstore/config"
	"bookstore/factory"
	"bookstore/infrastructure/aws/s3"
	"bookstore/infrastructure/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)
	session := s3.ConnectAws(cfg)

	e := echo.New()
	factory.InitFactory(e, db)
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("session", session)
			c.Set("bucket", cfg.BUCKET_NAME)
			return next(c)
		}
	})

	fmt.Println("Menjalankan program...")
	dsn := fmt.Sprintf(":%d", config.SERVERPORT)
	e.Logger.Fatal(e.Start(dsn))
}
