// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FamilyOverview is the golang structure for table family_overview.
type FamilyOverview struct {
	ElderId         int         `json:"elderId"         orm:"elderId"         description:"用户唯一ID"` // 用户唯一ID
	ElderName       string      `json:"elderName"       orm:"elderName"       description:"用户昵称"`   // 用户昵称
	ChildrenCount   int64       `json:"childrenCount"   orm:"childrenCount"   description:""`       //
	LastHealthCheck *gtime.Time `json:"lastHealthCheck" orm:"lastHealthCheck" description:"测量时间"`   // 测量时间
	PendingAlerts   int64       `json:"pendingAlerts"   orm:"pendingAlerts"   description:""`       //
}
