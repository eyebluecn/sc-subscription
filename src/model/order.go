package model

import (
	"time"
)

// 订单 领域模型
type Order struct {
	ID            int64       // 主键
	CreateTime    time.Time   // 创建时间
	UpdateTime    time.Time   // 更新时间
	No            string      // 订单唯一编号，整个系统唯一，带有前缀
	ReaderID      int64       // 读者id
	ColumnID      int64       // 专栏id
	ColumnQuoteID int64       // 专栏报价id
	PaymentID     int64       // 支付单id
	Price         int64       // 价格（单位：分）
	Status        OrderStatus // 状态 0/1/2/3/4
}
