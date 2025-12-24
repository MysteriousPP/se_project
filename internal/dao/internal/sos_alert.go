// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SosAlertDao is the data access object for the table sos_alert.
type SosAlertDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SosAlertColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SosAlertColumns defines and stores column names for the table sos_alert.
type SosAlertColumns struct {
	AlertId     string // 求助记录ID
	ElderUserId string // 触发求助的老人ID
	TriggerTime string // 触发时间
	Location    string // 求助时的详细地址描述
	Longitude   string // 经度
	Latitude    string // 纬度
	Status      string // 1=待处理, 2=已接听, 3=已处理
	HandledBy   string // 处理人ID
	HandledTime string // 处理时间
	Notes       string // 处理备注
}

// sosAlertColumns holds the columns for the table sos_alert.
var sosAlertColumns = SosAlertColumns{
	AlertId:     "alertId",
	ElderUserId: "elderUserId",
	TriggerTime: "triggerTime",
	Location:    "location",
	Longitude:   "longitude",
	Latitude:    "latitude",
	Status:      "status",
	HandledBy:   "handledBy",
	HandledTime: "handledTime",
	Notes:       "notes",
}

// NewSosAlertDao creates and returns a new DAO object for table data access.
func NewSosAlertDao(handlers ...gdb.ModelHandler) *SosAlertDao {
	return &SosAlertDao{
		group:    "default",
		table:    "sos_alert",
		columns:  sosAlertColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SosAlertDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SosAlertDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SosAlertDao) Columns() SosAlertColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SosAlertDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SosAlertDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SosAlertDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
