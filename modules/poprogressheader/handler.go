package poprogressheader

import (
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
