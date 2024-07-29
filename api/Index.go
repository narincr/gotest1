package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupIndexPage(router *gin.Engine) {
	ApiControl := router.Group("/Index")
	{
		ApiControl.GET("/Check", IndexCheck)
		ApiControl.GET("/Check2", Index2)
	}
}

func IndexCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Result": true, "Remark": "CheckOK"})
}

func Index2(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"Result": true, "Remark": "Index2"})
}
