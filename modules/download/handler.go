package download

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type downloadHandler struct {
	downloadService Service
}

func NewDownlaodHandler(v1 *gin.RouterGroup, downloadService Service) {
	handler := &downloadHandler{downloadService}

	download := v1.Group("download")
	download.GET("", handler.DownloadAttachment)
}

func (h *downloadHandler) DownloadAttachment(c *gin.Context) {
	filePath := c.Query("file")
	// Check if filePath is valid

	// Set Content-Disposition header untuk menentukan nama file saat didownload
	c.Header("Content-Disposition", "attachment; filename="+filepath.Base(filePath))

	// Set Content-Type header untuk menentukan tipe konten file
	c.Header("Content-Type", "application/octet-stream")

	// Mengirimkan file sebagai response
	c.File(filePath)
}
