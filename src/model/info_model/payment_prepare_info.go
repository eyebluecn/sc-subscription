package info_model

import (
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
)

// 准备订阅的信息模型
type PaymentPrepareInfo struct {
	PaymentVO          *vo_model.PaymentVO //订单模型
	ThirdTransactionNo string
	NonceStr           string
}
