package poprogressheaderaddendum

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type poProgressHeaderAddendumHandler struct {
	poProgressHeaderAddendumService Service
}

func NewPoProgressHeaderAddendumHandler(v1 *gin.RouterGroup, poProgressHeaderAddendumService Service) {
	handler := &poProgressHeaderAddendumHandler{poProgressHeaderAddendumService}

	header := v1.Group("progressheader")

	header.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)
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
