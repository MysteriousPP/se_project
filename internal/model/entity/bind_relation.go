// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BindRelation is the golang structure for table bind_relation.
type BindRelation struct {
	BindID      int         `json:"bindID"      orm:"bindID"      description:"绑定记录ID"`      // 绑定记录ID
	ChildUserID int         `json:"childUserID" orm:"childUserID" description:"子女端用户ID"`     // 子女端用户ID
	ElderUserId int         `json:"elderUserId" orm:"elderUserId" description:"银发端用户ID"`     // 银发端用户ID
	BindTime    *gtime.Time `json:"bindTime"    orm:"bindTime"    description:"绑定时间"`        // 绑定时间
	Status      int         `json:"status"      orm:"status"      description:"1=有效, 0=已解绑"` // 1=有效, 0=已解绑
}
