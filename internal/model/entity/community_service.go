// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// CommunityService is the golang structure for table community_service.
type CommunityService struct {
	ServiceId     int     `json:"serviceId"     orm:"serviceId"     description:"服务ID"`         // 服务ID
	ServiceName   string  `json:"serviceName"   orm:"serviceName"   description:"服务名称"`         // 服务名称
	Price         float64 `json:"price"         orm:"price"         description:"服务价格"`         // 服务价格
	Duration      int     `json:"duration"      orm:"duration"      description:"服务时长(分钟)"`     // 服务时长(分钟)
	Provider      string  `json:"provider"      orm:"provider"      description:"服务提供方"`        // 服务提供方
	ServiceStatus int     `json:"serviceStatus" orm:"serviceStatus" description:"1=可预约, 0=已下架"` // 1=可预约, 0=已下架
}
