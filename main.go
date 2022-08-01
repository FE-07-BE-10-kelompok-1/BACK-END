package main

import (
	"bookstore/config"
	"bookstore/factory"
	"bookstore/infrastructure/database/mysql"
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	mysql.MigrateData(db)

	fmt.Println(cfg)
	e := echo.New()
	factory.InitFactory(e, db)

	fmt.Println("Menjalankan program...")
	dsn := fmt.Sprintf(":%d", 8000)
	e.Logger.Fatal(e.Start(dsn))
}
