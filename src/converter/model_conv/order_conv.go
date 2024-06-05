package model_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/model/po"
)

// 转为枚举
func ConvertOrderStatus(status int32) model.OrderStatus {
	return model.OrderStatus(status)
}

// 数据库模型转换为领域模型
func ConvertOrder(thing *po.OrderPO) *model.Order {
	if thing == nil {
		return nil
	}

	return &model.Order{
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
func ConvertOrders(things []*po.OrderPO) []*model.Order {
	if things == nil {
		return nil
	}
	var orders []*model.Order
	for _, item := range things {
		orders = append(orders, ConvertOrder(item))
	}
	return orders
}
