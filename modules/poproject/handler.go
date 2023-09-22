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
	poProject.GET("/:id/:no/:var1/:var2/:var3", handler.GetByPo)
	poProject.POST("", handler.Store)

	//tambahan untuk header PO
	poProject.GET("/roll", handler.rollNum)
}

func NewPoProjectHandlerAddon(v1 *gin.RouterGroup, poProjectService Service) {
	handler := &poProjectHandler{poProjectService}

	addon := v1.Group("addons_project_company")
	addon2 := v1.Group("addons_project_plant")

	addon.GET("/:id", handler.GetCompany)
	addon2.GET("/:id", handler.GetPlant)
}

func (h *poProjectHandler) GetCompany(c *gin.Context) {
	ba := c.Param("id")

	addonProject, err := h.poProjectService.GetCompany(ba)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data addon tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data addon",
		})
		return
	}

	response := domain.AddonResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data addon",
		Data:    addonProject,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectHandler) GetPlant(c *gin.Context) {
	werks := c.Param("id")

	addonProject, err := h.poProjectService.GetPlant(werks)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data addon tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data addon",
		})
		return
	}

	response := domain.AddonResponse2{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data addon",
		Data:    addonProject,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poProjectHandler) rollNum(c *gin.Context) {
	rollNum, err := helper.GenerateHeaderBoq2(3)
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
	noPekerjaan := c.Param("no")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	addon := string("/")

	poProject, err := h.poProjectService.GetByPo(po, noPekerjaan+addon+var1+addon+var2+addon+var3)
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
	var input []domain.PoProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	for _, item := range input {
		createPoProject := domain.PoProject{
			Po:          item.Po,
			PekerjaanNo: item.PekerjaanNo,
			Item:        item.Item,
			Description: item.Description,
			Qty:         item.Qty,
			PoUnit:      item.PoUnit,
			Price:       item.Price,
			Currency:    item.Currency,
			Wbs:         item.Wbs,
			Cera:        item.Cera,
		}

		// Simpan setiap entri ke dalam basis data
		if _, err := h.poProjectService.Store(createPoProject); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data Po Project",
			})
			return
		}
	}

	response := domain.PoProjectResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Po Project",
	}

	c.JSON(http.StatusCreated, response)
}
