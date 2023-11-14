package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

func (u OrderController) GetOrderInfo(c *gin.Context) {
	Success(c, 0, "success", gin.H{
		"name": "Order",
		"age":  23,
	})
}

// 模拟函数名重复

func (u OrderController) List(c *gin.Context) {
	Error(c, 1, "orderError")
}

// 获取post请求参数

func (u OrderController) Post(c *gin.Context) {
	value := c.PostForm("name")
	form := c.DefaultPostForm("age", "100")
	fmt.Println(value, form)
	Success(c, 0, "success", gin.H{
		"name": "Order",
		"age":  23,
	})
}
