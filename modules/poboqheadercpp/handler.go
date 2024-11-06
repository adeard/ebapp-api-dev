package poboqheadercpp

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poBoqHeaderCppHandler struct {
	poBoqHeaderCppService Service
}

func NewPoBoqHeaderCppHandler(v1 *gin.RouterGroup, poBoqHeaderCppService Service) {
	handler := &poBoqHeaderCppHandler{poBoqHeaderCppService}

	header := v1.Group("poboq_header_cpp")
	header.Use(middlewares.AuthService())

	header.GET("/:id/:var1/:var2/:var3/:var4", handler.GetCpp)
	header.POST("", handler.Store)
	header.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)
	header.GET("getss/:id/:var1/:var2/:var3/:var4/:var5", handler.GetBodyServerSide)
	header.GET("getsscount/:id/:var1/:var2/:var3/:var4/:var5", handler.GetBodyServerSideCount)
}

func (h *poBoqHeaderCppHandler) GetCpp(c *gin.Context) {
	FinalId := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	headers, err := h.poBoqHeaderCppService.GetCpp(FinalId)
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

	response := domain.PoBoqHeaderCppResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Header",
		Data:    headers,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poBoqHeaderCppHandler) Store(c *gin.Context) {
	var input []domain.PoBoqHeaderCpp

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	for _, item := range input {
		createHeader := domain.PoBoqHeaderCpp{
			PekerjaanNo: item.PekerjaanNo,
			Po:          item.Po,
			Item:        item.Item,
			Description: item.Description,
			Qty:         item.Qty,
		}

		if _, err := h.poBoqHeaderCppService.Store(createHeader); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data header Cpp",
			})
			return
		}
	}

	response := domain.PoBoqHeaderCppResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data header",
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poBoqHeaderCppHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3 + "/" + var4

	if deleteErr := h.poBoqHeaderCppService.Delete(FinalId); deleteErr != nil {
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

func (h *poBoqHeaderCppHandler) GetBodyServerSide(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3 + "/" + var4 + "/" + var5

	var filter domain.PoBoqHeaderFilterRequestServerSide
	c.ShouldBindQuery(&filter)

	headers, err := h.poBoqHeaderCppService.GetBodyServerSide(FinalId, filter.Page, filter.PageSize, filter.Item, filter.Desc)
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

func (h *poBoqHeaderCppHandler) GetBodyServerSideCount(c *gin.Context) {

	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")

	FinalId := id + "/" + var1 + "/" + var2 + "/" + var3 + "/" + var4 + "/" + var5
	var filter domain.PoBoqHeaderFilterRequestServerSide
	c.ShouldBindQuery(&filter)

	datas, err := h.poBoqHeaderCppService.GetBodyServerSideCount(FinalId, filter.Item, filter.Desc)
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
