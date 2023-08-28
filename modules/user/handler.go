package user

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userHandlerService Service
}

func NewUserHandler(v1 *gin.RouterGroup, userService Service) {
	handler := &userHandler{userService}

	user := v1.Group("user")

	user.GET("", handler.GetAll)
	user.GET("/:id", handler.GetById)
}

func (h *userHandler) GetAll(c *gin.Context) {
	var input domain.UserRequest

	users, err := h.userHandlerService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data User",
		})
		return
	}

	response := domain.UserResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data User",
		Data:    users,
	}

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) GetById(c *gin.Context) {
	id := c.Param("id")

	users, err := h.userHandlerService.GetByID(id)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data User tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data User ->" + err.Error(),
		})
		return
	}

	response := domain.UserResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data User",
		Data:    []domain.User{users},
	}

	c.JSON(http.StatusOK, response)
}
