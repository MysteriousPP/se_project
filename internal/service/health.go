package service

import (
	"context"
	"se_project/internal/model/input/healthin"
)

type (
	IHealthV1 interface {
		UploadBp(ctx context.Context, req *healthin.UploadBpInp) (res *healthin.UploadBpModel, err error)
		SOSCreate(ctx context.Context, req *healthin.SOSCreateInp) (res *healthin.SOSCreateModel, err error)
	}
)

var (
	LocalHealthV1 IHealthV1
)

func RegisterHealthV1(i IHealthV1) {
	LocalHealthV1 = i
}

func HealthV1() IHealthV1 {
	if LocalHealthV1 == nil {
		panic("implement not found for interface IHealth, forgot register?")
	}
	return LocalHealthV1
}
