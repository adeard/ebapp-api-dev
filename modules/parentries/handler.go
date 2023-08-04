package parentries

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type parEntriesHandler struct {
	parEntriesService Service
}

func NewParEntriesHandler(v1 *gin.RouterGroup, parEntriesService Service) {
	handler := &parEntriesHandler{parEntriesService}

	parEntries := v1.Group("par_entries")

	parEntries.GET("", handler.GetAll)
}

func (h *parEntriesHandler) GetAll(c *gin.Context) {
	var input domain.ParEntriesRequest
	// Menggunakan parEntriesService untuk mendapatkan data ParEntries dari repository.
	parEntriess, err := h.parEntriesService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data ParEntries",
		})
		return
	}

	// Mengubah data ParEntries menjadi response yang sesuai dengan ParEntriesResponse.
	response := domain.ParEntriesResponse{
		Data:    parEntriess,
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data ParEntries",
	}

	// Mengirimkan response dengan data ParEntries yang sudah diubah formatnya.
	c.JSON(http.StatusOK, response)
}
