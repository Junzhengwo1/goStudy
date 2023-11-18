package service

import (
	"fmt"
	"goStudy/go_x_project_online_test/dao"
	"goStudy/go_x_project_online_test/model/dto"
	"goStudy/go_x_project_online_test/model/vo"
)

type ProblemService struct {
}

func (*ProblemService) QueryPage(param dto.ProblemDto) []vo.ProblemVo {
	page := (&dao.ProblemDao{}).QueryPage(param)
	fmt.Println(page)
	return []vo.ProblemVo{}
}
