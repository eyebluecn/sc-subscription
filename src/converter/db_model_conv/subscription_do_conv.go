package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func SubscriptionStatusToStorage(status model.SubscriptionStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDO(thing *model.Subscription) *po.SubscriptionPO {
	if thing == nil {
		return nil
	}

	return &po.SubscriptionPO{
		ID:         thing.ID,
		CreateTime: thing.CreateTime,
		UpdateTime: thing.UpdateTime,
		ReaderID:   thing.ReaderID,
		ColumnID:   thing.ColumnID,
		OrderID:    thing.OrderID,
		Status:     SubscriptionStatusToStorage(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDOs(things []*model.Subscription) []*po.SubscriptionPO {
	if things == nil {
		return nil
	}
	var results []*po.SubscriptionPO
	for _, item := range things {
		results = append(results, ConvertSubscriptionDO(item))
	}
	return results
}
