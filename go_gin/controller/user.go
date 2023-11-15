package controller

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (UserController) GetUserInfo(c *gin.Context) {
	Success(c, 0, "success", gin.H{
		"name": "test router",
		"age":  23,
	})

}

func (UserController) List(c *gin.Context) {
	Error(c, 1, "userError")
}
