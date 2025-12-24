// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CommunityServiceDao is the data access object for the table community_service.
type CommunityServiceDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  CommunityServiceColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// CommunityServiceColumns defines and stores column names for the table community_service.
type CommunityServiceColumns struct {
	ServiceId     string // 服务ID
	ServiceName   string // 服务名称
	Price         string // 服务价格
	Duration      string // 服务时长(分钟)
	Provider      string // 服务提供方
	ServiceStatus string // 1=可预约, 0=已下架
}

// communityServiceColumns holds the columns for the table community_service.
var communityServiceColumns = CommunityServiceColumns{
	ServiceId:     "serviceId",
	ServiceName:   "serviceName",
	Price:         "price",
	Duration:      "duration",
	Provider:      "provider",
	ServiceStatus: "serviceStatus",
}

// NewCommunityServiceDao creates and returns a new DAO object for table data access.
func NewCommunityServiceDao(handlers ...gdb.ModelHandler) *CommunityServiceDao {
	return &CommunityServiceDao{
		group:    "default",
		table:    "community_service",
		columns:  communityServiceColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *CommunityServiceDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *CommunityServiceDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *CommunityServiceDao) Columns() CommunityServiceColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *CommunityServiceDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *CommunityServiceDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *CommunityServiceDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
