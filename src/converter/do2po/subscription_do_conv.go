package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ConvertSubscriptionStatus(status enums.SubscriptionStatus) int32 {
	return int32(status)
}

// 数据库模型转换为领域模型
func ConvertSubscriptionPO(thing *do.SubscriptionDO) *po.SubscriptionPO {
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
		Status:     ConvertSubscriptionStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertSubscriptionPOs(things []*do.SubscriptionDO) []*po.SubscriptionPO {
	if things == nil {
		return nil
	}
	var results []*po.SubscriptionPO
	for _, item := range things {
		results = append(results, ConvertSubscriptionPO(item))
	}
	return results
}
