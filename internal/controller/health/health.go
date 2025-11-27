package health

import (
	"context"
	v1 "se_project/api/health/v1"
	"se_project/internal/service"
)

type ControllerHealth struct{}

var Health = ControllerHealth{}

func (c *ControllerHealth) UploadBp(ctx context.Context, req *v1.UploadBpReq) (res *v1.UploadBpRes, err error) {
	data, err := service.HealthV1().UploadBp(ctx, &req.UploadBpInp)
	if err != nil {
		return nil, err
	}
	res = &v1.UploadBpRes{}
	res.UploadBpModel = data
	return
}

func (c *ControllerHealth) SOSCreate(ctx context.Context, req *v1.SOSCreateReq) (res *v1.SOSCreateRes, err error) {
	data, err := service.HealthV1().SOSCreate(ctx, &req.SOSCreateInp)
	if err != nil {
		return nil, err
	}
	res = &v1.SOSCreateRes{}
	res.SOSCreateModel = data
	return
}
