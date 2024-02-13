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

	poboqbodyprogress.GET("/count/:id/:var1/:var2/:var3/:var4", handler.CountByRunNum)
	poboqbodyprogress.GET("/maxorder/:id/:var1/:var2/:var3/:var4", handler.SelectMaxOrder)
	poboqbodyprogress.POST("", handler.Store)
	poboqbodyprogress.PUT("/:id/:var1/:var2/:var3/:var4", handler.Update)
	poboqbodyprogress.GET("/:id/:var1/:var2/:var3/:var4/:var5", handler.GetBodyByID)
	poboqbodyprogress.DELETE("/:id/:var1/:var2/:var3/:var4", handler.Delete)
}

func groupItemsByParent(items []domain.PoBoqBodyProgressResponse, parentId int) []domain.PoBoqBodyProgressResponse {
	var result []domain.PoBoqBodyProgressResponse

	for _, item := range items {
		if item.ParentId == parentId {
			children := groupItemsByParent(items, item.Id)
			item.Children = children
			result = append(result, item)
		}
	}

	return result
}

func (h *poBoqBodyProgressHandler) GetBodyByID(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")
	addons := "/"

	poBoqBodyProgress, err := h.poBoqBodyProgressService.GetByRunNum(runNum+addons+var1+addons+var2+addons+var3+addons+var4, var5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	if len(poBoqBodyProgress) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data BoQ Body tidak ditemukan",
			"data":    nil,
		})
		return
	}

	var poBoqBodyProgressResponse []domain.PoBoqBodyProgressResponse
	for _, body := range poBoqBodyProgress {
		poBoqBodyProgressResponse = append(poBoqBodyProgressResponse, domain.PoBoqBodyProgressResponse{
			Id:                body.Id,
			ParentId:          body.ParentId,
			RunNum:            body.RunNum,
			Order:             body.Order,
			ItemNo:            body.ItemNo,
			ItemLevel:         body.ItemLevel,
			ItemDescription:   body.ItemDescription,
			ItemSpecification: body.ItemSpecification,
			Qty:               body.Qty,
			Unit:              body.Unit,
			Price:             body.Price,
			Currency:          body.Currency,
			Note:              body.Note,
			StartDate:         body.StartDate,
			EndDate:           body.EndDate,
			StartDateActual:   body.StartDateActual,
			EndDateActual:     body.EndDateActual,
			PreviousVolume:    body.PreviousVolume,
			CurrentVolume:     body.CurrentVolume,
		})
	}

	result := groupItemsByParent(poBoqBodyProgressResponse, 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data PO BoQ Body",
		"data":    result,
	})
}

func (h *poBoqBodyProgressHandler) CountByRunNum(c *gin.Context) {
	runNum := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	// Panggil service untuk menghitung jumlah entri dengan runNum tertentu
	total, err := h.poBoqBodyProgressService.CountByRunNum(runNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menghitung jumlah entri",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Jumlah entri berhasil dihitung",
		"total":   total,
	})
}

func (h *poBoqBodyProgressHandler) SelectMaxOrder(c *gin.Context) {
	runNum := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	total, err := h.poBoqBodyProgressService.SelectMaxOrder(runNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal menghitung jumlah entri",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Jumlah entri berhasil dihitung",
		"max":     total,
	})
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

func (h *poBoqBodyProgressHandler) Update(c *gin.Context) {
	runNum := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4")

	// Membaca data yang dikirimkan dalam body permintaan
	var requestBody struct {
		Order         string  `json:"order"`
		MainId        int     `json:"main_id"`
		ParentId      int     `json:"parent_id"`
		CurrentVolume float64 `json:"current_volume"`
	}

	// Melakukan penguraian data JSON yang diterima ke dalam struktur requestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		// Mengirimkan respons jika terjadi kesalahan saat penguraian JSON
		response := domain.PoBoqBodyProgressResponseFinal{
			Status:  http.StatusBadRequest,
			Message: "Gagal memperbarui data BoQ Body: " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Memanggil service untuk melakukan pembaruan data
	updatedProgress, err := h.poBoqBodyProgressService.Update(runNum, requestBody.Order, requestBody.MainId, requestBody.ParentId, requestBody.CurrentVolume)
	if err != nil {
		// Mengirimkan respons jika terjadi kesalahan saat melakukan pembaruan
		response := domain.PoBoqBodyProgressResponseFinal{
			Status:  http.StatusInternalServerError,
			Message: "Gagal memperbarui data BoQ Body " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Mengirimkan respons dengan data yang telah diperbarui
	response := domain.PoBoqBodyProgressResponseFinal{
		Status:  http.StatusOK,
		Message: "Berhasil memperbarui data BoQ Body",
		Data:    []domain.PoBoqBodyProgress{updatedProgress},
	}
	c.JSON(http.StatusOK, response)
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
