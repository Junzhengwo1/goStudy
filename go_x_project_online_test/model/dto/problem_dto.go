package dto

type ProblemDto struct {
	PageDTO
	Title   string `json:"title"` // 文章标题
	Content string `json:"content"`
}
