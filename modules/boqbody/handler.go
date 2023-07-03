package boqbody

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type boqBodyHandler struct {
	boqBodyService Service
}

func NewBoqBodyHandler(v1 *gin.RouterGroup, boqBodyService Service) {

	handler := &boqBodyHandler{boqBodyService}

	ztsTravis := v1.Group("boq_body")

	ztsTravis.GET("", handler.GetAll)
	ztsTravis.POST("", handler.Store)
}

func (h *boqBodyHandler) GetAll(c *gin.Context) {
	var input domain.BoqBodyRequest

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ztsTravis, err := h.boqBodyService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}

func (h *boqBodyHandler) Store(c *gin.Context) {
	var input domain.BoqBodyRequest

	c.ShouldBindJSON(&input)

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errorMessages := []interface{}{}

		for _, v := range err.(validator.ValidationErrors) {
			errorArray := map[string]string{
				"field":   v.Field(),
				"message": v.ActualTag(),
			}

			errorMessages = append(errorMessages, errorArray)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	ztsTravis, err := h.boqBodyService.Store(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}
