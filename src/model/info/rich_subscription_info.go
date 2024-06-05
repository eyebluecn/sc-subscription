package info

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/vo"
)

// 专栏领域模型  这个只是视图模型。
type RichSubscription struct {
	Subscription *do.SubscriptionDO
	Column       *vo.ColumnVO
	Reader       *vo.ReaderVO
	Order        *do.OrderDO
}
