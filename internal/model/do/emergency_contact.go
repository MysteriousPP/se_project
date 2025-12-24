// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// EmergencyContact is the golang structure of table emergency_contact for DAO operations like Where/Data.
type EmergencyContact struct {
	g.Meta       `orm:"table:emergency_contact, do:true"`
	ContactId    interface{} // 联系人ID
	UserId       interface{} // 所属用户ID
	Name         interface{} // 联系人姓名
	Phone        interface{} // 联系人电话
	Relationship interface{} // 关系描述
	IsCommunity  interface{} // 0=个人, 1=社区联系人
}
