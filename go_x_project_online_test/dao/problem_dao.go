package dao

import (
	"fmt"
	"goStudy/go_x_project_online_test/config"
	"goStudy/go_x_project_online_test/model/dto"
	"goStudy/go_x_project_online_test/model/pojo"
)

type ProblemDao struct {
}

func (a *ProblemDao) QueryPage(param dto.ProblemDto) ([]pojo.ProblemBasic, int64) {
	var problemList []pojo.ProblemBasic
	title := param.Title
	var count int64
	// 分页查询
	affected := config.Db.Model(&pojo.ProblemBasic{}).
		Where("title like ? OR content like ?", "%"+title+"%", "%"+title+"%").
		Count(&count).
		Offset(param.CalcPageNum()).
		Limit(param.PageDTO.PageSize).
		Find(&problemList).RowsAffected
	fmt.Println("affected:", affected)
	fmt.Println("count:", count)
	fmt.Println("problemList:", problemList)
	return problemList, count
}
