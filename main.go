package main

import (
	"ebapp-api-dev/config"
	"ebapp-api-dev/modules/auth"
	"ebapp-api-dev/modules/boqbody"
	"log"
	"net/http"
	"os"

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

	// Mengatur mode GIN menjadi release
	gin.SetMode(gin.ReleaseMode)

	//Penyesuaian Port ke IIS
	port := "88"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		log.Fatal("Koneksi gagal -> port "+port+":", err)
	}
}
