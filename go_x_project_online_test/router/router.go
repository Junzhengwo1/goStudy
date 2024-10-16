package router

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_x_project_online_test/handler"
)

func Router() *gin.Engine {
	r := gin.Default() // 拿到默认引擎 默认路由
	// 配置 日志收集 暂时先注释掉
	//r.Use(gin.LoggerWithConfig(self_logger.LoggerToFile()))
	//r.Use(self_logger.Recover)
	// 路由分组
	userGroup := r.Group("user")
	{
		// 这里方法绑定在 指针上 所以要用取地址符号
		userGroup.GET("query", (&handler.UserHandler{}).Query)
	}

	// 路由分组
	problemGroup := r.Group("problem")
	{
		problemHandler := handler.ProblemHandler{}
		problemGroup.POST("query", problemHandler.QueryPage)
	}

	return r
}
