package dto

import "goStudy/go_x_project_online_test/common"

type ProblemDto struct {
	common.PageDTO
	Title   string `json:"title"` // 文章标题
	Content string `json:"content"`
}
