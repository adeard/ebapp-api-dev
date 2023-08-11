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
	project := v1.Group("project")

	listProject.GET("", handler.GetAll)
	project.GET("/:id", handler.GetByID)
	project.POST("", handler.Store)
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

func (h *listProjectHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	project, err := h.listProjectService.GetByID(id)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data project tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data project",
		})
		return
	}

	response := domain.ListProjectsResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data BoQ Header",
		Data:    []domain.ListProject{project},
	}

	c.JSON(http.StatusOK, response)
}

func (h *listProjectHandler) Store(c *gin.Context) {
	var input domain.ListProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

}
