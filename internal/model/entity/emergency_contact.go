// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// EmergencyContact is the golang structure for table emergency_contact.
type EmergencyContact struct {
	ContactId    int    `json:"contactId"    orm:"contactId"    description:"联系人ID"`         // 联系人ID
	UserId       int    `json:"userId"       orm:"userId"       description:"所属用户ID"`        // 所属用户ID
	Name         string `json:"name"         orm:"name"         description:"联系人姓名"`         // 联系人姓名
	Phone        string `json:"phone"        orm:"phone"        description:"联系人电话"`         // 联系人电话
	Relationship string `json:"relationship" orm:"relationship" description:"关系描述"`          // 关系描述
	IsCommunity  int    `json:"isCommunity"  orm:"isCommunity"  description:"0=个人, 1=社区联系人"` // 0=个人, 1=社区联系人
}
