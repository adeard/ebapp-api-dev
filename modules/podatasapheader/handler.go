package podatasapheader

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poDataSapHeaderHandler struct {
	poDataSapHeaderService Service
}

func NewPoDataSapHeaderHandler(v1 *gin.RouterGroup, poDataSapHeaderService Service) {
	handler := &poDataSapHeaderHandler{poDataSapHeaderService}

	poDataSapHeader := v1.Group("po_sap_header")

	poDataSapHeader.GET("/:id", handler.GetTitle)
}

func (h *poDataSapHeaderHandler) GetTitle(c *gin.Context) {
	id := c.Param("id")

	poDataSapHeaderTitle, err := h.poDataSapHeaderService.GetTitle(id)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Header tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Header",
		})
		return
	}

	response := domain.PoDataSapHeaderTitleResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Header",
		Data:    poDataSapHeaderTitle,
	}

	c.JSON(http.StatusOK, response)
}
