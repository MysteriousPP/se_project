package svc

import (
	"context"
	v1 "se_project/api/svc/v1"
	"se_project/internal/service"
)

type ControllerSvc struct{}

var Svc = ControllerSvc{}

func (s *ControllerSvc) Create(ctx context.Context, req *v1.ServiceOrderCreateReq) (res *v1.ServiceOrderCreateRes, err error) {
	data, err := service.SvcV1().Create(ctx, &req.ServiceOrderCreateInp)
	if err != nil {
		return nil, err
	}
	res = &v1.ServiceOrderCreateRes{}
	res.ServiceOrderCreateModel = data
	return res, nil
}
