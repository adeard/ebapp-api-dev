package main

import (
	"ebapp-api-dev/config"
	"ebapp-api-dev/modules/auth"
	"ebapp-api-dev/modules/boqbody"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	db := config.Connect()

	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("api/v1")
	auth.NewAuthHandler(v1, auth.AuthRegistry(db))

	//v1.Use(middlewares.AuthService_Sample())

	boqbody.NewBoqBodyHandler(v1, boqbody.BoqBodyRegistry(db))

	router.Run(":88")
}
