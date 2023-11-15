package pojo

type User struct {
	//gorm.Model
	// 自增
	Id         int    `json:"id" gorm:"primary_key"`
	Username   string `json:"username" gorm:"type:varchar(20);not null"`
	Password   string `json:"password" gorm:"type:varchar(20);not null"`
	Email      string `json:"email" gorm:"type:varchar(20);not null"`
	Phone      string `json:"phone" gorm:"type:varchar(20)"`
	ClientIp   string `json:"client_ip" gorm:"type:varchar(20)"`
	ClientPort string `json:"client_port" gorm:"type:varchar(20)"`
	LoginTime  string `json:"login_time" gorm:"type:varchar(20)"`
	LogoutTime string `json:"logout_time" gorm:"type:varchar(20)"`
	DeviceInfo string `json:"device_info" gorm:"type:varchar(20)"`
}

// 必须实现 该方法

func (*User) TableName() string {
	return "user"
}
