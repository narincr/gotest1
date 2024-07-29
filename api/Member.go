package api

import (
	"Test1/db"
	"Test1/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetupMember(router *gin.Engine) {
	ApiControl := router.Group("/Member")
	{
		ApiControl.POST("/InfoId", MemberInfoId)
		ApiControl.POST("/List", MemberList)
		ApiControl.POST("/Save", MemberSave)
	}
}

func MemberSave(c *gin.Context) {
	var queryGet model.MemberGetSave
	if c.BindJSON(&queryGet) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "BindJSON Error", "Error": "BindJSON Error", "PostData": queryGet})
		return
	}

	if queryGet.FirstName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "Empty FirstName"})
		return
	}
	if queryGet.LastName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "Empty LastName"})
		return
	}
	if queryGet.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "Empty Phone"})
		return
	}

	var CheckRepeat int64
	db.GetDB().Table("Member").Where("FirstName = ?", queryGet.FirstName).Count(&CheckRepeat)
	if CheckRepeat > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "FirstName already exists"})
		return
	}

}

func MemberList(c *gin.Context) {
	var queryGet model.MemberGetList
	if c.BindJSON(&queryGet) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "BindJSON Error", "Error": "BindJSON Error", "PostData": queryGet})
		return
	}

	var queryInfo []model.MemberInfo
	result := db.GetDB().Table("Member")
	if queryGet.FirstName != "" {
		result.Where("FirstName LIKE ?", "%"+queryGet.FirstName+"%")
	}
	if queryGet.LastName != "" {
		result.Where("LastName LIKE ?", "%"+queryGet.LastName+"%")
	}
	if queryGet.Phone != "" {
		result.Where("Phone LIKE ?", "%"+queryGet.Phone+"%")
	}
	result.Order("FirstName asc")
	result.Find(&queryInfo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "Error", "Error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Result": true, "Remark": "Success", "Count": len(queryInfo), "Data": queryInfo})

}

func MemberInfoId(c *gin.Context) {
	var queryGet model.MemberGetId
	if c.BindJSON(&queryGet) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "BindJSON Error", "Error": "BindJSON Error", "PostData": queryGet})
		return
	}

	var queryInfo model.MemberInfo
	result := db.GetDB().Table("Member").Where("MemberID=?", queryGet.Id)
	result.First(&queryInfo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Result": false, "Remark": "Error", "Error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Result": true, "Remark": "Success", "Data": queryInfo})
}
