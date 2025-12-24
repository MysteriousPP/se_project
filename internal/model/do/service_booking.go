// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// ServiceBooking is the golang structure of table service_booking for DAO operations like Where/Data.
type ServiceBooking struct {
	g.Meta      `orm:"table:service_booking, do:true"`
	BookingId   interface{} // 预约ID
	ServiceId   interface{} // 服务ID
	ElderUserId interface{} // 预约老人ID
	BookedBy    interface{} // 预约人ID
	BookingTime *gtime.Time // 预约服务时间
	Status      interface{} // 1=待确认, 2=已确认, 3=已完成
}
