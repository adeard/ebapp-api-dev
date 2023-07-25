package boqheader

import (
	"ebapp-api-dev/domain"
	"net/http"
	"time"

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
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Header",
		})
		return
	}

	// Mengubah data BoQ Header menjadi response yang sesuai dengan BoqHeaderResponse.
	response := domain.BoqHeaderResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data BoQ Header",
		Data:    boqHeaders,
	}

	// Mengirimkan response dengan data BoQ Header yang sudah diubah formatnya.
	c.JSON(http.StatusOK, response)
}

func (h *boqHeaderHandler) Store(c *gin.Context) {
	var input domain.BoqHeaderRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createdBoqHeader := domain.BoqHeader{
		RunNum:            "14",
		BoqNo:             input.BoqNo,
		HeaderDescription: input.HeaderDescription,
		HeaderVersion:     input.HeaderVersion,
		HeaderStatus:      input.HeaderStatus,
		Created:           time.Now(),
		CreatedBy:         input.CreatedBy,
		LastUpdated:       time.Now(),
		LastUpdatedBy:     input.LastUpdatedBy,
		Category:          input.Category,
		Remarks:           input.Remarks,
	}

	boqHeaders, err := h.boqHeaderService.Store(createdBoqHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Header",
		})
		return
	}

	response := domain.BoqHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data BoQ Header",
		Data:    []domain.BoqHeader{boqHeaders},
	}

	c.JSON(http.StatusCreated, response)
}
