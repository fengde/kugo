package dao

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ WxWjsqzUserModel = (*customWxWjsqzUserModel)(nil)

type (
	// WxWjsqzUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWxWjsqzUserModel.
	WxWjsqzUserModel interface {
		wxWjsqzUserModel
	}

	customWxWjsqzUserModel struct {
		*defaultWxWjsqzUserModel
	}
)

// NewWxWjsqzUserModel returns a model for the database table.
func NewWxWjsqzUserModel(conn sqlx.SqlConn) WxWjsqzUserModel {
	return &customWxWjsqzUserModel{
		defaultWxWjsqzUserModel: newWxWjsqzUserModel(conn),
	}
}
