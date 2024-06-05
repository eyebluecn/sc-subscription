package query

import (
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	enums2 "github.com/eyebluecn/sc-misc/src/model/query/enums"
	"time"
)

type SubscriptionPageQuery struct {

	//创建时间晚于
	CreateTimeGte time.Time
	//状态
	Status *enums.SubscriptionStatus
	//读者id
	ReaderId *int64
	//专栏id
	ColumnIds []int64
	//订单Id
	OrderId *int64

	//按照时间排序
	OrderBy enums2.SubscriptionPageOrderBy
	//当前页 1基
	PageNum int64
	//每页大小
	PageSize int64
}
