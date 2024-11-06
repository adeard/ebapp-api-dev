package poboqheaderprogress

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poBoqHeaderProgressHandler struct {
	poBoqHeaderProgressService Service
}

func NewPoBoqHeaderProgressHandler(v1 *gin.RouterGroup, poBoqHeaderProgressService Service) {
	handler := &poBoqHeaderProgressHandler{poBoqHeaderProgressService}

	header := v1.Group("poboq_header_progress")
	header.Use(middlewares.AuthService())

	header.GET("/:id/:var1/:var2/:var3/:var4", handler.GetProgress)
	header.POST("", handler.Store)
	header.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)

	header.GET("getss/:id/:var1/:var2/:var3/:var4", handler.GetBodyServerSide)
	header.GET("getsscount/:id/:var1/:var2/:var3/:var4", handler.GetBodyServerSideCount)
	header.GET("sync/:id/:var1/:var2/:var3/:var4", handler.SyncProgress)
}

func (h *poBoqHeaderProgressHandler) GetProgress(c *gin.Context) {
	FinalId := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	headers, err := h.poBoqHeaderProgressService.GetProgress(FinalId)
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

	response := domain.PoBoqHeaderProgressResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Header",
		Data:    headers,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poBoqHeaderProgressHandler) Store(c *gin.Context) {
	var input []domain.PoBoqHeaderProgress

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	for _, item := range input {
		createHeader := domain.PoBoqHeaderProgress{
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
			Percentage:  item.Percentage,
		}

		if _, err := h.poBoqHeaderProgressService.Store(createHeader); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data header progress",
			})
			return
		}
	}

	response := domain.PoBoqHeaderProgressResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data header",
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poBoqHeaderProgressHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3 + "/" + var4

	if deleteErr := h.poBoqHeaderProgressService.Delete(FinalId); deleteErr != nil {
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

func (h *poBoqHeaderProgressHandler) GetBodyServerSide(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3 + "/" + var4

	var filter domain.PoBoqHeaderFilterRequestServerSide
	c.ShouldBindQuery(&filter)

	headers, err := h.poBoqHeaderProgressService.GetBodyServerSide(FinalId, filter.Page, filter.PageSize, filter.Item, filter.Desc)
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

func (h *poBoqHeaderProgressHandler) GetBodyServerSideCount(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3 + "/" + var4
	var filter domain.PoBoqHeaderFilterRequestServerSide
	c.ShouldBindQuery(&filter)

	datas, err := h.poBoqHeaderProgressService.GetBodyServerSideCount(FinalId, filter.Item, filter.Desc)
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

func (h *poBoqHeaderProgressHandler) SyncProgress(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addon := "/"

	err := h.poBoqHeaderProgressService.SyncProgress(id + addon + var1 + addon + var2 + addon + var3 + addon + var4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal Sync data Progress " + id + addon + var1 + addon + var2 + addon + var3 + addon + var4,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil Sync data Progress " + id + addon + var1 + addon + var2 + addon + var3 + addon + var4,
		"data":    nil,
	})
}
