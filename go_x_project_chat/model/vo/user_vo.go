package vo

import "time"

type UserVo struct {
	Id         int       `json:"id"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	ClientIp   string    `json:"clientIp"`
	ClientPort string    `json:"clientPort"`
	LoginTime  time.Time `json:"loginTime"`
	LogoutTime time.Time `json:"logoutTime"`
	DeviceInfo string    `json:"deviceInfo"`
}
