// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// HealthData is the golang structure of table health_data for DAO operations like Where/Data.
type HealthData struct {
	g.Meta      `orm:"table:health_data, do:true"`
	DataID      interface{} // 数据唯一ID
	UserID      interface{} // 银发端用户ID
	HighBp      interface{} // 高压值(mmHg)
	LowBp       interface{} // 低压值(mmHg)
	MeasureTime *gtime.Time // 测量时间
	UploadType  interface{} // 1=手动，2=设备上传
}
