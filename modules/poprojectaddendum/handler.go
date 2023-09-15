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
