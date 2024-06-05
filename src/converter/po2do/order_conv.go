package po2do

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 转为枚举
func ConvertOrderStatus(status int32) enums.OrderStatus {
	return enums.OrderStatus(status)
}

// 数据库模型转换为领域模型
func ConvertOrderDO(thing *po.OrderPO) *do.OrderDO {
	if thing == nil {
		return nil
	}

	return &do.OrderDO{
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
func ConvertOrderDOs(things []*po.OrderPO) []*do.OrderDO {
	if things == nil {
		return nil
	}
	var orders []*do.OrderDO
	for _, item := range things {
		orders = append(orders, ConvertOrderDO(item))
	}
	return orders
}
