// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SosAlert is the golang structure for table sos_alert.
type SosAlert struct {
	AlertId     int         `json:"alertId"     orm:"alertId"     description:"求助记录ID"`              // 求助记录ID
	ElderUserId int         `json:"elderUserId" orm:"elderUserId" description:"触发求助的老人ID"`           // 触发求助的老人ID
	TriggerTime *gtime.Time `json:"triggerTime" orm:"triggerTime" description:"触发时间"`                // 触发时间
	Location    string      `json:"location"    orm:"location"    description:"求助时的详细地址描述"`          // 求助时的详细地址描述
	Longitude   string      `json:"longitude"   orm:"longitude"   description:"经度"`                  // 经度
	Latitude    string      `json:"latitude"    orm:"latitude"    description:"纬度"`                  // 纬度
	Status      int         `json:"status"      orm:"status"      description:"1=待处理, 2=已接听, 3=已处理"` // 1=待处理, 2=已接听, 3=已处理
	HandledBy   int         `json:"handledBy"   orm:"handledBy"   description:"处理人ID"`               // 处理人ID
	HandledTime *gtime.Time `json:"handledTime" orm:"handledTime" description:"处理时间"`                // 处理时间
	Notes       string      `json:"notes"       orm:"notes"       description:"处理备注"`                // 处理备注
}
