package popic

import (
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
