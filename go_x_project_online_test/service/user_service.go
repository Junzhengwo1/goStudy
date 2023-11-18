package service

import (
	"goStudy/go_x_project_online_test/dao"
	"goStudy/go_x_project_online_test/model/pojo"
	"goStudy/go_x_project_online_test/model/vo"
)

type UserService struct {
}

func (*UserService) Query(id int) vo.UserVo {
	user := (&dao.UserDao{}).Query(id)
	// user 转 userVo
	return buildData(user)
}

func buildData(user *pojo.UserBasic) vo.UserVo {
	return vo.UserVo{
		ID:        user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Mail:      user.Mail,
		SubmitNum: user.SubmitNum,
		PassNum:   user.PassNum,
		DeletedAt: user.DeletedAt,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		IsAdmin:   user.IsAdmin,
		Identity:  user.Identity,
	}
}
