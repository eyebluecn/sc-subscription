package po2do

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status int32) enums.SubscriptionStatus {
	return enums.SubscriptionStatus(status)
}

// 转为枚举
func ConvertSubscriptionStatusPtr(status *sc_subscription_api.SubscriptionStatus) *enums.SubscriptionStatus {
	if status == nil {
		return nil
	}
	subscriptionStatus := enums.SubscriptionStatus(*status)
	return &subscriptionStatus
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDO(thing *po.SubscriptionPO) *do.SubscriptionDO {
	if thing == nil {
		return nil
	}

	return &do.SubscriptionDO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		ReaderID:   thing.ReaderID,
		ColumnID:   thing.ColumnID,
		OrderID:    thing.OrderID,
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDOs(things []*po.SubscriptionPO) []*do.SubscriptionDO {
	if things == nil {
		return nil
	}
	var subscriptions []*do.SubscriptionDO
	for _, item := range things {
		subscriptions = append(subscriptions, ConvertSubscriptionDO(item))
	}
	return subscriptions
}
