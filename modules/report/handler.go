package report

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type reportHandler struct {
	reportService Service
}

func NewReportHandler(v1 *gin.RouterGroup, reportService Service) {
	handler := &reportHandler{reportService}

	report := v1.Group("report")
	report.Use(middlewares.AuthService())

	report.GET("/plant", handler.GetByPlant)
}

func (h *reportHandler) GetByPlant(c *gin.Context) {
	idsJSON := c.GetHeader("Ids")

	var idSlice []string
	err := json.Unmarshal([]byte(idsJSON), &idSlice)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Gagal memproses data ID",
		})
		return
	}

	// Memanggil service untuk mendapatkan data dengan menggunakan IDs yang diberikan
	listReport, err := h.reportService.GetByPlant(idSlice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Project List",
		})
		return
	}

	response := domain.ListReportResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data List Project",
		Data:    listReport,
	}

	c.JSON(http.StatusOK, response)
}
