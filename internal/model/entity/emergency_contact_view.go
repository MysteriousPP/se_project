// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// EmergencyContactView is the golang structure for table emergency_contact_view.
type EmergencyContactView struct {
	ContactId    int    `json:"contactId"    orm:"contactId"    description:""`              //
	UserId       int    `json:"userId"       orm:"userId"       description:"所属用户ID"`        // 所属用户ID
	Name         string `json:"name"         orm:"name"         description:"联系人姓名"`         // 联系人姓名
	Phone        string `json:"phone"        orm:"phone"        description:"联系人电话"`         // 联系人电话
	Relationship string `json:"relationship" orm:"relationship" description:"关系描述"`          // 关系描述
	IsCommunity  int    `json:"isCommunity"  orm:"isCommunity"  description:"0=个人, 1=社区联系人"` // 0=个人, 1=社区联系人
	UserNickname string `json:"userNickname" orm:"userNickname" description:"用户昵称"`          // 用户昵称
	UserPhone    string `json:"userPhone"    orm:"userPhone"    description:"登陆账号（手机号）"`     // 登陆账号（手机号）
}
