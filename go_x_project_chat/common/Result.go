package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int         `json:"code"` //返回到前端时 小写
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
func Error(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
	})
}
