// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Reminder is the golang structure of table reminder for DAO operations like Where/Data.
type Reminder struct {
	g.Meta      `orm:"table:reminder, do:true"`
	ReminderId  interface{} //
	CreatedBy   interface{} // 创建人ID
	ElderUserId interface{} // 提醒对象ID
	Title       interface{} // 提醒标题
	Content     interface{} // 提醒内容
	RemindTime  *gtime.Time // 提醒时间
	RepeatType  interface{} // 0=不重复, 1=每天, 2=每周, 3=每月
	Status      interface{} // 1=待处理, 2=已完成
}
