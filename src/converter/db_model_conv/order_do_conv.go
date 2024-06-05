package db_model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func OrderStatusToStorage(status model.OrderStatus) int32 {
	return int32(status)
}

// 枚举转为存储的整数
func OrderStatusesToStorage(things []model.OrderStatus) []int32 {
	if things == nil {
		return nil
	}
	var resultList []int32
	for _, item := range things {
		resultList = append(resultList, OrderStatusToStorage(item))
	}
	return resultList
}

// 数据库模型转换为领域模型
func ConvertOrderDO(thing *model.Order) *po.OrderPO {
	if thing == nil {
		return nil
	}

	return &po.OrderPO{
		ID:            thing.ID,
		CreateTime:    thing.CreateTime,
		UpdateTime:    thing.UpdateTime,
		No:            thing.No,
		ReaderID:      thing.ReaderID,
		ColumnID:      thing.ColumnID,
		ColumnQuoteID: thing.ColumnQuoteID,
		PaymentID:     thing.PaymentID,
		Price:         thing.Price,
		Status:        OrderStatusToStorage(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertOrderDOs(things []*model.Order) []*po.OrderPO {
	if things == nil {
		return nil
	}
	var results []*po.OrderPO
	for _, item := range things {
		results = append(results, ConvertOrderDO(item))
	}
	return results
}
