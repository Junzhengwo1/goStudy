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
		// 这里方法绑定在 指针上 所以要用取地址符号
		userGroup.GET("query", (&controller.UserController{}).Query)
	}
	return r
}
