package boqbody

import (
	"ebapp-api-dev/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type boqBodyHandler struct {
	boqBodyService Service
}

func NewBoqBodyHandler(v1 *gin.RouterGroup, boqBodyService Service) {

	handler := &boqBodyHandler{boqBodyService}

	boqBody := v1.Group("boq_body")

	boqBody.GET("", handler.GetAll) // FOR TEST PURPOSE ONLY
	boqBody.GET("/:id", handler.GetByID)
	boqBody.POST("", handler.Store)
	boqBody.PUT("/:id", handler.Update)
	boqBody.DELETE("/:id", handler.Delete)
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
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	res := []domain.BoqBodyResponse{}
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
				ParentId:          boqBodyData.ParentId,
			})

			continue
		}

		previousValue := res[len(res)-1]

		if previousValue.ItemLevel < boqBodyData.ItemLevel {
			if HighestLevel <= boqBodyData.ItemLevel {
				HighestLevel = boqBodyData.ItemLevel
			}
			boqBodyData.ParentId = previousValue.Id
		}

		if previousValue.ItemLevel == boqBodyData.ItemLevel {
			boqBodyData.ParentId = previousValue.ParentId
		}

		if previousValue.ItemLevel > boqBodyData.ItemLevel {
			var ParentBefore int
			for _, parentData := range res {
				if parentData.Id == previousValue.ParentId {
					ParentBefore = parentData.ParentId

					break
				}
			}

			boqBodyData.ParentId = ParentBefore
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
			ParentId:          boqBodyData.ParentId,
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
		"status":  200,
		"message": "Berhasil mengambil data BoQ Body",
		"data":    resultFix,
	})
}

func (h *boqBodyHandler) GetByID(c *gin.Context) {
	runNum := c.Param("id")

	var input domain.BoqBodyRequest

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	boqBody, err := h.boqBodyService.GetByRunNum(runNum)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	if len(boqBody) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Data BoQ Body tidak ditemukan",
			"data":    nil,
		})
		return
	}

	res := []domain.BoqBodyResponse{}
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
				ParentId:          boqBodyData.ParentId,
			})

			continue
		}

		previousValue := res[len(res)-1]

		if previousValue.ItemLevel < boqBodyData.ItemLevel {
			if HighestLevel <= boqBodyData.ItemLevel {
				HighestLevel = boqBodyData.ItemLevel
			}

			boqBodyData.ParentId = previousValue.Id
		}

		if previousValue.ItemLevel == boqBodyData.ItemLevel {
			boqBodyData.ParentId = previousValue.ParentId
		}

		if previousValue.ItemLevel > boqBodyData.ItemLevel {
			var ParentBefore int
			for _, parentData := range res {
				if parentData.Id == previousValue.ParentId {
					ParentBefore = parentData.ParentId

					break
				}
			}

			boqBodyData.ParentId = ParentBefore
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
			ParentId:          boqBodyData.ParentId,
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
		"status":  200,
		"message": "Berhasil mengambil data BoQ Body",
		"data":    resultFix,
	})
}

func (h *boqBodyHandler) Store(c *gin.Context) {
	var input domain.BoqBodyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid",
		})
		return
	}

	existingBoqBody, _ := h.boqBodyService.FindByItemNo(input.ItemNo)

	if existingBoqBody.Id != 0 && existingBoqBody.RunNum == input.RunNum {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ItemNo sudah ada di database untuk RunNum yang sama",
		})
		return
	}

	createdBoqBody := domain.BoqBody{
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

	boqBodies, err := h.boqBodyService.Store(createdBoqBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal meneruskan data BoQ Body",
		})
		return
	}

	response := domain.BoqBodyResponseFinal{
		Status:  http.StatusCreated,
		Message: "Berhasil menyimpan data BoQ Body",
		Data:    []domain.BoqBody{boqBodies},
	}

	c.JSON(http.StatusCreated, response)
}

func (h *boqBodyHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var input domain.BoqBodyRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "Request tidak valid" + err.Error(),
		})
		return
	}

	// Mengecek apakah ItemNo yang baru sudah ada di database untuk RunNum yang sama.
	existingBoqBody, _ := h.boqBodyService.FindByItemNo(input.ItemNo)
	// Jika ItemNo yang ditemukan memiliki Id yang tidak sama dengan Id yang sedang diupdate, berarti ItemNo sudah ada di database untuk RunNum yang sama.
	if existingBoqBody.RunNum == input.RunNum && existingBoqBody.Id != input.Id {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "ItemNo sudah ada di database untuk RunNum yang sama",
		})
		return
	}

	// Membuat objek domain.BoqBody yang akan diupdate.
	updateBoqBody := domain.BoqBody{
		ItemNo:            input.ItemNo,
		ParentId:          input.ParentId,
		ItemDescription:   input.ItemDescription,
		ItemSpecification: input.ItemSpecification,
		Qty:               input.Qty,
		Unit:              input.Unit,
		Price:             input.Price,
		Currency:          input.Currency,
		Note:              input.Note,
	}

	// Memanggil service untuk melakukan update data BoQ Body.
	_, updateErr := h.boqBodyService.Update(updateBoqBody, id)
	if updateErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengupdate data BoQ Body",
		})
		return
	}

	response := domain.BoqBodyResponseFinal{
		Status:  http.StatusOK,
		Message: "Berhasil mengupdate data BoQ Body",
		Data:    nil,
	}

	c.JSON(http.StatusOK, response)
}

func (h *boqBodyHandler) Delete(c *gin.Context) {

}
