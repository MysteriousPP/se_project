// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FamilyOverviewDao is the data access object for the table family_overview.
type FamilyOverviewDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  FamilyOverviewColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// FamilyOverviewColumns defines and stores column names for the table family_overview.
type FamilyOverviewColumns struct {
	ElderId         string // 用户唯一ID
	ElderName       string // 用户昵称
	ChildrenCount   string //
	LastHealthCheck string // 测量时间
	PendingAlerts   string //
}

// familyOverviewColumns holds the columns for the table family_overview.
var familyOverviewColumns = FamilyOverviewColumns{
	ElderId:         "elderId",
	ElderName:       "elderName",
	ChildrenCount:   "childrenCount",
	LastHealthCheck: "lastHealthCheck",
	PendingAlerts:   "pendingAlerts",
}

// NewFamilyOverviewDao creates and returns a new DAO object for table data access.
func NewFamilyOverviewDao(handlers ...gdb.ModelHandler) *FamilyOverviewDao {
	return &FamilyOverviewDao{
		group:    "default",
		table:    "family_overview",
		columns:  familyOverviewColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FamilyOverviewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FamilyOverviewDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FamilyOverviewDao) Columns() FamilyOverviewColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FamilyOverviewDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FamilyOverviewDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FamilyOverviewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
