package info_model

import "github.com/eyebluecn/sc-misc/src/model"

// 准备订阅的信息模型
type PrepareSubscribeInfo struct {
	Order              *model.Order //订单模型
	ThirdTransactionNo string
	NonceStr           string
}
