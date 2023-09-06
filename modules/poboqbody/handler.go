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

	poboqbody.GET("/:id", handler.GetPoBoqBodyByRunNum)
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

	poBoqBody, err := h.poBoqBodyService.GetByRunNum(runNum)
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
		"message": "Berhasil mengambil data BoQ Body",
		"data":    result,
	})
}

func (h *poBoqBodyHandler) Store(c *gin.Context) {
	var input domain.PoBoqBodyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	existingBoqBody, _ := h.poBoqBodyService.FindByItemNo(input.ItemNo)

	if existingBoqBody.Id != 0 && existingBoqBody.RunNum == input.RunNum {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ItemNo sudah ada di database untuk RunNum yang sama",
		})
		return
	}

	createdPoBoqBody := domain.PoBoqBody{
		RunNum:            input.RunNum,
		ParentId:          input.ParentId,
		ItemNo:            input.ItemNo,
		ItemLevel:         input.ItemLevel,
		ItemDescription:   input.ItemDescription,
		ItemSpecification: input.ItemSpecification,
		Qty:               input.Qty,
		Unit:              input.Unit,
		Price:             input.Price,
		Currency:          input.Currency,
		Note:              input.Note,
	}

	poBoqBodies, err := h.poBoqBodyService.Store(createdPoBoqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data BoQ Body",
		})
		return
	}

	response := domain.PoBoqBodyResponseFinal{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data BoQ Body",
		Data:    []domain.PoBoqBody{poBoqBodies},
	}

	c.JSON(http.StatusCreated, response)
}
