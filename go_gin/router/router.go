package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goStudy/go_gin/controller"
	"goStudy/go_gin/util/self_logger"
	"net/http"
)

func UserRouter() *gin.Engine {
	r := gin.Default() // 拿到默认引擎 默认路由
	// 配置 日志收集
	r.Use(gin.LoggerWithConfig(self_logger.LoggerToFile()))
	r.Use(gin.Recovery())

	// 路由分组
	userGroup := r.Group("user")
	{
		userGroup.GET("get", func(c *gin.Context) {
			data := gin.H{"name": "king", "age": 18}
			c.JSON(http.StatusOK, data)
		})
		userGroup.GET("query", func(c *gin.Context) {
			name := c.Query("name")
			query, b := c.GetQuery("age")
			if !b {
				fmt.Println("默认")
			}
			fmt.Println(query)
			c.JSON(http.StatusOK, name)
		})
	}
	userGroup.GET("test/router", controller.UserController{}.GetUserInfo)
	userGroup.GET("error", controller.UserController{}.List)

	orderGroup := r.Group("order")
	orderGroup.GET("order", controller.OrderController{}.GetOrderInfo)
	orderGroup.GET("error", controller.OrderController{}.List)
	orderGroup.POST("post", controller.OrderController{}.Post)

	goGroup := r.Group("go")
	goGroup.GET("go", controller.GoTestController{}.FindOne)

	return r
}
