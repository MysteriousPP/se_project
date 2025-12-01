package service

import (
	"context"
	"se_project/internal/model/input/svcin"
)

type (
	ISvcV1 interface {
		Create(ctx context.Context, req *svcin.ServiceOrderCreateInp) (res *svcin.ServiceOrderCreateModel, err error)
	}
)

var (
	LocalSvcV1 ISvcV1
)

func RegisterSvcV1(i ISvcV1) {
	LocalSvcV1 = i
}

func SvcV1() ISvcV1 {
	if LocalSvcV1 == nil {
		panic("implement not found for interface SvcV1, forgot register?")
	}
	return LocalSvcV1
}
