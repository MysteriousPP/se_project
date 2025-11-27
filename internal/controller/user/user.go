package user

import (
	"context"
	v1 "se_project/api/user/v1"
	"se_project/internal/service"
)

type ControllerUser struct{}

var User = ControllerUser{}

func (c *ControllerUser) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	data, err := service.UserV1().Create(ctx, &req.UserCreateInp)
	if err != nil {
		return nil, err
	}
	res = &v1.CreateRes{}
	res.UserCreateModel = data
	return
}

func (c *ControllerUser) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	data, err := service.UserV1().Login(ctx, &req.UserLoginInp)
	if err != nil {
		return nil, err
	}
	res = &v1.LoginRes{}
	res.UserLoginModel = data
	return
}

func (c *ControllerUser) BindElder(ctx context.Context, req *v1.BindElderReq) (res *v1.BindElderRes, err error) {
	data, err := service.UserV1().BindElder(ctx, &req.BindElderInp)
	if err != nil {
		return nil, err
	}
	res = &v1.BindElderRes{}
	res.BindElderModel = data
	return
}
