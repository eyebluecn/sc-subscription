package do2dto

import (
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertSubscriptionStatus(status enums.SubscriptionStatus) sc_subscription_api.SubscriptionStatus {
	return sc_subscription_api.SubscriptionStatus(status)
}

// 领域模型转为传输模型
func ConvertSubscriptionDTO(thing *do.SubscriptionDO) *sc_subscription_api.SubscriptionDTO {
	if thing == nil {
		return nil
	}

	return &sc_subscription_api.SubscriptionDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		ReaderId:   thing.ReaderID,
		ColumnId:   thing.ColumnID,
		OrderId:    thing.OrderID,
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDTOs(things []*do.SubscriptionDO) []*sc_subscription_api.SubscriptionDTO {
	if things == nil {
		return nil
	}
	var subscriptionDTOs []*sc_subscription_api.SubscriptionDTO
	for _, item := range things {
		subscriptionDTOs = append(subscriptionDTOs, ConvertSubscriptionDTO(item))
	}
	return subscriptionDTOs
}
