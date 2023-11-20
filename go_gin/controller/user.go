package Handler

import "github.com/gin-gonic/gin"

type UserHandler struct {
}

func (UserHandler) GetUserInfo(c *gin.Context) {
	Success(c, 0, "success", gin.H{
		"name": "test router",
		"age":  23,
	})

}

func (UserHandler) List(c *gin.Context) {
	Error(c, 1, "userError")
}
