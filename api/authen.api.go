package api

import (
	"main/db"
	"main/model"

	"github.com/gin-gonic/gin"
) 

func setupAuthenAPI(router *gin.Engine){
	authenAPI := router.Group("/api/v2")
	{

		authenAPI.POST("/login", login)
	
		authenAPI.POST("/getuser", getuser)
	}
}
func login(c *gin.Context)  {
	var employees  model.Employee
	if c.ShouldBind(&employees) == nil{
		c.JSON(200, gin.H{"data":employees})
	}

}
func getuser(c *gin.Context){
	var employee []model.Employee
	db.GetDB().Find(&employee)
	c.JSON(200, gin.H{"data":employee})
}	

func sex(c *gin.Context){
	var employee []model.Employee
	db.GetDB().Find(&employee)
	c.JSON(200, gin.H{"data":employee})
}
