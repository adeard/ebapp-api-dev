package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Middleware untuk mencatat log setiap permintaan
func RequestLoggerMiddleware(c *gin.Context) {
	c.Next()

	// Cek apakah permintaan berhasil (kode status 2xx atau 3xx)
	if c.Writer.Status() < 300 {
		log.Printf("Request %s %s Success - Status: %d", c.Request.Method, c.Request.URL, c.Writer.Status())
	} else {
		log.Printf("Request %s %s Fail - Status: %d", c.Request.Method, c.Request.URL, c.Writer.Status())
	}
}
