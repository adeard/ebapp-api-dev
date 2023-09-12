package listproject

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type listProjectHandler struct {
	listProjectService Service
}

func NewListProjectHandler(v1 *gin.RouterGroup, listProjectService Service) {
	handler := &listProjectHandler{listProjectService}

	listProject := v1.Group("list_project")
	project := v1.Group("project")

	listProject.GET("", handler.GetAll)
	project.GET("/:id", handler.GetByID)
	project.POST("", handler.Store)
	project.POST("/draft2", handler.Store2)
	project.POST("/draft3", handler.Store3)
	project.POST("/draft4", handler.Store4)
}

func (h *listProjectHandler) GetAll(c *gin.Context) {
	var input domain.ListProjectRequest

	listProjects, err := h.listProjectService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Project List",
		})
		return
	}

	response := domain.ListProjectsResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data List Project",
		Data:    listProjects,
	}

	c.JSON(http.StatusOK, response)
}

func (h *listProjectHandler) GetByID(c *gin.Context) {
	id := c.Param("id")

	project, err := h.listProjectService.GetByID(id)
	if err != nil {
		// Cek apakah error disebabkan oleh data tidak ditemukan.
		if err == domain.ErrNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusNotFound,
				"message": "Data project tidak ditemukan",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data project",
		})
		return
	}

	response := domain.ListProjectsResponse{
		Status:  http.StatusOK,
		Message: "Berhasil mengambil data BoQ Header",
		Data:    []domain.ListProject{project},
	}

	c.JSON(http.StatusOK, response)
}

func (h *listProjectHandler) Store(c *gin.Context) {
	var input domain.ListProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createProject := domain.ListProject{
		Po:              input.Po,
		PoDate:          input.PoDate,
		PekerjaanNo:     input.PekerjaanNo,
		PekerjaanName:   input.PekerjaanName,
		PekerjaanDate:   input.PekerjaanDate,
		StartDate:       input.StartDate,
		EndDate:         input.EndDate,
		StartDateActual: input.StartDateActual,
		EndDateActual:   input.EndDateActual,
		UnitUsaha:       input.UnitUsaha,
		Pt:              input.Pt,
		Vendor:          input.Vendor,
		Status:          input.Status,
		Total:           input.Total,
		Currency:        input.Currency,
		Area:            input.Area,
	}

	project, err := h.listProjectService.Store(createProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Project",
		})
		return
	}

	response := domain.ListProjectsResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Po Project",
		Data:    []domain.ListProject{project},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *listProjectHandler) Store2(c *gin.Context) {
	var input domain.ListProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createProject := domain.ListProject2{
		Po:            input.Po,
		PoDate:        input.PoDate,
		PekerjaanNo:   input.PekerjaanNo,
		PekerjaanName: input.PekerjaanName,
		PekerjaanDate: input.PekerjaanDate,
		StartDate:     input.StartDate,
		EndDate:       input.EndDate,
		UnitUsaha:     input.UnitUsaha,
		Pt:            input.Pt,
		Vendor:        input.Vendor,
		Status:        input.Status,
		Total:         input.Total,
		Currency:      input.Currency,
		Area:          input.Area,
	}

	project, err := h.listProjectService.Store2(createProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Project",
		})
		return
	}

	response := domain.ListProjectsResponse2{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Po Project",
		Data:    []domain.ListProject2{project},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *listProjectHandler) Store3(c *gin.Context) {
	var input domain.ListProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createProject := domain.ListProject3{
		Po:              input.Po,
		PoDate:          input.PoDate,
		PekerjaanNo:     input.PekerjaanNo,
		PekerjaanName:   input.PekerjaanName,
		PekerjaanDate:   input.PekerjaanDate,
		StartDateActual: input.StartDateActual,
		EndDateActual:   input.EndDateActual,
		UnitUsaha:       input.UnitUsaha,
		Pt:              input.Pt,
		Vendor:          input.Vendor,
		Status:          input.Status,
		Total:           input.Total,
		Currency:        input.Currency,
		Area:            input.Area,
	}

	project, err := h.listProjectService.Store3(createProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Project",
		})
		return
	}

	response := domain.ListProjectsResponse3{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Po Project",
		Data:    []domain.ListProject3{project},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *listProjectHandler) Store4(c *gin.Context) {
	var input domain.ListProjectRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createProject := domain.ListProject4{
		Po:            input.Po,
		PoDate:        input.PoDate,
		PekerjaanNo:   input.PekerjaanNo,
		PekerjaanName: input.PekerjaanName,
		PekerjaanDate: input.PekerjaanDate,
		UnitUsaha:     input.UnitUsaha,
		Pt:            input.Pt,
		Vendor:        input.Vendor,
		Status:        input.Status,
		Total:         input.Total,
		Currency:      input.Currency,
		Area:          input.Area,
	}

	project, err := h.listProjectService.Store4(createProject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Project",
		})
		return
	}

	response := domain.ListProjectsResponse4{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Po Project",
		Data:    []domain.ListProject4{project},
	}

	c.JSON(http.StatusCreated, response)
}
