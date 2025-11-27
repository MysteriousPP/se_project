// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FamilyOverview is the golang structure of table family_overview for DAO operations like Where/Data.
type FamilyOverview struct {
	g.Meta          `orm:"table:family_overview, do:true"`
	ElderId         interface{} // 用户唯一ID
	ElderName       interface{} // 用户昵称
	ChildrenCount   interface{} //
	LastHealthCheck *gtime.Time // 测量时间
	PendingAlerts   interface{} //
}
