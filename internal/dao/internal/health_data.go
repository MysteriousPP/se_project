// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// HealthDataDao is the data access object for the table health_data.
type HealthDataDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  HealthDataColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// HealthDataColumns defines and stores column names for the table health_data.
type HealthDataColumns struct {
	DataID      string // 数据唯一ID
	UserID      string // 银发端用户ID
	HighBp      string // 高压值(mmHg)
	LowBp       string // 低压值(mmHg)
	MeasureTime string // 测量时间
	UploadType  string // 1=手动, 2=设备上传
}

// healthDataColumns holds the columns for the table health_data.
var healthDataColumns = HealthDataColumns{
	DataID:      "dataID",
	UserID:      "userID",
	HighBp:      "highBp",
	LowBp:       "lowBp",
	MeasureTime: "measureTime",
	UploadType:  "uploadType",
}

// NewHealthDataDao creates and returns a new DAO object for table data access.
func NewHealthDataDao(handlers ...gdb.ModelHandler) *HealthDataDao {
	return &HealthDataDao{
		group:    "default",
		table:    "health_data",
		columns:  healthDataColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *HealthDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *HealthDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *HealthDataDao) Columns() HealthDataColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *HealthDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *HealthDataDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *HealthDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
