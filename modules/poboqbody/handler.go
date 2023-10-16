package poboqbody

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type poBoqBodyHandler struct {
	poBoqBodyService Service
}

func NewPoBoqBodyHandler(v1 *gin.RouterGroup, poBoqService Service) {

	handler := &poBoqBodyHandler{poBoqService}

	poboqbody := v1.Group("po_boq_body")

	poboqbody.GET("/:id/:var1/:var2/:var3", handler.GetPoBoqBodyByRunNum)
	poboqbody.POST("", handler.Store)

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
	addons := "/"

	poBoqBody, err := h.poBoqBodyService.GetByRunNum(runNum + addons + var1 + addons + var2 + addons + var3)
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

		if existingBoqBody.Id != 0 && existingBoqBody.RunNum == requestData.RunNum {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"message": "ItemNo sudah ada di database untuk RunNum yang sama",
			})
			return
		}

		createdPoBoqBody := domain.PoBoqBody{
			RunNum:            requestData.RunNum,
			Id:                requestData.Id,
			ParentId:          requestData.ParentId,
			ItemNo:            requestData.ItemNo,
			ItemLevel:         requestData.ItemLevel,
			ItemDescription:   requestData.ItemDescription,
			ItemSpecification: requestData.ItemSpecification,
			Qty:               requestData.Qty,
			Unit:              requestData.Unit,
			Price:             requestData.Price,
			Currency:          requestData.Currency,
			Note:              requestData.Note,
			Order:             requestData.Order,
		}

		poBoqBodies, err := h.poBoqBodyService.Store(createdPoBoqBody)
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
