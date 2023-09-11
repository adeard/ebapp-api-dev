package poproject

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poProjectHandler struct {
	poProjectService Service
}

func NewPoProjectHandler(v1 *gin.RouterGroup, poProjectService Service) {
	handler := &poProjectHandler{poProjectService}

	poProject := v1.Group("po_project")

	poProject.GET("", handler.GetAll)
	poProject.GET("/:id", handler.GetByPo)
	poProject.POST("", handler.Store)
	poProject.GET("/roll", handler.rollNum)
}

func (h *poProjectHandler) rollNum(c *gin.Context) {
	rollNum, err := helper.GenerateHeaderBoq(3)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	response := domain.PoProjectResponse2{
		Status:  http.StatusOK,
		Message: "Data berhasil dibuat",
		Data:    rollNum,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectHandler) GetAll(c *gin.Context) {
	var input domain.PoProjectRequest

	poProject, err := h.poProjectService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data PO Project List",
		})
		return
	}

	response := domain.PoProjectResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data PO Project List",
		Data:    poProject,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectHandler) GetByPo(c *gin.Context) {
	po := c.Param("id")

	poProject, err := h.poProjectService.GetByPo(po)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data Po Project tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Po Project",
		})
		return
	}

	response := domain.PoProjectResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data Po Project",
		Data:    poProject,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectHandler) Store(c *gin.Context) {
	var input domain.PoProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createPoProject := domain.PoProject{
		Po:          input.Po,
		Item:        input.Item,
		Description: input.Description,
		Qty:         input.Qty,
		Price:       input.Price,
		Wbs:         input.Wbs,
		Cera:        input.Cera,
	}

	poProjects, err := h.poProjectService.Store(createPoProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Po Project",
		})
		return
	}

	response := domain.PoProjectResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Po Project",
		Data:    []domain.PoProject{poProjects},
	}

	c.JSON(http.StatusCreated, response)
}
