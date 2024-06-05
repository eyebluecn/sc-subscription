package dto2vo

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/model/vo/enums"
)

// 转为枚举
func ConvertPaymentStatus(status sc_misc_api.PaymentStatus) enums.PaymentStatus {
	return enums.PaymentStatus(status)
}

// 转为VO
func ConvertPaymentVO(thing *sc_misc_api.PaymentDTO) *vo.PaymentVO {
	if thing == nil {
		return nil
	}
	return &vo.PaymentVO{
		ID:                 thing.Id,
		CreateTime:         util.ParseTimestamp(thing.CreateTime),
		UpdateTime:         util.ParseTimestamp(thing.UpdateTime),
		OrderNo:            thing.OrderNo,
		Method:             thing.Method,
		ThirdTransactionNo: thing.ThirdTransactionNo,
		Amount:             thing.Amount,
		Status:             ConvertPaymentStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertPaymentVOs(things []*sc_misc_api.PaymentDTO) []*vo.PaymentVO {
	if things == nil {
		return nil
	}
	var resultList []*vo.PaymentVO
	for _, item := range things {
		resultList = append(resultList, ConvertPaymentVO(item))
	}
	return resultList
}
