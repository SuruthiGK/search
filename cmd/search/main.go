package main

import (
	"fmt"
	"runtime"
	"search/internal/search"

	"github.com/gin-gonic/gin"
)

//main --gin Engine instance is created
//Routes to the end-points
//host and port is specified for the app
func main() {
	fmt.Println("server is up and running!!")
	runtime.GOMAXPROCS(4)

	app := gin.Default()

	search.RouterMain(app)

	err := app.Run("0.0.0.0:5000")
	if err != nil {
		panic(err)
	}
	fmt.Println("server got fired!!!!")
}
