package common

type PageDTO struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}

// 计算分页 默认值为1,10

func CalcPage() PageDTO {
	dto := PageDTO{
		PageNum:  PageNum,
		PageSize: PageSize,
	}
	return dto
}
