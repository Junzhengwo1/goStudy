package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/**
gin框架:
		1.
		2.
*/

type Person struct {
	Name string //属性大写
	Age  int
}

func ping(c *gin.Context) {
	data := gin.H{"name": "king", "age": 18}
	c.JSON(http.StatusOK, data)
}

// 返回文件
func file(c *gin.Context) {
	c.File("./file/123.jpeg")
}

func main() {

	r := gin.Default() // 拿到默认引擎 默认路由
	// 前面这个是 浏览器访问的地址  第二个参数是 项目目录
	r.StaticFile("st/file", "./file/123.jpeg") // 配置文件可以访问
	r.StaticFS("kou", http.Dir("./file"))      // 配置文件目录可以访问

	r.GET("/ping", ping) //函数本身作为变量传递

	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.GET("/test", func(c *gin.Context) {
		data := Person{"kou", 20}
		c.JSON(http.StatusOK, data)
	})
	r.GET("/file", file) //函数本身作为变量传递

	//------------------------以上是响应-------------------------------

	r.GET("/web", func(c *gin.Context) {
		// 获取浏览器那边携带请求参数 eg localhost:8080/web?name=kou
		name := c.Query("name")
		query, b := c.GetQuery("age")
		if !b {
			fmt.Println("默认")
		}
		fmt.Println(query)
		c.JSON(http.StatusOK, name)
	})

	r.POST("/post", func(c *gin.Context) {
		// 获取 body里面的 参数
		name := c.Query("name")
		c.JSON(http.StatusOK, name)
	})

	// 必须写在最后执行
	err := r.Run(":8088")
	if err != nil {
		return
	} // 监听并在 0.0.0.0:8089 上启动服务

}
