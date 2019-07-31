package main

import (
	"fmt"
	"runtime"
	"search/internal/search"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("server started!!")
	runtime.GOMAXPROCS(4)

	// Config goes here
	app := gin.Default()

	//Routes
	search.RouterMain(app)

	err := app.Run("0.0.0.0:5000")
	if err != nil {
		panic(err)
	}
	//search.SearchIntegersFromString()
	fmt.Println("server is up and running!!")
}
