// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BindRelation is the golang structure of table bind_relation for DAO operations like Where/Data.
type BindRelation struct {
	g.Meta      `orm:"table:bind_relation, do:true"`
	BindID      interface{} // 绑定记录ID
	ChildUserID interface{} // 子女端用户ID
	ElderUserId interface{} // 银发端用户ID
	BindTime    *gtime.Time // 绑定时间
	Status      interface{} // 1=有效，0=已解绑
}
