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

	poDataSapHeader.GET("/:id", handler.GetData)
}

func (h *poDataSapHeaderHandler) GetData(c *gin.Context) {
	po := c.Param("id")

	poDataSapHeader, err := h.poDataSapHeaderService.GetByPo(po)
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
			"message": "Gagal mengambil data Header" + err.Error(),
		})
		return
	}

	response := domain.PoDataSapHeaderResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Header",
		Data:    poDataSapHeader,
	}

	c.JSON(http.StatusOK, response)
}
