package do2po

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 枚举转为存储的整数
func ConvertOrderStatus(status enums.OrderStatus) int32 {
	return int32(status)
}

// 枚举转为存储的整数
func ConvertOrderStatuses(things []enums.OrderStatus) []int32 {
	if things == nil {
		return nil
	}
	var resultList []int32
	for _, item := range things {
		resultList = append(resultList, ConvertOrderStatus(item))
	}
	return resultList
}

// 数据库模型转换为领域模型
func ConvertOrderPO(thing *do.OrderDO) *po.OrderPO {
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
		Status:        ConvertOrderStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertOrderPOs(things []*do.OrderDO) []*po.OrderPO {
	if things == nil {
		return nil
	}
	var results []*po.OrderPO
	for _, item := range things {
		results = append(results, ConvertOrderPO(item))
	}
	return results
}
