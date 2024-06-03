package api_conv

import (
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertOrderStatus(status model.OrderStatus) sc_subscription_api.OrderStatus {
	return sc_subscription_api.OrderStatus(status)
}

// 领域模型转为传输模型
func ConvertOrderDTO(thing *model.Order) *sc_subscription_api.OrderDTO {
	if thing == nil {
		return nil
	}

	return &sc_subscription_api.OrderDTO{
		Id:            thing.ID,
		CreateTime:    util.Timestamp(thing.CreateTime),
		UpdateTime:    util.Timestamp(thing.UpdateTime),
		No:            thing.No,
		ReaderId:      thing.ReaderID,
		ColumnId:      thing.ColumnID,
		ColumnQuoteId: thing.ColumnQuoteID,
		PaymentId:     thing.PaymentID,
		Price:         thing.Price,
		Status:        ConvertOrderStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertOrderDTOs(things []*model.Order) []*sc_subscription_api.OrderDTO {
	if things == nil {
		return nil
	}
	var subscriptionDTOs []*sc_subscription_api.OrderDTO
	for _, item := range things {
		subscriptionDTOs = append(subscriptionDTOs, ConvertOrderDTO(item))
	}
	return subscriptionDTOs
}
