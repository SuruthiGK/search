package api

import (
	"github.com/gin-gonic/gin"
)

//Routes takes the request to the end-point
func Routes(route *gin.Engine) {
	search := route.Group("/api/search")
	{
		search.POST("/integers", SearchIntegersFromString)
	}
}
