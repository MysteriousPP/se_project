// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// HealthData is the golang structure for table health_data.
type HealthData struct {
	DataID      int         `json:"dataID"      orm:"dataID"      description:"数据唯一ID"`       // 数据唯一ID
	UserID      int         `json:"userID"      orm:"userID"      description:"银发端用户ID"`      // 银发端用户ID
	HighBp      int         `json:"highBp"      orm:"highBp"      description:"高压值(mmHg)"`    // 高压值(mmHg)
	LowBp       int         `json:"lowBp"       orm:"lowBp"       description:"低压值(mmHg)"`    // 低压值(mmHg)
	MeasureTime *gtime.Time `json:"measureTime" orm:"measureTime" description:"测量时间"`         // 测量时间
	UploadType  int         `json:"uploadType"  orm:"uploadType"  description:"1=手动, 2=设备上传"` // 1=手动, 2=设备上传
}
