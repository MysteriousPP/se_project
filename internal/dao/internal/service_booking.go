// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ServiceBookingDao is the data access object for the table service_booking.
type ServiceBookingDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  ServiceBookingColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// ServiceBookingColumns defines and stores column names for the table service_booking.
type ServiceBookingColumns struct {
	BookingId   string //
	ServiceId   string // 服务ID
	ElderUserId string // 预约老人ID
	BookedBy    string // 预约人ID
	BookingTime string // 预约服务时间
	Status      string // 1=待确认, 2=已确认, 3=已完成
}

// serviceBookingColumns holds the columns for the table service_booking.
var serviceBookingColumns = ServiceBookingColumns{
	BookingId:   "bookingId",
	ServiceId:   "serviceId",
	ElderUserId: "elderUserId",
	BookedBy:    "bookedBy",
	BookingTime: "bookingTime",
	Status:      "status",
}

// NewServiceBookingDao creates and returns a new DAO object for table data access.
func NewServiceBookingDao(handlers ...gdb.ModelHandler) *ServiceBookingDao {
	return &ServiceBookingDao{
		group:    "default",
		table:    "service_booking",
		columns:  serviceBookingColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ServiceBookingDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ServiceBookingDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ServiceBookingDao) Columns() ServiceBookingColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ServiceBookingDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ServiceBookingDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *ServiceBookingDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
