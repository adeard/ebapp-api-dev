package poboqheader

import (
	"ebapp-api-dev/domain"
	"net/http"
	"strconv"

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
	hHeader.POST("sync/price", handler.SyncPrice)
	hHeader.GET("get/:id/:var1/:var2/:var3", handler.GetByPekerjaanNoWithBody)
	hHeader.GET("get/:id/:var1/:var2/:var3/:var4", handler.GetByPekerjaanNoWithBodyByOrder)

	hHeader.GET("getss/:id/:var1/:var2/:var3", handler.GetByPekerjaanNoWithBodyServerSide)
	hHeader.GET("getsscount/:id/:var1/:var2/:var3", handler.GetByPekerjaanNoWithBodyServerSideCount)

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

func (h *poBoqHeaderHandler) SyncPrice(c *gin.Context) {

	pekerjaanNo := struct {
		PekerjaanNo string `json:"pekerjaan_no"`
	}{}

	c.ShouldBindJSON(&pekerjaanNo)

	response := domain.PoBoqHeaderResponse{
		Status:  http.StatusOK,
		Message: "Berhasil sync data header",
	}

	err := h.poBoqHeaderService.SyncActualPrice(pekerjaanNo.PekerjaanNo)
	if err != nil {
		response = domain.PoBoqHeaderResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}

	}

	c.JSON(response.Status, response)
}

func (h *poBoqHeaderHandler) GetByPekerjaanNoWithBody(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3

	var filter domain.PoBoqHeaderFilterRequest
	c.ShouldBindQuery(&filter)

	headers, err := h.poBoqHeaderService.GetByPekerjaanNoWithBody(FinalId, filter.Page, filter.PageSize)
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

	c.JSON(http.StatusOK, headers)
}

func (h *poBoqHeaderHandler) GetByPekerjaanNoWithBodyServerSide(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3

	var filter domain.PoBoqHeaderFilterRequestServerSide
	c.ShouldBindQuery(&filter)

	headers, err := h.poBoqHeaderService.GetByPekerjaanNoWithBodyServerSide(FinalId, filter.Page, filter.PageSize, filter.Item, filter.Desc)
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

	c.JSON(http.StatusOK, headers)
}

func (h *poBoqHeaderHandler) GetByPekerjaanNoWithBodyServerSideCount(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3
	var filter domain.PoBoqHeaderFilterRequestServerSide
	c.ShouldBindQuery(&filter)

	datas, err := h.poBoqHeaderService.GetByPekerjaanNoWithBodyServerSideCount(FinalId, filter.Item, filter.Desc)
	if err != nil {
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data",
		})
		return
	}

	c.JSON(http.StatusOK, datas)
}

func (h *poBoqHeaderHandler) GetByPekerjaanNoWithBodyByOrder(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4") // Order dari parameter URL

	// Gabungkan id menjadi FinalId
	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3

	// Konversi var4 (order) menjadi integer
	order, err := strconv.Atoi(var4)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Order harus berupa angka",
		})
		return
	}

	// Panggil service GetByPekerjaanNoWithBodyByOrder
	headers, err := h.poBoqHeaderService.GetByPekerjaanNoWithBodyByOrder(FinalId, order)
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

	// Berikan respon dengan data headers
	c.JSON(http.StatusOK, headers)
}
