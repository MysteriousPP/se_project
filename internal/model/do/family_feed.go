// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FamilyFeed is the golang structure of table family_feed for DAO operations like Where/Data.
type FamilyFeed struct {
	g.Meta     `orm:"table:family_feed, do:true"`
	FeedId     interface{} //
	AuthorId   interface{} // 发布者ID
	Content    interface{} // 动态内容
	Images     interface{} // 图片URL数组
	CreateTime *gtime.Time //
	IsDeleted  interface{} // 0=正常, 1=已删除
}
