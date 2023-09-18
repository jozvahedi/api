package main

import (
	"api/Utility"
	"api/config"
	"api/routing"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {

	err := config.SetConfig()
	if err != nil {
		log.Fatal(err)
	}

	//init server
	server := echo.New()
	server.Validator = &Utility.CustomValidator{Validator: validator.New()}
	//routing
	err = routing.SetRouting(server)
	if err != nil {
		log.Fatal(err)
	}
	//middleware
	//server.Use(middleware.Logger())
	server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			apiContext := &Utility.ApiContext{Context: c}
			return next(apiContext)
		}
	})

	//start server
	server.Start(":" + config.AppConfig.Server.Port)
}
