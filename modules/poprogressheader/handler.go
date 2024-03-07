package poprogressheader

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poProgressHeaderHandler struct {
	poProgressHeaderService Service
}

func NewPoProgressHeaderHandler(v1 *gin.RouterGroup, poProgressHeaderService Service) {
	handler := &poProgressHeaderHandler{poProgressHeaderService}

	header := v1.Group("progressheader")

	header.GET("/:id/:var1/:var2/:var3/:var4", handler.GetProgrssByRunNum)
	header.GET("/:id/:var1/:var2/:var3", handler.GetAllProgressByRunNum)
	header.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)
	header.PUT("/:id/:var1/:var2/:var3/:var4", handler.Update)
	header.PUT("/isebapp/:id/:var1/:var2/:var3/:var4", handler.UpdateEbapp)
	header.POST("", handler.Store)
}

func (h *poProgressHeaderHandler) GetProgrssByRunNum(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addon := "/"

	data, err := h.poProgressHeaderService.FindProg(id + addon + var1 + addon + var2 + addon + var3 + addon + var4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Progress " + id,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data Progress",
		"data":    data,
	})
}

func (h *poProgressHeaderHandler) GetAllProgressByRunNum(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	addon := "/"

	datas, err := h.poProgressHeaderService.FindAllProg(id + addon + var1 + addon + var2 + addon + var3 + addon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Progress " + id + addon + var1 + addon + var2 + addon + var3,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data Progress",
		"data":    datas,
	})
}

func (h *poProgressHeaderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addon := "/"

	err := h.poProgressHeaderService.Delete(id + addon + var1 + addon + var2 + addon + var3 + addon + var4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menghapus data Progress " + id + addon + var1 + addon + var2 + addon + var3 + addon + var4,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil menghapus data Progress " + id + addon + var1 + addon + var2 + addon + var3 + addon + var4,
		"data":    nil,
	})
}

func (h *poProgressHeaderHandler) Update(c *gin.Context) {
	id := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	var input domain.PoProgressHeaderUpdate
	c.BindJSON(&input)

	data, err := h.poProgressHeaderService.Update(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Main Header Progress " + err.Error(),
		})
		return
	}

	response := domain.PoProgressHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Main Header Progress",
		Data:    []domain.PoProgressHeader{data},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poProgressHeaderHandler) UpdateEbapp(c *gin.Context) {
	id := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	var input domain.PoProgressHeaderUpdateEbapp
	c.BindJSON(&input)

	data, err := h.poProgressHeaderService.EbappUpdate(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Main Header Progress " + err.Error(),
		})
		return
	}

	response := domain.PoProgressHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Main Header Progress",
		Data:    []domain.PoProgressHeader{data},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poProgressHeaderHandler) Store(c *gin.Context) {
	var input domain.PoProgressHeader

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	data, err := h.poProgressHeaderService.Store(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Header Progress",
		})
		return
	}

	response := domain.PoProgressHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Header Progress",
		Data:    []domain.PoProgressHeader{data},
	}

	c.JSON(http.StatusCreated, response)
}
