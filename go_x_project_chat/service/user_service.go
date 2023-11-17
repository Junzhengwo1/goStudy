package service

import (
	"goStudy/go_x_project_chat/dao"
	"goStudy/go_x_project_chat/model/vo"
)

type UserService struct {
}

func (*UserService) Query(id int) vo.UserVo {
	user := (&dao.UserDao{}).Query(id)
	// user 转 userVo
	return vo.UserVo{
		ClientIp:   user.ClientIp,
		LogoutTime: user.LogoutTime,
		Email:      user.Email,
		ClientPort: user.ClientPort,
		LoginTime:  user.LoginTime,
		Id:         user.Id,
		DeviceInfo: user.DeviceInfo,
		Password:   user.Password,
		Phone:      user.Phone,
		Username:   user.Username,
	}
}
