package poprojectaddendum

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poProjectAddendumHandler struct {
	poProjectAddendumService Service
}

func NewPoProjectAddendumHandler(v1 *gin.RouterGroup, poProjectAddendumService Service) {
	handler := &poProjectAddendumHandler{poProjectAddendumService}

	poProjectAddendum := v1.Group("po_addendum")

	poProjectAddendum.GET("/:id/:var1/:var2/:var3", handler.GetByPo)
	poProjectAddendum.DELETE("/:id/:var1/:var2/:var3/:var4/:var5", handler.Delete)
	poProjectAddendum.POST("", handler.Store)
}

func (h *poProjectAddendumHandler) GetByPo(c *gin.Context) {
	po := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	addon := string("/")

	poProjectAddendum, err := h.poProjectAddendumService.GetByPo(po + addon + var1 + addon + var2 + addon + var3)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Addendum tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Addendum",
		})
		return
	}

	response := domain.PoProjectAddendumResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Addendum",
		Data:    poProjectAddendum,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectAddendumHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3
	po := c.Param("var4")
	item := c.Param("var5")

	if deleteErr := h.poProjectAddendumService.Delete(FinalId, po, item); deleteErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menghapus data",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Data berhasil dihapus",
		"data":    nil,
	})
}

func (h *poProjectAddendumHandler) Store(c *gin.Context) {
	var input []domain.PoProjectAddendumRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	// Iterasi melalui array input dan menyimpan setiap entri ke dalam basis data
	for _, item := range input {
		createPoProject := domain.PoProjectAddendum{
			Po:          item.Po,
			PekerjaanNo: item.PekerjaanNo,
			Item:        item.Item,
			Description: item.Description,
			Qty:         item.Qty,
			PoUnit:      item.PoUnit,
			Price:       item.Price,
			Currency:    item.Currency,
			Wbs:         item.Wbs,
			Cera:        item.Cera,
			Order:       item.Order,
		}

		// Simpan setiap entri ke dalam basis data
		if _, err := h.poProjectAddendumService.Store(createPoProject); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data Addendum",
			})
			return
		}
	}

	response := domain.PoProjectAddendumResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Addendum",
	}

	c.JSON(http.StatusCreated, response)
}
