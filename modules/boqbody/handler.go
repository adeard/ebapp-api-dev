package boqbody

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type boqBodyHandler struct {
	boqBodyService Service
}

func NewBoqBodyHandler(v1 *gin.RouterGroup, boqBodyService Service) {

	handler := &boqBodyHandler{boqBodyService}

	boqBody := v1.Group("boq_body")

	boqBody.GET("", handler.GetAll)
	boqBody.POST("", handler.Store)
}

func (h *boqBodyHandler) GetAll(c *gin.Context) {
	var input domain.BoqBodyRequest

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boqBody, err := h.boqBodyService.GetAll(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})
	}

	res := []domain.BoqBodyResponse{}

	ParentId := 0

	HighestLevel := 1

	for _, boqBodyData := range boqBody {
		if res == nil || boqBodyData.ItemLevel == 1 {
			res = append(res, domain.BoqBodyResponse{
				Id:                boqBodyData.Id,
				RunNum:            boqBodyData.RunNum,
				ItemNo:            boqBodyData.ItemNo,
				ItemLevel:         boqBodyData.ItemLevel,
				ItemDescription:   boqBodyData.ItemDescription,
				ItemSpecification: boqBodyData.ItemSpecification,
				Qty:               boqBodyData.Qty,
				Unit:              boqBodyData.Unit,
				Price:             boqBodyData.Price,
				Currency:          boqBodyData.Currency,
				Note:              boqBodyData.Note,
				Children:          nil,
				ParentId:          0,
			})

			continue
		}

		previousValue := res[len(res)-1]

		if previousValue.ItemLevel < boqBodyData.ItemLevel {
			if HighestLevel <= boqBodyData.ItemLevel {
				HighestLevel = boqBodyData.ItemLevel
			}

			ParentId = previousValue.Id
		}

		if previousValue.ItemLevel == boqBodyData.ItemLevel {
			ParentId = previousValue.ParentId
		}

		if previousValue.ItemLevel > boqBodyData.ItemLevel {
			var ParentBefore int
			for _, parentData := range res {
				if parentData.Id == previousValue.ParentId {
					ParentBefore = parentData.ParentId

					break
				}
			}

			ParentId = ParentBefore
		}

		res = append(res, domain.BoqBodyResponse{
			Id:                boqBodyData.Id,
			RunNum:            boqBodyData.RunNum,
			ItemNo:            boqBodyData.ItemNo,
			ItemLevel:         boqBodyData.ItemLevel,
			ItemDescription:   boqBodyData.ItemDescription,
			ItemSpecification: boqBodyData.ItemSpecification,
			Qty:               boqBodyData.Qty,
			Unit:              boqBodyData.Unit,
			Price:             boqBodyData.Price,
			Currency:          boqBodyData.Currency,
			Note:              boqBodyData.Note,
			Children:          nil,
			ParentId:          ParentId,
		})

		continue

	}

	for i := HighestLevel; i > 0; i-- {
		for _, resData := range res {
			if resData.ParentId != 0 && resData.ItemLevel == i {
				for index, parentTemp := range res {
					if parentTemp.Id == resData.ParentId {
						parentTemp.Children = append(parentTemp.Children, resData)
						res[index] = parentTemp
						break
					}
				}
			}
		}
	}

	resultFix := []domain.BoqBodyResponse{}

	for _, final := range res {
		if final.ItemLevel == 1 {
			resultFix = append(resultFix, final)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"data": resultFix,
	})
}

func (h *boqBodyHandler) Store(c *gin.Context) {
	var input domain.BoqBodyRequest

	c.ShouldBindJSON(&input)

	validate := validator.New()
	err := validate.Struct(input)
	if err != nil {
		errorMessages := []interface{}{}

		for _, v := range err.(validator.ValidationErrors) {
			errorArray := map[string]string{
				"field":   v.Field(),
				"message": v.ActualTag(),
			}

			errorMessages = append(errorMessages, errorArray)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})

		return
	}

	ztsTravis, err := h.boqBodyService.Store(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"errors ": err,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": ztsTravis,
	})
}
