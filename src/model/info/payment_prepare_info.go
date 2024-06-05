package info

import (
	"github.com/eyebluecn/sc-misc/src/model/vo"
)

// 准备订阅的信息模型
type PaymentPrepareInfo struct {
	PaymentVO          *vo.PaymentVO //订单模型
	ThirdTransactionNo string
	NonceStr           string
}
