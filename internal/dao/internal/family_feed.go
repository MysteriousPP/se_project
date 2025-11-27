// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FamilyFeedDao is the data access object for the table family_feed.
type FamilyFeedDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  FamilyFeedColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// FamilyFeedColumns defines and stores column names for the table family_feed.
type FamilyFeedColumns struct {
	FeedId     string //
	AuthorId   string // 发布者ID
	Content    string // 动态内容
	Images     string // 图片URL数组
	CreateTime string //
	IsDeleted  string // 0=正常, 1=已删除
}

// familyFeedColumns holds the columns for the table family_feed.
var familyFeedColumns = FamilyFeedColumns{
	FeedId:     "feedId",
	AuthorId:   "authorId",
	Content:    "content",
	Images:     "images",
	CreateTime: "createTime",
	IsDeleted:  "isDeleted",
}

// NewFamilyFeedDao creates and returns a new DAO object for table data access.
func NewFamilyFeedDao(handlers ...gdb.ModelHandler) *FamilyFeedDao {
	return &FamilyFeedDao{
		group:    "default",
		table:    "family_feed",
		columns:  familyFeedColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *FamilyFeedDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *FamilyFeedDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *FamilyFeedDao) Columns() FamilyFeedColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *FamilyFeedDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *FamilyFeedDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *FamilyFeedDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
