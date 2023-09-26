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
