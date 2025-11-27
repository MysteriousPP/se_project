package service

import (
	"context"
	"se_project/internal/model/input/userin"
)

type (
	IUserV1 interface {
		Create(ctx context.Context, req *userin.UserCreateInp) (res *userin.UserCreateModel, err error)
		Login(ctx context.Context, req *userin.UserLoginInp) (res *userin.UserLoginModel, err error)
		BindElder(ctx context.Context, req *userin.BindElderInp) (res *userin.BindElderModel, err error)
	}
)

var (
	LocalUserV1 IUserV1
)

func RegisterUserV1(i IUserV1) {
	LocalUserV1 = i
}

func UserV1() IUserV1 {
	if LocalUserV1 == nil {
		panic("implement not found for interface UserV1, forgot register?")
	}
	return LocalUserV1
}
