package vo

import (
	"goStudy/go_x_project_online_test/model/pojo"
	"gorm.io/gorm"
)

type UserVo struct {
	ID        uint           `gorm:"primary_key;" json:"id"`
	CreatedAt pojo.MyTime    `gorm:"type:timestamp;" json:"created_at"`
	UpdatedAt pojo.MyTime    `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index;" json:"deleted_at"`
	Identity  string         `gorm:"column:identity;type:varchar(36);" json:"identity"` // 用户的唯一标识
	Name      string         `gorm:"column:name;type:varchar(100);" json:"name"`        // 用户名
	Password  string         `gorm:"column:password;type:varchar(32);" json:"-"`        // 密码
	Phone     string         `gorm:"column:phone;type:varchar(20);" json:"phone"`       // 手机号
	Mail      string         `gorm:"column:mail;type:varchar(100);" json:"mail"`        // 邮箱
	PassNum   int64          `gorm:"column:pass_num;type:int(11);" json:"passNum"`      // 通过的次数
	SubmitNum int64          `gorm:"column:submit_num;type:int(11);" json:"submitNum"`  // 提交次数
	IsAdmin   int            `gorm:"column:is_admin;type:tinyint(1);" json:"isAdmin"`
}
