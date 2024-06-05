package dto2vo

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
)

// 转为VO
func ConvertReaderVO(thing *sc_misc_api.ReaderDTO) *vo.ReaderVO {
	if thing == nil {
		return nil
	}
	return &vo.ReaderVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Username:   thing.Username,
		Password:   "",
	}
}

// 转为VO数组
func ConvertReaderVOs(subscriptionDTOS []*sc_misc_api.ReaderDTO) []*vo.ReaderVO {
	if subscriptionDTOS == nil {
		return nil
	}
	var subscriptionVOS []*vo.ReaderVO
	for _, item := range subscriptionDTOS {
		subscriptionVOS = append(subscriptionVOS, ConvertReaderVO(item))
	}
	return subscriptionVOS
}
