package v1

import (
	"se_project/internal/model/input/userin"

	"github.com/gogf/gf/v2/frame/g"
)

type CreateReq struct {
	g.Meta `path:"/user" method:"post" tags:"User" summary:"Create user"`
	userin.UserCreateInp
}

type CreateRes struct {
	*userin.UserCreateModel
}

type LoginReq struct {
	g.Meta `path:"/user/login" method:"post" tags:"User" summary:"user login"`
	userin.UserLoginInp
}

type LoginRes struct {
	*userin.UserLoginModel
}

type BindElderReq struct {
	g.Meta `path:"/user/bindElder" method:"post" tags:"User" summary:"bind elder"`
	userin.BindElderInp
}

type BindElderRes struct {
	*userin.BindElderModel
}
