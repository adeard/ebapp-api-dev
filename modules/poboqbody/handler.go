package poboqbody

import (
	"ebapp-api-dev/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type poBoqBodyHandler struct {
	poBoqBodyService Service
}

func NewPoBoqBodyHandler(v1 *gin.RouterGroup, poBoqService Service) {

	handler := &poBoqBodyHandler{poBoqService}

	poboqbody := v1.Group("po_boq_body")

	poboqbody.GET("/:id/:var1/:var2/:var3/:var4", handler.GetPoBoqBodyByRunNum)
	poboqbody.POST("", handler.Store)
	poboqbody.DELETE("/:id/:var1/:var2/:var3/:var4/:var5", handler.Delete)
	poboqbody.DELETE("/:id/:var1/:var2/:var3/:var4", handler.DeleteByOrder)
	poboqbody.PUT("", handler.Update)
}

func groupItemsByParent(items []domain.PoBoqBodyResponse, parentId int) []domain.PoBoqBodyResponse {
	var result []domain.PoBoqBodyResponse

	for _, item := range items {
		if item.ParentId == parentId {
			children := groupItemsByParent(items, item.Id)
			item.Children = children
			result = append(result, item)
		}
	}

	return result
}

func (h *poBoqBodyHandler) GetPoBoqBodyByRunNum(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addons := "/"

	poBoqBody, err := h.poBoqBodyService.GetByRunNum(runNum+addons+var1+addons+var2+addons+var3, var4)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	if len(poBoqBody) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data BoQ Body tidak ditemukan",
			"data":    nil,
		})
		return
	}

	var poBoqBodyResponse []domain.PoBoqBodyResponse
	for _, body := range poBoqBody {
		poBoqBodyResponse = append(poBoqBodyResponse, domain.PoBoqBodyResponse{
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
		})
	}

	result := groupItemsByParent(poBoqBodyResponse, 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data PO BoQ Body",
		"data":    result,
	})
}

func (h *poBoqBodyHandler) Store(c *gin.Context) {
	var input []domain.PoBoqBodyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	createdPoBoqBodies := []domain.PoBoqBody{}

	for _, requestData := range input {
		existingBoqBody, _ := h.poBoqBodyService.FindByItemNo(requestData.ItemNo)

		if existingBoqBody.Id != 0 && existingBoqBody.RunNum == requestData.RunNum && existingBoqBody.Order == requestData.Order {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "ItemNo sudah ada di database untuk RunNum yang sama dan Order yang sama",
			})
			return
		}

		poBoqBodies, err := h.poBoqBodyService.Store(domain.PoBoqBody(requestData))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal meneruskan data BoQ Body",
			})
			return
		}

		createdPoBoqBodies = append(createdPoBoqBodies, poBoqBodies)
	}

	response := domain.PoBoqBodyResponseFinal{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data BoQ Body",
		Data:    createdPoBoqBodies,
	}

	c.JSON(http.StatusCreated, response)
}

func (h *poBoqBodyHandler) Update(c *gin.Context) {
	var input domain.PoBoqBody

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	// Cek apakah item ada di database
	result, err := h.poBoqBodyService.CheckBoqBody(input.RunNum, input.Order, strconv.Itoa(input.Id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak dapat menemukan Item",
		})
		return
	}

	if len(result) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusNotFound,
			"message": "Item yang dicari tidak ditemukan",
		})
		return
	}

	// Melakukan pembaruan item jika ada
	_, updateErr := h.poBoqBodyService.Update(input)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengupdate data PoBoqBody",
		})
		return
	}

	response := domain.PoBoqBodyResponseFinal{
		Status:  http.StatusOK,
		Message: "Berhasil mengupdate data PoBoqBody",
		Data:    nil,
	}

	c.JSON(http.StatusOK, response)
}

func (h *poBoqBodyHandler) Delete(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	var5 := c.Param("var5")
	addons := "/"

	_, err := h.poBoqBodyService.GetByRunNum(runNum+addons+var1+addons+var2+addons+var3, var4)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data",
			"data":    nil,
		})
		return
	}

	err = h.poBoqBodyService.Delete(runNum+addons+var1+addons+var2+addons+var3, var4, var5)
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

func (h *poBoqBodyHandler) DeleteByOrder(c *gin.Context) {
	runNum := c.Param("id")
	var1 := c.Param("var1")
	var2 := c.Param("var2")
	var3 := c.Param("var3")
	var4 := c.Param("var4")
	addons := "/"

	_, err := h.poBoqBodyService.GetByRunNum(runNum+addons+var1+addons+var2+addons+var3, var4)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data",
			"data":    nil,
		})
		return
	}

	err = h.poBoqBodyService.DeleteByOrder(runNum+addons+var1+addons+var2+addons+var3, var4)
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
