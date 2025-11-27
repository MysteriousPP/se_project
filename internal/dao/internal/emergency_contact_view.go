// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EmergencyContactViewDao is the data access object for the table emergency_contact_view.
type EmergencyContactViewDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  EmergencyContactViewColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// EmergencyContactViewColumns defines and stores column names for the table emergency_contact_view.
type EmergencyContactViewColumns struct {
	ContactId    string //
	UserId       string // 所属用户ID
	Name         string // 联系人姓名
	Phone        string // 联系人电话
	Relationship string // 关系描述
	IsCommunity  string // 0=个人, 1=社区联系人
	UserNickname string // 用户昵称
	UserPhone    string // 登陆账号（手机号）
}

// emergencyContactViewColumns holds the columns for the table emergency_contact_view.
var emergencyContactViewColumns = EmergencyContactViewColumns{
	ContactId:    "contactId",
	UserId:       "userId",
	Name:         "name",
	Phone:        "phone",
	Relationship: "relationship",
	IsCommunity:  "isCommunity",
	UserNickname: "userNickname",
	UserPhone:    "userPhone",
}

// NewEmergencyContactViewDao creates and returns a new DAO object for table data access.
func NewEmergencyContactViewDao(handlers ...gdb.ModelHandler) *EmergencyContactViewDao {
	return &EmergencyContactViewDao{
		group:    "default",
		table:    "emergency_contact_view",
		columns:  emergencyContactViewColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *EmergencyContactViewDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *EmergencyContactViewDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *EmergencyContactViewDao) Columns() EmergencyContactViewColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *EmergencyContactViewDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *EmergencyContactViewDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *EmergencyContactViewDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
