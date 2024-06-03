package repo

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"time"
)

type SubscriptionPageRequest struct {

	//创建时间晚于
	CreateTimeGte time.Time
	//状态
	Status *model.SubscriptionStatus
	//读者id
	ReaderId *int64
	//专栏id
	ColumnIds []int64
	//订单Id
	OrderId *int64

	//按照时间排序
	OrderBy SubscriptionPageOrderBy
	//当前页 1基
	PageNum int64
	//每页大小
	PageSize int64
}

// 查询排序
type SubscriptionPageOrderBy int32

const (
	OrderByCreateTimeDesc SubscriptionPageOrderBy = 0
	OrderByCreateTimeAsc  SubscriptionPageOrderBy = 1
	OrderByIdAsc          SubscriptionPageOrderBy = 2
	OrderByIdDesc         SubscriptionPageOrderBy = 3
)
