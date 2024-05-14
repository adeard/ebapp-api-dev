package pocppheader

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poCppHeaderHandler struct {
	poCppHeaderService Service
}

func NewPoCppHeaderHandler(v1 *gin.RouterGroup, poCppHeaderService Service) {
	handler := &poCppHeaderHandler{poCppHeaderService}

	header := v1.Group("cppheader")

	header.GET("/:id/:var1/:var2/:var3", handler.GetCppByRunNum)
	header.GET("/by_run_num_progress/:id/:var1/:var2/:var3/:var4", handler.GetCppByRunNumProgress)
	header.GET("/:id/:var1/:var2/:var3/:var4", handler.GetAllCppByRunNum)
	header.DELETE("/:id/:var1/:var2/:var3/:var4/:var5", handler.Delete)
	header.PUT("/:id/:var1/:var2/:var3/:var4/:var5", handler.Update)
	header.PUT("/status/:id/:var1/:var2/:var3/:var4/:var5", handler.UpdateStatus)
	header.POST("", handler.Store)
}

func (h *poCppHeaderHandler) GetCppByRunNum(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	addon := "/"

	data, err := h.poCppHeaderService.FindCpp(id + addon + var1 + addon + var2 + addon + var3)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    nil,
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": err.Error(),
				"data":    nil,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data Cpp",
		"data":    data,
	})
}

func (h *poCppHeaderHandler) GetCppByRunNumProgress(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addon := "/"

	data, err := h.poCppHeaderService.FindCppByRunNumProgress(id + addon + var1 + addon + var2 + addon + var3 + addon + var4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data Cpp",
		"data":    data,
	})
}

func (h *poCppHeaderHandler) GetAllCppByRunNum(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addon := "/"

	datas, err := h.poCppHeaderService.FindAllProg(id + addon + var1 + addon + var2 + addon + var3 + addon + var4 + addon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data Cpp " + id + addon + var1 + addon + var2 + addon + var3,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data Cpp",
		"data":    datas,
	})
}

func (h *poCppHeaderHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")
	addon := "/"

	err := h.poCppHeaderService.Delete(id + addon + var1 + addon + var2 + addon + var3 + addon + var4 + addon + var5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menghapus data Cpp " + id + addon + var1 + addon + var2 + addon + var3 + addon + var4 + addon + var5,
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil menghapus data Cpp " + id + addon + var1 + addon + var2 + addon + var3 + addon + var4 + addon + var5,
		"data":    nil,
	})
}

func (h *poCppHeaderHandler) Update(c *gin.Context) {
	id := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4") + "/" + c.Param("var5")

	var input domain.PoCppHeaderUpdate
	c.BindJSON(&input)

	data, err := h.poCppHeaderService.Update(id, input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Main Header Cpp " + err.Error(),
		})
		return
	}

	response := domain.PoCppHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Main Header Cpp",
		Data:    []domain.PoCppHeader{data},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poCppHeaderHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4") + "/" + c.Param("var5")

	var input domain.PoCppHeaderUpdateStatus

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	data, err := h.poCppHeaderService.UpdateStatus(id, input.StatusProgress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Main Header Cpp " + err.Error(),
		})
		return
	}

	response := domain.PoCppHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data data Main Header Cpp",
		Data:    []domain.PoCppHeader{data},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poCppHeaderHandler) Store(c *gin.Context) {
	var input domain.PoCppHeader

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	data, err := h.poCppHeaderService.Store(input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data Header Cpp",
		})
		return
	}

	response := domain.PoCppHeaderResponse{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data Header Cpp",
		Data:    []domain.PoCppHeader{data},
	}

	c.JSON(http.StatusCreated, response)
}
