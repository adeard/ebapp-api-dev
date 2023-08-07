package boqheader

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/helper"
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
	boqHeader.GET("/:id", handler.GetByID)
	boqHeader.POST("", handler.Store)
	boqHeader.PUT("/:id", handler.Update)
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

func (h *boqHeaderHandler) GetByID(c *gin.Context) {
	// Mendapatkan nilai ID dari parameter URL.
	id := c.Param("id")

	// Menggunakan boqHeaderService untuk mendapatkan data BoQ Header berdasarkan ID dari repository.
	boqHeader, err := h.boqHeaderService.GetByID(id)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data BoQ Header tidak ditemukan",
			})
			return
		}

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
		Data:    []domain.BoqHeader{boqHeader}, // Menggunakan slice dari BoqHeader
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

	boqHeaderID, _ := helper.GenerateHeaderBoq(3)

	createdBoqHeader := domain.BoqHeader{
		RunNum:            boqHeaderID,
		BoqNo:             input.BoqNo,
		HeaderDescription: input.HeaderDescription,
		HeaderVersion:     input.HeaderVersion,
		HeaderStatus:      input.HeaderStatus,
		Created:           time.Now(),
		CreatedBy:         input.CreatedBy,
		Category:          input.Category,
		Remarks:           input.Remarks,
	}

	boqHeaders, err := h.boqHeaderService.Store(createdBoqHeader)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data BoQ Header",
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

func (h *boqHeaderHandler) Update(c *gin.Context) {
	// Mendapatkan nilai ID dari parameter URL.
	id := c.Param("id")

	var input domain.BoqHeaderRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	// Membuat objek domain.BoqHeader yang akan diupdate.
	updatedBoqHeader := domain.BoqHeader{
		RunNum:            id,
		BoqNo:             input.BoqNo,
		HeaderDescription: input.HeaderDescription,
		HeaderVersion:     input.HeaderVersion,
		HeaderStatus:      input.HeaderStatus,
		LastUpdated:       time.Now(),
		LastUpdatedBy:     input.LastUpdatedBy,
		Category:          input.Category,
		Remarks:           input.Remarks,
	}

	// Memanggil service untuk melakukan update data BoQ Header.
	_, err := h.boqHeaderService.Update(updatedBoqHeader, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengupdate data BoQ Header",
		})
		return
	}

	response := domain.BoqHeaderResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengupdate data BoQ Header",
	}

	c.JSON(http.StatusOK, response)
}
