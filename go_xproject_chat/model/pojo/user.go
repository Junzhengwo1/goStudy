package pojo

import (
	"time"
)

type User struct {
	//gorm.Model
	// 自增
	Id         int       `json:"id" gorm:"primary_key"`
	Username   string    `json:"username" gorm:"type:varchar(20);not null"`
	Password   string    `json:"password" gorm:"type:varchar(20);not null"`
	Email      string    `json:"email" gorm:"type:varchar(20);not null"`
	Phone      string    `json:"phone" gorm:"type:varchar(20)"`
	ClientIp   string    `json:"clientIp" gorm:"type:varchar(20)"`
	ClientPort string    `json:"clientPort" gorm:"type:varchar(20)"`
	LoginTime  time.Time `json:"loginTime"`
	LogoutTime time.Time `json:"logoutTime"`
	DeviceInfo string    `json:"deviceInfo" gorm:"type:varchar(20)"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	DeletedAt  time.Time `json:"deletedAt"`
}

// 必须实现 该方法

func (*User) TableName() string {
	return "user"
}
