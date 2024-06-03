package vo_model

// 专栏状态
type ColumnStatus int32

const (
	ColumnStatusNew      ColumnStatus = 0 //未发布
	ColumnStatusOk       ColumnStatus = 1 //已生效
	ColumnStatusDisabled ColumnStatus = 2 //已禁用
)

// 支付单状态
type PaymentStatus int32

const (
	PaymentStatusUnpaid PaymentStatus = 0 //未支付
	PaymentStatusPaid   PaymentStatus = 1 //已支付
	PaymentStatusClosed PaymentStatus = 2 //已关闭
)

// 专栏定价状态
type ColumnQuoteStatus int32

const (
	ColumnQuoteStatusNew ColumnQuoteStatus = 0 //未生效
	ColumnQuoteStatusOk  ColumnQuoteStatus = 1 //已生效
)
