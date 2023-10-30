package poboqheader

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poBoqHeaderHandler struct {
	poBoqHeaderService Service
}

func NewPoBoqHeaderHandler(v1 *gin.RouterGroup, poBoqHeaderService Service) {
	handler := &poBoqHeaderHandler{poBoqHeaderService}

	hHeader := v1.Group("poboq_header")

	hHeader.GET("/:id/:var1/:var2/:var3", handler.GetByPekerjaanNo)
	hHeader.DELETE("/:id/:var1/:var2/:var3/:var4/:var5", handler.Delete)
	hHeader.POST("", handler.Store)

}

func (h *poBoqHeaderHandler) GetByPekerjaanNo(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")

	headers, err := h.poBoqHeaderService.GetByPekerjaanNo(id + "/" + var1 + "/" + var2 + "/" + var3)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Header tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Header",
		})
		return
	}

	response := domain.PoBoqHeaderResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Header",
		Data:    headers,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poBoqHeaderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3
	po := c.Param("var4")
	order := c.Param("var5")

	if deleteErr := h.poBoqHeaderService.Delete(FinalId, po, order); deleteErr != nil {
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

func (h *poBoqHeaderHandler) Store(c *gin.Context) {
	var input []domain.PoBoqHeader

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	for _, item := range input {
		createHeader := domain.PoBoqHeader{
			PekerjaanNo: item.PekerjaanNo,
			Po:          item.Po,
			Item:        item.Item,
			Description: item.Description,
			Qty:         item.Qty,
			Unit:        item.Unit,
			Price:       item.Price,
			Currency:    item.Currency,
			Order:       item.Order,
			IsAddendum:  item.IsAddendum,
		}

		if _, err := h.poBoqHeaderService.Store(createHeader); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data header",
			})
			return
		}
	}

	response := domain.PoBoqHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data header",
	}

	c.JSON(http.StatusCreated, response)
}
