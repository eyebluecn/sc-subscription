package model

import "github.com/eyebluecn/sc-misc/src/model/vo_model"

// 专栏领域模型  这个只是视图模型。
type RichSubscription struct {
	Subscription *Subscription
	Column       *vo_model.ColumnVO
	Reader       *vo_model.ReaderVO
	Order        *Order
}
