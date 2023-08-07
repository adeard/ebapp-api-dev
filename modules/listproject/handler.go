package listproject

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listProjectHandler struct {
	listProjectService Service
}

func NewListProjectHandler(v1 *gin.RouterGroup, listProjectService Service) {
	handler := &listProjectHandler{listProjectService}

	listProject := v1.Group("list_project")

	listProject.GET("", handler.GetAll)
}

func (h *listProjectHandler) GetAll(c *gin.Context) {
	var input domain.ListProjectRequest

	listProjects, err := h.listProjectService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Project List",
		})
		return
	}

	response := domain.ListProjectsResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data List Project",
		Data:    listProjects,
	}

	c.JSON(http.StatusOK, response)
}
