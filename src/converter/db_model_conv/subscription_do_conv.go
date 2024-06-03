package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/db_model"
)

// 枚举转为存储的整数
func SubscriptionStatusToStorage(status model.SubscriptionStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertSubscriptionDO(thing *model.Subscription) *db_model.SubscriptionDO {
	if thing == nil {
		return nil
	}

	return &db_model.SubscriptionDO{
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
func ConvertSubscriptionDOs(things []*model.Subscription) []*db_model.SubscriptionDO {
	if things == nil {
		return nil
	}
	var results []*db_model.SubscriptionDO
	for _, item := range things {
		results = append(results, ConvertSubscriptionDO(item))
	}
	return results
}
