package controller

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_x_project_online_test/common"
	"goStudy/go_x_project_online_test/model/dto"
	"goStudy/go_x_project_online_test/service"
	"log"
)

type ProblemController struct {
}

func (*ProblemController) QueryPage(c *gin.Context) {
	// 接受 body格式的
	var problemDto = dto.ProblemDto{
		PageDTO: dto.PageDTO{PageNum: common.PageNum, PageSize: common.PageSize}, //默认分页
	}
	if err := c.ShouldBindJSON(&problemDto); err != nil {
		log.Fatalln(err)
	}
	problemService := service.ProblemService{}
	query := problemService.QueryPage(problemDto)
	common.Success(c, query)
}
