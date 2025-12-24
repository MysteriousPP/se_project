// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmergencyContactDao is the data access object for the table emergency_contact.
type EmergencyContactDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  EmergencyContactColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// EmergencyContactColumns defines and stores column names for the table emergency_contact.
type EmergencyContactColumns struct {
	ContactId    string // 联系人ID
	UserId       string // 所属用户ID
	Name         string // 联系人姓名
	Phone        string // 联系人电话
	Relationship string // 关系描述
	IsCommunity  string // 0=个人, 1=社区联系人
}

// emergencyContactColumns holds the columns for the table emergency_contact.
var emergencyContactColumns = EmergencyContactColumns{
	ContactId:    "contactId",
	UserId:       "userId",
	Name:         "name",
	Phone:        "phone",
	Relationship: "relationship",
	IsCommunity:  "isCommunity",
}

// NewEmergencyContactDao creates and returns a new DAO object for table data access.
func NewEmergencyContactDao(handlers ...gdb.ModelHandler) *EmergencyContactDao {
	return &EmergencyContactDao{
		group:    "default",
		table:    "emergency_contact",
		columns:  emergencyContactColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmergencyContactDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmergencyContactDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmergencyContactDao) Columns() EmergencyContactColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmergencyContactDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmergencyContactDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *EmergencyContactDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
