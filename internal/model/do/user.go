// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// User is the golang structure of table user for DAO operations like Where/Data.
type User struct {
	g.Meta   `orm:"table:user, do:true"`
	UserID   interface{} // 用户唯一ID
	Phone    interface{} // 登录账号（手机号）
	Password interface{} // SHA256加盐加密存储
	Role     interface{} // 1=银发端, 2=子女端
	Nickname interface{} // 用户昵称
}
