// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FamilyFeed is the golang structure for table family_feed.
type FamilyFeed struct {
	FeedId     int         `json:"feedId"     orm:"feedId"     description:"动态ID"`        // 动态ID
	AuthorId   int         `json:"authorId"   orm:"authorId"   description:"发布者ID"`       // 发布者ID
	Content    string      `json:"content"    orm:"content"    description:"动态内容"`        // 动态内容
	Images     string      `json:"images"     orm:"images"     description:"图片URL数组"`     // 图片URL数组
	CreateTime *gtime.Time `json:"createTime" orm:"createTime" description:"创建时间"`        // 创建时间
	IsDeleted  int         `json:"isDeleted"  orm:"isDeleted"  description:"0=正常, 1=已删除"` // 0=正常, 1=已删除
}
