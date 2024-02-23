package poprogressheaderaddendum

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poProgressHeaderAddendumHandler struct {
	poProgressHeaderAddendumService Service
}

func NewPoProgressHeaderAddendumHandler(v1 *gin.RouterGroup, poProgressHeaderAddendumService Service) {
	handler := &poProgressHeaderAddendumHandler{poProgressHeaderAddendumService}

	header := v1.Group("progressheaderaddendum")

	header.GET("/:id/:var1/:var2/:var3/:var4", handler.GetAllProgressByRunNum)
	header.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)
	header.POST("", handler.Store)
	header.PUT("/:id/:var1/:var2/:var3/:var4", handler.Update)
}

func (h *poProgressHeaderAddendumHandler) GetAllProgressByRunNum(c *gin.Context) {
	id := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	datas, err := h.poProgressHeaderAddendumService.FindAllProg(id)
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
		"data":    datas,
	})
}

func (h *poProgressHeaderAddendumHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addon := "/"

	err := h.poProgressHeaderAddendumService.Delete(id + addon + var1 + addon + var2 + addon + var3 + addon + var4)
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

func (h *poProgressHeaderAddendumHandler) Store(c *gin.Context) {
	var input domain.PoProgressHeaderAddendum

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	data, err := h.poProgressHeaderAddendumService.Store(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Header Progress",
		})
		return
	}

	response := domain.PoProgressHeaderAddendumResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Header Progress",
		Data:    []domain.PoProgressHeaderAddendum{data},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poProgressHeaderAddendumHandler) Update(c *gin.Context) {
	id := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	var input domain.PoProgressHeaderAddendumUpdate
	c.BindJSON(&input)

	data, err := h.poProgressHeaderAddendumService.Update(id, input.Po, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Header Progress " + err.Error(),
		})
		return
	}

	response := domain.PoProgressHeaderAddendumResponseUpdate{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Header Progress",
		Data:    []domain.PoProgressHeaderAddendumUpdate{data},
	}

	c.JSON(http.StatusCreated, response)
}
