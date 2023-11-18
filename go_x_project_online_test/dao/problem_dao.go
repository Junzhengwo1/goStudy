package dao

import (
	"fmt"
	"goStudy/go_x_project_online_test/model/dto"
	"goStudy/go_x_project_online_test/model/pojo"
	"goStudy/go_x_project_online_test/util"
)

type ProblemDao struct {
}

func (*ProblemDao) QueryPage(param dto.ProblemDto) []pojo.ProblemBasic {

	var problemList []pojo.ProblemBasic
	title := param.Title
	// 分页查询
	util.Db.Model(pojo.ProblemBasic{}).Where("title like ? OR content like ?", "%"+title+"%", "%"+title+"%").
		Offset((param.PageDTO.PageNum - 1) * param.PageDTO.PageSize).
		Limit(param.PageDTO.PageSize).
		Find(&problemList)
	fmt.Println("problemList:", problemList)

	return problemList
}
