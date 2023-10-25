package poprogressheader

import "github.com/gin-gonic/gin"

type poProgressHeaderHandler struct {
	poProgressHeaderService Service
}

func NewPoProgressHeaderHandler(v1 *gin.RouterGroup, poProgressHeaderService Service) {

}
