package popic

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poPicHandler struct {
	poPicService Service
}

func NewPoPicHandler(v1 *gin.RouterGroup, poPicService Service) {
	handler := &poPicHandler{poPicService}

	poPic := v1.Group("pic")

	poPic.GET("/:id/:var1/:var2/:var3", handler.GetByRunNum)
	poPic.POST("", handler.Store)
}

func (h *poPicHandler) GetByRunNum(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	addons := "/"

	pic, err := h.poPicService.FindPicByPo(id + addons + var1 + addons + var2 + addons + var3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data BoQ Body",
		"data":    pic,
	})
}

func (h *poPicHandler) Store(c *gin.Context) {
	var input []domain.PoPicRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	var addedEntries []domain.PoPic // Untuk menyimpan item yang berhasil ditambahkan

	for _, item := range input {
		// Panggil fungsi untuk memeriksa apakah entri sudah ada
		existingEntry, err := h.poPicService.FindPic(item.Uid, item.RunNum, item.Role)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusNotFound,
				"message": "Terdeteksi data baru! mencoba memasukan data baru uid : " + item.Uid + " ke Nomor Pekerjaan : " + item.RunNum,
			})
		}

		// Jika entri tidak ditemukan, tambahkan
		if existingEntry.Name == "" {
			setPIC := domain.PoPic{
				RunNum: item.RunNum,
				Id:     item.Id,
				Uid:    item.Uid,
				Name:   item.Name,
				Email:  item.Email,
				Role:   item.Role,
				Status: item.Status,
			}

			// Simpan setiap entri ke dalam basis data
			if _, err := h.poPicService.Store(setPIC); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  http.StatusInternalServerError,
					"message": "Gagal meneruskan data PIC",
				})
				return
			}

			addedEntries = append(addedEntries, setPIC) // Tambahkan item yang berhasil ditambahkan ke daftar
		} else if existingEntry.Name != "" {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal memasukan data baru",
			})
		}
	}

	response := domain.PoPicResponse{
		Status:  http.StatusCreated,
		Message: "Berikut adalah data PIC yang berhasil disimpan",
		Data:    addedEntries, // Mengembalikan daftar item yang berhasil ditambahkan
	}

	c.JSON(http.StatusCreated, response)
}
