package info

import (
	"github.com/eyebluecn/sc-misc/src/model/do"
)

// 准备订阅的信息模型
type PrepareSubscribeInfo struct {
	Order              *do.OrderDO //订单模型
	ThirdTransactionNo string
	NonceStr           string
}
