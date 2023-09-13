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
	poDataSapHeader.GET("/wbs/:id", handler.GetWbs)
	poDataSapHeader.GET("/area/:id", handler.GetArea)
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
			"message": "Gagal mengambil data Header => error dari sini : " + err.Error(),
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

func (h *poDataSapHeaderHandler) GetWbs(c *gin.Context) {
	id := c.Param("id")

	poDataSapHeaderWbs, err := h.poDataSapHeaderService.GetWbs(id)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data WBS tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data WBS",
		})
		return
	}

	response := domain.PoDataSapHeaderWbsResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Header",
		Data:    poDataSapHeaderWbs,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poDataSapHeaderHandler) GetArea(c *gin.Context) {
	id := c.Param("id")

	dataMasterPlant, err := h.poDataSapHeaderService.GetArea(id)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Master Plant tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil Data Master Plant",
		})
		return
	}

	response := domain.DataMasterPlantResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil Data Master Plant",
		Data:    []domain.DataMasterPlant{dataMasterPlant},
	}

	c.JSON(http.StatusOK, response)
}
