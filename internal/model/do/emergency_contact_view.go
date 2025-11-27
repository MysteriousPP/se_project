// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// EmergencyContactView is the golang structure of table emergency_contact_view for DAO operations like Where/Data.
type EmergencyContactView struct {
	g.Meta       `orm:"table:emergency_contact_view, do:true"`
	ContactId    interface{} //
	UserId       interface{} // 所属用户ID
	Name         interface{} // 联系人姓名
	Phone        interface{} // 联系人电话
	Relationship interface{} // 关系描述
	IsCommunity  interface{} // 0=个人, 1=社区联系人
	UserNickname interface{} // 用户昵称
	UserPhone    interface{} // 登陆账号（手机号）
}
