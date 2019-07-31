package search

import (
	"search/internal/search/api"

	"github.com/gin-gonic/gin"
)

func RouterMain(router *gin.Engine) {
	api.Routes(router)
}
