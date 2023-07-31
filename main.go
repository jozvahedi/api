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

	//start server
	server.Start(":" + config.AppConfig.Server.Port)
}
