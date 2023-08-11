package poproject

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poProjectHandler struct {
	poProjectService Service
}

func NewPoProjectHandler(v1 *gin.RouterGroup, poProjectService Service) {
	handler := &poProjectHandler{poProjectService}

	poProject := v1.Group("po_project")

	poProject.GET("", handler.GetAll)
	poProject.GET("/:id", handler.GetByPo)
}

func (h *poProjectHandler) GetAll(c *gin.Context) {
	var input domain.PoProjectRequest

	poProject, err := h.poProjectService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data PO Project List",
		})
		return
	}

	response := domain.PoProjectResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data PO Project List",
		Data:    poProject,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectHandler) GetByPo(c *gin.Context) {
	po := c.Param("id")

	poProject, err := h.poProjectService.GetByPo(po)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Po Project tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Po Project",
		})
		return
	}

	response := domain.PoProjectResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Po Project",
		Data:    poProject,
	}

	c.JSON(http.StatusOK, response)
}
