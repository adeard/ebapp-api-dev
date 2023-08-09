package boqbody

import (
	"ebapp-api-dev/domain"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type boqBodyHandler struct {
	boqBodyService Service
}

func NewBoqBodyHandler(v1 *gin.RouterGroup, boqBodyService Service) {

	handler := &boqBodyHandler{boqBodyService}

	boqBody := v1.Group("boq_body")

	boqBody.GET("", handler.GetAll)                      // FOR TEST PURPOSE ONLY
	boqBody.GET("/tree/:id", handler.GetChildByParentId) // FOR TEST PURPOSE ONLY
	boqBody.GET("/:id", handler.GetBoqByRunNum)
	boqBody.POST("", handler.Store)
	boqBody.PUT("/:id", handler.Update)
	boqBody.DELETE("/:id", handler.Delete)
}

func groupItemsByParent(items []domain.BoqBodyResponse, parentId int) []domain.BoqBodyResponse {
	var result []domain.BoqBodyResponse

	for _, item := range items {
		if item.ParentId == parentId {
			children := groupItemsByParent(items, item.Id)
			item.Children = children
			result = append(result, item)
		}
	}

	return result
}

func (h *boqBodyHandler) GetAll(c *gin.Context) {
	boqBody, err := h.boqBodyService.GetAll(domain.BoqBodyRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

	// Konversi tipe data []domain.BoqBody menjadi []domain.BoqBodyResponse
	var boqBodyResponse []domain.BoqBodyResponse
	for _, body := range boqBody {
		boqBodyResponse = append(boqBodyResponse, domain.BoqBodyResponse{
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

	result := groupItemsByParent(boqBodyResponse, 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data BoQ Body",
		"data":    result,
	})
}

func (h *boqBodyHandler) GetChildByParentId(c *gin.Context) {
	id := c.Param("id")

	boqBody, err := h.boqBodyService.GetByParentId(id)
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

	// Algoritma rekurisnya tabahkan disini
	var result []int
	var findData func(parentID int)
	findData = func(parentID int) {
		boqBody, err := h.boqBodyService.GetByParentId(strconv.Itoa(parentID))
		if err != nil {
			return
		}

		for _, item := range boqBody {
			result = append(result, item.Id)
			findData(item.Id)
		}
	}

	for _, item := range boqBody {
		result = append(result, item.Id)
		findData(item.Id)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil mengambil data BoQ Body",
		"data":    result,
	})
}

func (h *boqBodyHandler) GetBoqByRunNum(c *gin.Context) {
	runNum := c.Param("id")

	boqBody, err := h.boqBodyService.GetByRunNum(runNum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
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

	// Konversi tipe data []domain.BoqBody menjadi []domain.BoqBodyResponse
	var boqBodyResponse []domain.BoqBodyResponse
	for _, body := range boqBody {
		boqBodyResponse = append(boqBodyResponse, domain.BoqBodyResponse{
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

	result := groupItemsByParent(boqBodyResponse, 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Berhasil mengambil data BoQ Body",
		"data":    result,
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
	id := c.Param("id")
	fmt.Println("Ini Id yang didapat : ", id)

	_, err := h.boqBodyService.GetByParentId(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Gagal mengambil data BoQ Body",
			"data":    nil,
		})
		return
	}

	// boqBody, err := h.boqBodyService.GetByParentId(id)
	// if err != nil {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  http.StatusInternalServerError,
	// 		"message": "Gagal mengambil data BoQ Body",
	// 		"data":    nil,
	// 	})
	// 	return
	// }

	// if len(boqBody) == 0 {
	// 	c.JSON(http.StatusNotFound, gin.H{
	// 		"status":  http.StatusNotFound,
	// 		"message": "Data BoQ Body tidak ditemukan",
	// 		"data":    nil,
	// 	})
	// 	return
	// }

	// Algoritma rekurisnya
	var result []int
	parentID, _ := strconv.Atoi(id)
	result = append(result, parentID)

	var findData func(parentID int)
	findData = func(parentID int) {
		boqBody, err := h.boqBodyService.GetByParentId(strconv.Itoa(parentID))
		if err != nil {
			return
		}

		for _, item := range boqBody {
			result = append(result, item.Id)
			findData(item.Id)
		}
	}

	findData(parentID)

	//buatkan funsi delete dibawah ini berdasarkan id yang ada di result -> contoh hasil result [19,20,21,22,23]
	// Hapus data berdasarkan ID yang ada di result
	for _, id := range result {
		err := h.boqBodyService.DeleteByID(id)
		if err != nil {
			// Jika ada kesalahan saat menghapus, tangani sesuai kebutuhan (misalnya kembalikan pesan kesalahan)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Gagal menghapus data BoQ Body",
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Berhasil menhapus data BoQ Body",
	})
}
