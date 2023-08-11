package main

import (
	"ebapp-api-dev/config"
	"ebapp-api-dev/middlewares"
	"ebapp-api-dev/modules/auth"
	"ebapp-api-dev/modules/boqbody"
	"ebapp-api-dev/modules/boqheader"
	"ebapp-api-dev/modules/listproject"
	"ebapp-api-dev/modules/parentries"
	"ebapp-api-dev/modules/poproject"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	// Buka file log.txt untuk ditulis (create or append)
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Gagal membuka file log.txt: %s", err)
	}
	defer file.Close()

	// Pengaturan log output ke file log.txt
	log.SetOutput(file)

	// Menampilkan log saat aplikasi dimulai
	log.Println("Start App Service...")

	db := config.Connect()

	router := gin.Default()
	router.Use(cors.AllowAll())

	v1 := router.Group("api/v1")
	auth.NewAuthHandler(v1, auth.AuthRegistry(db))

	// Menambahkan middleware untuk mencatat log setiap permintaan
	v1.Use(middlewares.RequestLoggerMiddleware)

	boqbody.NewBoqBodyHandler(v1, boqbody.BoqBodyRegistry(db))
	boqheader.NewBoqHeaderHandler(v1, boqheader.BoqHeaderRegistry(db))
	listproject.NewListProjectHandler(v1, listproject.ListProjectRegistry(db))
	parentries.NewParEntriesHandler(v1, parentries.ParEntriesRegistry(db))
	poproject.NewPoProjectHandler(v1, poproject.PoProjectRegistry(db))

	// Mengatur mode GIN menjadi release
	gin.SetMode(gin.ReleaseMode)

	//Penyesuaian Port ke IIS
	port := "88"
	if os.Getenv("ASPNETCORE_PORT") != "" {
		port = os.Getenv("ASPNETCORE_PORT")
	}

	// Menampilkan log koneksi sukses
	log.Println("App Service run in port:", port)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		// Menampilkan log ketika koneksi gagal
		log.Fatal("Connection Fail -> port "+port+":", err)
	}
}
