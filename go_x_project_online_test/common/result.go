package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Code int    `json:"code"` //返回到前端时 小写
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

type PageResult struct {
	PageDTO
	List interface{} `json:"list"`
}

func Success(c *gin.Context, data any) {
	c.JSON(http.StatusOK, Result{
		Code: SuccessCode,
		Msg:  SuccessMsg,
		Data: data,
	})
}

func SuccessWithCode(c *gin.Context, code int, msg string, data any) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Result{
		Code: code,
		Msg:  msg,
	})
}
