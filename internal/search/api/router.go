package api

import (
	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	search := route.Group("/api/search")
	{
		search.POST("/integers", SearchIntegersFromString)
	}
}
