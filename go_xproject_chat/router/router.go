package router

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default() // 拿到默认引擎 默认路由
	// 路由分组
	r.Group("user")
	return r
}
