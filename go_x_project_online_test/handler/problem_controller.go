package handler

import (
	"github.com/gin-gonic/gin"
	"goStudy/go_x_project_online_test/common"
	"goStudy/go_x_project_online_test/model/dto"
	"goStudy/go_x_project_online_test/service"
	"log"
)

type ProblemHandler struct {
}

func (*ProblemHandler) QueryPage(c *gin.Context) {
	// 接受 body格式的
	page := common.CalcPage()
	var problemDto = dto.ProblemDto{
		PageDTO: page, //默认分页
	}

	if err := c.ShouldBindJSON(&problemDto); err != nil {
		log.Fatalln(err)
	}
	problemService := service.ProblemService{}
	query, count := problemService.QueryPage(problemDto)
	res := common.PageResult{
		PageDTO: common.PageDTO{PageNum: problemDto.PageNum,
			PageSize: problemDto.PageSize,
			Total:    count},
		List: query,
	}
	common.Success(c, res)
}
