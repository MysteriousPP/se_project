// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// ServiceBooking is the golang structure for table service_booking.
type ServiceBooking struct {
	BookingId   int         `json:"bookingId"   orm:"bookingId"   description:"预约ID"`                // 预约ID
	ServiceId   int         `json:"serviceId"   orm:"serviceId"   description:"服务ID"`                // 服务ID
	ElderUserId int         `json:"elderUserId" orm:"elderUserId" description:"预约老人ID"`              // 预约老人ID
	BookedBy    int         `json:"bookedBy"    orm:"bookedBy"    description:"预约人ID"`               // 预约人ID
	BookingTime *gtime.Time `json:"bookingTime" orm:"bookingTime" description:"预约服务时间"`              // 预约服务时间
	Status      int         `json:"status"      orm:"status"      description:"1=待确认, 2=已确认, 3=已完成"` // 1=待确认, 2=已确认, 3=已完成
}
