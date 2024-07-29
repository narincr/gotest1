package api

import (
	"Test1/db"
	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {
	db.SetupDB()

	SetupIndexPage(router)
	SetupMember(router)
}
