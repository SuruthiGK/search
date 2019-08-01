package search

import (
	"search/internal/search/api"

	"github.com/gin-gonic/gin"
)

//RouterMain is the main router
//routes the request to the api
func RouterMain(router *gin.Engine) {
	api.Routes(router)
}
