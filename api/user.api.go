package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	"github.com/gin-gonic/gin"
)

func setupUserAPI(router *gin.Engine) {
	authenAPI := router.Group("/api")
	{

		authenAPI.GET("/", hompage)
		authenAPI.POST("/getuser", interceptor.JwtVerify, getuser)
	}
}

func getuser(c *gin.Context) {
	var employee []model.Employee
	db.GetDB().Find(&employee)
	c.JSON(200, gin.H{"data": employee})
}

func hompage(c *gin.Context) {
	//fmt.Println("test")
	c.JSON(200, gin.H{"stus": "ok1"})
}
