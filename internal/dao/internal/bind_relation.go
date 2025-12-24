// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BindRelationDao is the data access object for the table bind_relation.
type BindRelationDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  BindRelationColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// BindRelationColumns defines and stores column names for the table bind_relation.
type BindRelationColumns struct {
	BindID      string // 绑定记录ID
	ChildUserID string // 子女端用户ID
	ElderUserId string // 银发端用户ID
	BindTime    string // 绑定时间
	Status      string // 1=有效, 0=已解绑
}

// bindRelationColumns holds the columns for the table bind_relation.
var bindRelationColumns = BindRelationColumns{
	BindID:      "bindID",
	ChildUserID: "childUserID",
	ElderUserId: "elderUserId",
	BindTime:    "bindTime",
	Status:      "status",
}

// NewBindRelationDao creates and returns a new DAO object for table data access.
func NewBindRelationDao(handlers ...gdb.ModelHandler) *BindRelationDao {
	return &BindRelationDao{
		group:    "default",
		table:    "bind_relation",
		columns:  bindRelationColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BindRelationDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BindRelationDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BindRelationDao) Columns() BindRelationColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BindRelationDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BindRelationDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BindRelationDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
