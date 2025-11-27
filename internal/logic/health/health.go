package health

import (
	"context"
	"se_project/internal/dao"
	"se_project/internal/model/do"
	"se_project/internal/model/input/healthin"
	"se_project/internal/service"
)

type HealthV1 struct{}

func NewHealthV1() *HealthV1 {
	return &HealthV1{}
}

func init() {
	service.RegisterHealthV1(NewHealthV1())
}

func (h *HealthV1) UploadBp(ctx context.Context, req *healthin.UploadBpInp) (res *healthin.UploadBpModel, err error) {
	_, err = dao.HealthData.Ctx(ctx).Data(do.HealthData{
		UserID:      req.UserId,
		HighBp:      req.HighBp,
		LowBp:       req.LowBp,
		MeasureTime: req.MeasureTime,
		UploadType:  req.UploadType,
	}).Insert()
	// TODO 判断是否异常的依据？
	if err != nil {
		return nil, err
	}
	return
}

func (h *HealthV1) SOSCreate(ctx context.Context, req *healthin.SOSCreateInp) (res *healthin.SOSCreateModel, err error) {
	_, err = dao.SosAlert.Ctx(ctx).Data(do.SosAlert{
		ElderUserId: req.ElderUserId,
		Longitude:   req.Longitude,
		Latitude:    req.Latitude,
		Notes:       req.Notes,
	}).Insert()
	// TODO 是否需要后续操作？
	if err != nil {
		return nil, err
	}
	return
}
