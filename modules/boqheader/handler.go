package boqheader

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type boqHeaderHandler struct {
	boqHeaderService Service
}

func NewBoqHeaderHandler(v1 *gin.RouterGroup, boqHeaderService Service) {
	handler := &boqHeaderHandler{boqHeaderService}

	boqHeader := v1.Group("boq_header")

	boqHeader.GET("", handler.GetAll)
	boqHeader.POST("", handler.Store)
}

func (h *boqHeaderHandler) GetAll(c *gin.Context) {
	var input domain.BoqHeaderRequest
	// Menggunakan boqHeaderService untuk mendapatkan data BoQ Header dari repository.
	boqHeaders, err := h.boqHeaderService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Gagal mengambil data BoQ Header",
		})
		return
	}

	// Mengubah data BoQ Header menjadi response yang sesuai dengan BoqHeaderResponse.
	var response domain.BoqHeaderResponse

	response.Data = boqHeaders
	response.Status = 200
	response.Message = "Berhasil mengambil data BoQ Header"

	// Mengirimkan response dengan data BoQ Header yang sudah diubah formatnya.
	c.JSON(http.StatusOK, response)
}

func (h *boqHeaderHandler) Store(c *gin.Context) {}
