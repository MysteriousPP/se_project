// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Reminder is the golang structure for table reminder.
type Reminder struct {
	ReminderId  int         `json:"reminderId"  orm:"reminderId"  description:"提醒ID"`                    // 提醒ID
	CreatedBy   int         `json:"createdBy"   orm:"createdBy"   description:"创建人ID"`                   // 创建人ID
	ElderUserId int         `json:"elderUserId" orm:"elderUserId" description:"提醒对象ID"`                  // 提醒对象ID
	Title       string      `json:"title"       orm:"title"       description:"提醒标题"`                    // 提醒标题
	Content     string      `json:"content"     orm:"content"     description:"提醒内容"`                    // 提醒内容
	RemindTime  *gtime.Time `json:"remindTime"  orm:"remindTime"  description:"提醒时间"`                    // 提醒时间
	RepeatType  int         `json:"repeatType"  orm:"repeatType"  description:"0=不重复, 1=每天, 2=每周, 3=每月"` // 0=不重复, 1=每天, 2=每周, 3=每月
	Status      int         `json:"status"      orm:"status"      description:"1=待处理, 2=已完成"`            // 1=待处理, 2=已完成
}
