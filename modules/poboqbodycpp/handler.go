package poboqbodycpp

import (
	"ebapp-api-dev/domain"
	"ebapp-api-dev/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poBoqBodyCppHandler struct {
	poBoqBodyCppService Service
}

func NewPoBoqBodyCppHandler(v1 *gin.RouterGroup, poBoqBodyCppService Service) {
	handler := &poBoqBodyCppHandler{poBoqBodyCppService}

	poboqbodyCpp := v1.Group("po_boq_body_cpp")
	poboqbodyCpp.Use(middlewares.AuthService())

	poboqbodyCpp.GET("/count/:id/:var1/:var2/:var3/:var4/:var5", handler.CountByRunNum)
	poboqbodyCpp.GET("/maxorder/:id/:var1/:var2/:var3/:var4/:var5", handler.SelectMaxOrder)
	poboqbodyCpp.POST("", handler.Store)
	poboqbodyCpp.PUT("/:id/:var1/:var2/:var3/:var4/:var5", handler.Update)
	poboqbodyCpp.GET("/:id/:var1/:var2/:var3/:var4/:var5/:var6", handler.GetBodyByID)
	poboqbodyCpp.DELETE("/:id/:var1/:var2/:var3/:var4/:var5", handler.Delete)
	poboqbodyCpp.POST("/clone", handler.CloneCpp)
}

func groupItemsByParent(items []domain.PoBoqBodyCppProgressResponse, parentId int) []domain.PoBoqBodyCppProgressResponse {
	var result []domain.PoBoqBodyCppProgressResponse

	for _, item := range items {
		if item.ParentId == parentId {
			children := groupItemsByParent(items, item.Id)
			item.Children = children
			result = append(result, item)
		}
	}

	return result
}

func (h *poBoqBodyCppHandler) GetBodyByID(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")
	var6 := c.Param("var6")
	addons := "/"

	poBoqBodyCpp, err := h.poBoqBodyCppService.GetByRunNum(runNum+addons+var1+addons+var2+addons+var3+addons+var4+addons+var5, runNum+addons+var1+addons+var2+addons+var3+addons+var4, var6)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	if len(poBoqBodyCpp) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data BoQ Body tidak ditemukan",
			"data":    nil,
		})
		return
	}

	var poBoqBodyCppResponse []domain.PoBoqBodyCppProgressResponse
	for _, body := range poBoqBodyCpp {
		poBoqBodyCppResponse = append(poBoqBodyCppResponse, domain.PoBoqBodyCppProgressResponse{
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
			StartDate:         body.StartDate,
			EndDate:           body.EndDate,
			StartDateActual:   body.StartDateActual,
			EndDateActual:     body.EndDateActual,
			PreviousVolume:    body.PreviousVolume,
			CurrentVolume:     body.CurrentVolume,
			Status:            body.Status,
			Note:              body.Note,
			RunNumCpp:         body.RunNumCpp,
		})
	}

	result := groupItemsByParent(poBoqBodyCppResponse, 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data PO BoQ Body",
		"data":    result,
	})
}

func (h *poBoqBodyCppHandler) CountByRunNum(c *gin.Context) {
	runNum := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4") + "/" + c.Param("var5")

	// Panggil service untuk menghitung jumlah entri dengan runNum tertentu
	total, err := h.poBoqBodyCppService.CountByRunNum(runNum)
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

func (h *poBoqBodyCppHandler) SelectMaxOrder(c *gin.Context) {
	runNum := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4") + "/" + c.Param("var5")

	total, err := h.poBoqBodyCppService.SelectMaxOrder(runNum)
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

func (h *poBoqBodyCppHandler) Store(c *gin.Context) {
	var input []domain.PoBoqBodyCppRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createdPoBoqBodiesCpp := []domain.PoBoqBodyCpp{}

	for _, requestData := range input {
		existingBoqBody, _ := h.poBoqBodyCppService.FindByItemNo(requestData.ItemNo)

		if existingBoqBody.Id != 0 && existingBoqBody.RunNum == requestData.RunNum && existingBoqBody.Order == requestData.Order {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "ItemNo sudah ada di database untuk RunNum yang sama dan Order yang sama",
			})
			return
		}

		poBoqBodies, err := h.poBoqBodyCppService.Store(domain.PoBoqBodyCpp(requestData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data Boq Body",
			})
			return
		}

		createdPoBoqBodiesCpp = append(createdPoBoqBodiesCpp, poBoqBodies)
	}

	response := domain.PoBoqBodyCppResponseFinal{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data BoQ Body",
		Data:    createdPoBoqBodiesCpp,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poBoqBodyCppHandler) Update(c *gin.Context) {
	runNum := c.Param("id") + "/" + c.Param("var1") + "/" + c.Param("var2") + "/" + c.Param("var3") + "/" + c.Param("var4") + "/" + c.Param("var5")

	// Membaca data yang dikirimkan dalam body permintaan
	var requestBody struct {
		Order    string `json:"order"`
		MainId   int    `json:"main_id"`
		ParentId int    `json:"parent_id"`
		Status   bool   `json:"status"`
		Note     string `json:"note"`
	}

	// Melakukan penguraian data JSON yang diterima ke dalam struktur requestBody
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		// Mengirimkan respons jika terjadi kesalahan saat penguraian JSON
		response := domain.PoBoqBodyCppResponseFinal{
			Status:  http.StatusBadRequest,
			Message: "Gagal memperbarui data BoQ Body: " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusBadRequest, response)
		return
	}

	// Memanggil service untuk melakukan pembaruan data
	updatedCpp, err := h.poBoqBodyCppService.Update(runNum, requestBody.Order, requestBody.MainId, requestBody.ParentId, requestBody.Status, requestBody.Note)
	if err != nil {
		// Mengirimkan respons jika terjadi kesalahan saat melakukan pembaruan
		response := domain.PoBoqBodyCppResponseFinal{
			Status:  http.StatusInternalServerError,
			Message: "Gagal memperbarui data BoQ Body " + err.Error(),
			Data:    nil,
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Mengirimkan respons dengan data yang telah diperbarui
	response := domain.PoBoqBodyCppResponseUpdate{
		Status:  http.StatusOK,
		Message: "Berhasil memperbarui data BoQ Body",
		Data:    []domain.PoBoqBodyCppProgress{updatedCpp},
	}
	c.JSON(http.StatusOK, response)
}

func (h *poBoqBodyCppHandler) Delete(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")
	addons := "/"

	err := h.poBoqBodyCppService.Delete(runNum + addons + var1 + addons + var2 + addons + var3 + addons + var4 + addons + var5)
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

func (h *poBoqBodyCppHandler) CloneCpp(c *gin.Context) {

	pekerjaanNo := struct {
		PreviousRunNum string `json:"previous_run_num"`
		NextRunNum     string `json:"next_run_num"`
	}{}

	c.ShouldBindJSON(&pekerjaanNo)

	response := domain.PoBoqHeaderResponse{
		Status:  http.StatusOK,
		Message: "Berhasil clone data body",
	}

	err := h.poBoqBodyCppService.CloneCpp(pekerjaanNo.PreviousRunNum, pekerjaanNo.NextRunNum)
	if err != nil {
		response = domain.PoBoqHeaderResponse{
			Status:  http.StatusBadRequest,
			Message: err.Error(),
		}

	}

	c.JSON(response.Status, response)
}
