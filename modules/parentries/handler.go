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
	parEntries.GET("/:id", handler.GetByID)
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

func (h *parEntriesHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	parEntries, err := h.parEntriesService.GetByID(id)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Par Entries tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Par Entries",
		})
		return
	}

	response := domain.ParEntriesResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Par Entries",
		Data:    parEntries,
	}

	c.JSON(http.StatusOK, response)
}
