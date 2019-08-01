package api

import (
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

//RequestParams accepts the data in the form of string
type RequestParams struct {
	Data string `json:"data"`
}

//SearchIntegersFromString searches for integers in the provided string
//regex pattern is used to identify the integers in the string
func SearchIntegersFromString(c *gin.Context) {

	var params RequestParams
	result := []string{}

	err := c.BindJSON(&params)

	if err != nil {
		glog.Error(err)
		c.JSON(404, gin.H{"message": "Bad Params."})
		return
	}

	glog.Info("searching integer ...")

	re := regexp.MustCompile("[0-9]+")

	result = re.FindAllString(params.Data, -1)
	glog.Info("result: ", result)

	if len(result) == 0 {
		c.JSON(200, gin.H{"message": "No integer found in the string"})
		return
	}

	c.JSON(200, result)
	return
}
