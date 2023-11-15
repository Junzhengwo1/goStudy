package router

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_xproject_chat/controller"
)

func Router() *gin.Engine {
	r := gin.Default() // 拿到默认引擎 默认路由
	// 路由分组
	userGroup := r.Group("user")
	{
		userController := controller.UserController{}
		userGroup.GET("query", userController.Query)
	}
	return r
}
