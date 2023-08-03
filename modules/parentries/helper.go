package parentries

import "github.com/gin-gonic/gin"

type parEntriesHandler struct {
	parEntriesService Service
}

func NewParEntriesHandler(v1 *gin.RouterGroup, parEntriesService Service) {
	handler := &parEntriesHandler{parEntriesService}

	parEntries := v1.Group("par_entries")

	parEntries.GET("", handler.GetAll)
}

func (h *parEntriesHandler) GetAll(c *gin.Context) {

}
