package user

import (
	"context"
	"se_project/internal/dao"
	"se_project/internal/lib/crypto"
	"se_project/internal/lib/jwt"
	"se_project/internal/model/do"
	"se_project/internal/model/input/userin"
	"se_project/internal/service"
	"strconv"

	"github.com/gogf/gf/v2/errors/gerror"
)

type UserV1 struct{}

func NewUserV1() *UserV1 {
	return &UserV1{}
}

func init() {
	service.RegisterUserV1(NewUserV1())
}

func (u *UserV1) Create(ctx context.Context, req *userin.UserCreateInp) (res *userin.UserCreateModel, err error) {
	cls := dao.User.Columns()

	count, err := dao.User.Ctx(ctx).Where(cls.Phone, req.Phone).Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, gerror.New("account exists")
	}
	crypted_password, _ := crypto.SHA256WithSalt(req.Password)
	_, err = dao.User.Ctx(ctx).Data(do.User{
		Phone:    req.Phone,
		Password: crypted_password,
		Role:     req.Role,
		Nickname: req.Nickname,
	}).Insert()
	if err != nil {
		return nil, err
	}
	return
}
func (u *UserV1) Login(ctx context.Context, req *userin.UserLoginInp) (res *userin.UserLoginModel, err error) {
	res = &userin.UserLoginModel{}
	cls := dao.User.Columns()
	type IdAndPswd struct {
		Id       int
		Password string
	}
	data := &IdAndPswd{}
	err = dao.User.Ctx(ctx).Where(cls.Phone, req.Phone).Scan(data)
	if err != nil {
		return nil, err
	}
	err = dao.User.Ctx(ctx).Where(cls.Phone, req.Phone).Scan(res)
	if err != nil {
		return nil, err
	}
	isEqual, err := crypto.VerifySHA256WithSalt(req.Password, data.Password)
	if err != nil {
		return nil, err
	}
	if !isEqual {
		return nil, gerror.New("password wrong")
	}
	tokenString, err := jwt.GenAccessToken(strconv.Itoa(data.Id))
	if err != nil {
		return nil, err
	}
	res.Token = tokenString
	res.UserId = data.Id
	return res, nil

}
func (h *UserV1) BindElder(ctx context.Context, req *userin.BindElderInp) (res *userin.BindElderModel, err error) {
	res = &userin.BindElderModel{}
	// cls := dao.BindRelation.Columns()
	_, err = dao.BindRelation.Ctx(ctx).Data(
		do.BindRelation{
			ChildUserID: req.ChildUserId,
			ElderUserId: req.ElderUserId,
			Status:      1,
		}).Insert()
	if err != nil {
		return nil, err
	}
	return res, nil
}
