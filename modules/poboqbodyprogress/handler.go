package poboqbodyprogress

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poBoqBodyProgressHandler struct {
	poBoqBodyProgressService Service
}

func NewPoBoqBodyProgressHandler(v1 *gin.RouterGroup, poBoqBodyProgressService Service) {
	handler := &poBoqBodyProgressHandler{poBoqBodyProgressService}

	poboqbodyprogress := v1.Group("po_boq_body_progress")

	poboqbodyprogress.POST("", handler.Store)
	poboqbodyprogress.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)
}

func (h *poBoqBodyProgressHandler) Store(c *gin.Context) {
	var input []domain.PoBoqBodyProgressRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createdPoBoqBodiesProgress := []domain.PoBoqBodyProgress{}

	for _, requestData := range input {
		existingBoqBody, _ := h.poBoqBodyProgressService.FindByItemNo(requestData.ItemNo)

		if existingBoqBody.Id != 0 && existingBoqBody.RunNum == requestData.RunNum && existingBoqBody.Order == requestData.Order {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "ItemNo sudah ada di database untuk RunNum yang sama dan Order yang sama",
			})
			return
		}

		poBoqBodies, err := h.poBoqBodyProgressService.Store(domain.PoBoqBodyProgress(requestData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data Boq Body",
			})
			return
		}

		createdPoBoqBodiesProgress = append(createdPoBoqBodiesProgress, poBoqBodies)

		response := domain.PoBoqBodyProgressResponseFinal{
			Status:  http.StatusCreated,
			Message: "Berhasil menyimpan data BoQ Body",
			Data:    createdPoBoqBodiesProgress,
		}

		c.JSON(http.StatusCreated, response)
	}
}

func (h *poBoqBodyProgressHandler) Delete(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addons := "/"

	err := h.poBoqBodyProgressService.Delete(runNum + addons + var1 + addons + var2 + addons + var3 + addons + var4)
	if err != nil {
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
