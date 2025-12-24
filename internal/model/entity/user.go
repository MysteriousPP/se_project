// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// User is the golang structure for table user.
type User struct {
	UserID   int    `json:"userID"   orm:"userID"   description:"用户唯一ID"`       // 用户唯一ID
	Phone    string `json:"phone"    orm:"phone"    description:"登录账号（手机号）"`    // 登录账号（手机号）
	Password string `json:"password" orm:"password" description:"SHA256加盐加密存储"` // SHA256加盐加密存储
	Role     int    `json:"role"     orm:"role"     description:"1=银发端, 2=子女端"` // 1=银发端, 2=子女端
	Nickname string `json:"nickname" orm:"nickname" description:"用户昵称"`         // 用户昵称
}
