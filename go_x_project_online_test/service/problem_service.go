package service

import (
	"goStudy/go_x_project_online_test/dao"
	"goStudy/go_x_project_online_test/model/dto"
	"goStudy/go_x_project_online_test/model/vo"
)

type ProblemService struct {
}

func (*ProblemService) QueryPage(param dto.ProblemDto) ([]vo.ProblemVo, int64) {
	page, count := (&dao.ProblemDao{}).QueryPage(param)
	var result []vo.ProblemVo
	for _, v := range page {
		result = append(result, vo.ProblemVo{
			ID:                v.ID,
			Title:             v.Title,
			Content:           v.Content,
			SubmitNum:         v.SubmitNum,
			PassNum:           v.PassNum,
			MaxMem:            v.MaxMem,
			MaxRuntime:        v.MaxRuntime,
			Identity:          v.Identity,
			TestCases:         v.TestCases,
			ProblemCategories: v.ProblemCategories,
		})

	}

	return result, count
}
