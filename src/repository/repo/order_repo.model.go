package repo

import "github.com/eyebluecn/sc-misc/src/model"

type OrderPageRequest struct {

	//读者id
	ReaderId *int64
	//专栏id
	ColumnId *int64
	//状态
	Statuses []model.OrderStatus

	//当前页 1基
	PageNum int64
	//每页大小
	PageSize int64
}
