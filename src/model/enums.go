package model

// 订阅状态
type SubscriptionStatus int32

const (
	SubscriptionStatusCreated  SubscriptionStatus = 0 //已创建
	SubscriptionStatusOk       SubscriptionStatus = 1 //已生效
	SubscriptionStatusDisabled SubscriptionStatus = 2 //已失效
)

// 订单状态
type OrderStatus int32

const (
	OrderStatusCreated    OrderStatus = 0 //已创建
	OrderStatusPaid       OrderStatus = 1 //已支付
	OrderStatusSubscribed OrderStatus = 2 //已订阅
	OrderStatusClosed     OrderStatus = 3 //已关闭
	OrderStatusCanceled   OrderStatus = 4 //已取消
)
