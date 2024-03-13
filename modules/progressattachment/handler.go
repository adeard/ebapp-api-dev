package progressattachment

import (
	"ebapp-api-dev/domain"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type progressAttachmentHandler struct {
	progressAttachmentService Service
}

func NewProgressAttachmentHandler(v1 *gin.RouterGroup, progressAttachmentService Service) {
	handler := &progressAttachmentHandler{progressAttachmentService}

	attachment := v1.Group("progress_attachment")

	attachment.POST("", handler.Store)
}

func (h *progressAttachmentHandler) Store(c *gin.Context) {
	var input domain.ProgressAttachment

	// Mengambil nilai dari form
	input.RunNum = c.PostForm("run_num")
	input.UploadBy = c.PostForm("upload_by")
	file, err := c.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	// Mengecek ukuran file
	maxFileSize := int64(2 << 20) // 2MB
	if file.Size > maxFileSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Ukuran file melebihi batas maksimum (2MB)",
		})
		return
	}

	// Menetapkan tahun saat ini
	currentYear := time.Now().Year()

	//mengambil nama folder yang ditetapkan
	attachmentFolder := os.Getenv("ATTACHMENT_FOLDER_NAME")

	// Menetapkan direktori UPLOAD/tahun_saat_ini
	uploadDir := filepath.Join(attachmentFolder, strconv.Itoa(currentYear))

	// Memeriksa apakah folder UPLOAD/tahun_saat_ini sudah ada
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		// Jika folder tidak ada, buat folder UPLOAD/tahun_saat_ini
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal membuat folder tahun_saat_ini di dalam UPLOAD",
			})
			return
		}
	}

	// Menetapkan nama baru untuk file
	timestamp := time.Now().Format("02-01-2006") // Format tanggal dd-MM-yyyy
	newFileName := input.RunNum + "_" + timestamp + "_" + file.Filename

	// Membuat path lengkap untuk menyimpan file
	filePath := filepath.Join(uploadDir, newFileName)

	// Menetapkan waktu saat ini sebagai tanggal unggah
	input.Date = time.Now()

	// Menetapkan file name
	input.FileName = file.Filename

	// Simpan file ke direktori yang ditentukan
	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menyimpan file",
		})
		return
	}

	// Menyimpan nama baru file ke dalam kolom FilePath
	input.FilePath = filepath.Join(strconv.Itoa(currentYear), newFileName)

	// Menyimpan informasi attachment ke dalam database
	data, err := h.progressAttachmentService.Store(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan Attachment",
		})
		return
	}

	response := domain.ProgressAttachmentResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan Attachment",
		Data:    []domain.ProgressAttachment{data},
	}

	c.JSON(http.StatusCreated, response)
}
