package healthin

import "github.com/gogf/gf/v2/os/gtime"

type UploadBpInp struct {
	UserId      int `v:"required"`
	HighBp      int `v:"required"`
	LowBp       int `v:"required"`
	MeasureTime *gtime.Time
	UploadType  int8 `v:"required"`
}

type UploadBpModel struct {
	DataId     int
	IsAbnormal int
}
