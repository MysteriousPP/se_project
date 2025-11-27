// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SosAlert is the golang structure of table sos_alert for DAO operations like Where/Data.
type SosAlert struct {
	g.Meta      `orm:"table:sos_alert, do:true"`
	AlertId     interface{} //
	ElderUserId interface{} // 触发求助的老人ID
	TriggerTime *gtime.Time // 触发时间
	Location    interface{} // 求助时的位置信息
	Longitude   interface{} // 经度
	Latitude    interface{} // 纬度
	Status      interface{} // 1=待处理, 2=已接听, 3=已处理
	HandledBy   interface{} // 处理人ID
	HandledTime *gtime.Time // 处理时间
	Notes       interface{} // 处理备注
}
