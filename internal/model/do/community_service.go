// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityService is the golang structure of table community_service for DAO operations like Where/Data.
type CommunityService struct {
	g.Meta        `orm:"table:community_service, do:true"`
	ServiceId     interface{} //
	ServiceName   interface{} // 服务名称
	Price         interface{} // 服务价格
	Duration      interface{} // 服务时长(分钟)
	Provider      interface{} // 服务提供方
	ServiceStatus interface{} // 1=可预约, 0=已下架
}
