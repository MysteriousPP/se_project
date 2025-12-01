package svc

import (
	"context"
	"se_project/internal/dao"
	"se_project/internal/model/do"
	"se_project/internal/model/input/svcin"
	"se_project/internal/service"
)

type SvcV1 struct{}

func NewSvcV1() *SvcV1 {
	return &SvcV1{}
}

func init() {
	service.RegisterSvcV1(NewSvcV1())
}

func (s *SvcV1) Create(ctx context.Context, req *svcin.ServiceOrderCreateInp) (res *svcin.ServiceOrderCreateModel, err error) {
	result, err := dao.ServiceBooking.Ctx(ctx).Data(
		do.ServiceBooking{
			ServiceId:   req.ServiceId,
			ElderUserId: req.ElderUserId,
			BookedBy:    req.BookedBy,
			Status:      1,
		}).Insert()
	if err != nil {
		return nil, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	res = &svcin.ServiceOrderCreateModel{}
	res.BookingId = int(lastInsertId)
	res.Status = 1
	return res, nil
}
