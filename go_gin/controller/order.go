package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Stu struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type OrderController struct {
}

func (o OrderController) GetOrderInfo(context *gin.Context) {
	Success(context, 0, "success", gin.H{
		"name": "Order",
		"age":  23,
	})
}

// 模拟函数名重复

func (o OrderController) List(c *gin.Context) {
	Error(c, 1, "orderError")
}

// 获取post请求参数

func (o OrderController) Post(c *gin.Context) {
	//value := c.PostForm("name")
	//form := c.DefaultPostForm("age", "100")
	// 接受 body格式的
	var stu = Stu{}
	if err := c.ShouldBindJSON(&stu); err != nil {
		fmt.Println(err)
	}
	fmt.Println(stu)
	Success(c, 0, "success", stu)
}
