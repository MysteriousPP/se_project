package healthin

import "github.com/gogf/gf/v2/os/gtime"

type UploadBpInp struct {
	UserId      int         `v:"required"`
	HighBp      int         `v:"required"`
	LowBp       int         `v:"required"`
	MeasureTime *gtime.Time `json:"time"`
	UploadType  int8
	Value       string `json:"value"`
}

type UploadBpModel struct {
	DataId      int         `json:"recordId"`
	ElderId     int         `json:"elderId"`
	MeasureTime *gtime.Time `json:"time"`
	Value       string      `json:"value"`
	IsAbnormal  int
}
