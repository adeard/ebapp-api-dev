package auth

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authService Service
}

func NewAuthHandler(v1 *gin.RouterGroup, authService Service) {

	handler := &authHandler{authService}

	auth := v1.Group("auth")

	auth.POST("sign_in", handler.SignIn)

	auth.Use(middlewares.AuthService_Sample())

	auth.GET("", handler.GetAuth)
}

func (h *authHandler) GetAuth(c *gin.Context) {
	auth, err := h.authService.AuthTest()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": auth,
	})
}

func (h *authHandler) SignIn(c *gin.Context) {
	var authrequest domain.AuthRequest

	err := c.ShouldBindJSON(&authrequest)
	if err != nil {

		errorMessages := []string{}

		for _, v := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s , condition : %s", v.Field(), v.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": errorMessages,
		})

		return
	}

	auth, err := h.authService.SignIn(authrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": auth,
	})
}
