package auth

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type authHandler struct {
	authService Service
}

func NewAuthHandler(v1 *gin.RouterGroup, authService Service) {

	handler := &authHandler{authService}

	auth := v1.Group("auth")

	//auth.POST("sign_in", handler.SignIn)

	//auth.Use(middlewares.AuthService_Sample())

	//auth.GET("", handler.GetAuth)
	auth.POST("get_token", handler.GetToken)
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

func (h *authHandler) GetToken(c *gin.Context) {
	var authtoken domain.AuthToken

	if err := c.ShouldBindJSON(&authtoken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}
	_, err := middlewares.DecryptAndValidate(authtoken.Token, middlewares.JWTKEY)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": err.Error(),
		})
		return
	}

	cookie := http.Cookie{
		Name:     "session_token",
		Value:    url.QueryEscape(authtoken.Token),
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
		Secure:   false,
		Path:     "/",
		Domain:   os.Getenv("URL_APP"),
	}
	http.SetCookie(c.Writer, &cookie)
	if retrievedCookie, err := c.Cookie("session_token"); err == nil {
		c.JSON(http.StatusOK, gin.H{
			"token":   true,
			"message": "Cookie berhasil di-set",
			"cookie":  retrievedCookie,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"token":   false,
			"message": err.Error(),
		})
	}
}
