package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status int32) model.SubscriptionStatus {
	return model.SubscriptionStatus(status)
}

// 转为枚举
func ConvertSubscriptionStatusPtr(status *sc_subscription_api.SubscriptionStatus) *model.SubscriptionStatus {
	if status == nil {
		return nil
	}
	subscriptionStatus := model.SubscriptionStatus(*status)
	return &subscriptionStatus
}

// 数据库模型转换为领域模型
func ConvertSubscription(thing *db_model.SubscriptionDO) *model.Subscription {
	if thing == nil {
		return nil
	}

	return &model.Subscription{
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
func ConvertSubscriptions(things []*db_model.SubscriptionDO) []*model.Subscription {
	if things == nil {
		return nil
	}
	var subscriptions []*model.Subscription
	for _, item := range things {
		subscriptions = append(subscriptions, ConvertSubscription(item))
	}
	return subscriptions
}
