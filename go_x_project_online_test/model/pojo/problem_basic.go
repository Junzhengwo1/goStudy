package pojo

import (
	"gorm.io/gorm"
)

type ProblemBasic struct {
	ID                uint               `gorm:"primary_key;" json:"id"`
	CreatedAt         MyTime             `json:"created_at"`
	UpdatedAt         MyTime             `json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `gorm:"index;" json:"deleted_at"`
	Identity          string             `gorm:"column:identity;type:varchar(36);" json:"identity"`                 // 问题表的唯一标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:id" json:"problemCategories"`      // 关联问题分类表
	Title             string             `gorm:"column:title;type:varchar(255);" json:"title"`                      // 文章标题
	Content           string             `gorm:"column:content;type:text;" json:"content"`                          // 文章正文
	MaxRuntime        int                `gorm:"column:max_runtime;type:int(11);" json:"maxRuntime"`                // 最大运行时长
	MaxMem            int                `gorm:"column:max_mem;type:int(11);" json:"maxMem"`                        // 最大运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:problem_identity;references:identity;" json:"testCases"` // 关联测试用例表
	PassNum           int64              `gorm:"column:pass_num;type:int(11);" json:"passNum"`                      // 通过次数
	SubmitNum         int64              `gorm:"column:submit_num;type:int(11);" json:"submitNum"`                  // 提交次数
}

func (table *ProblemBasic) TableName() string {
	return "problem_basic"
}
