package api_conv

import (
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
	"github.com/eyebluecn/sc-misc/src/util"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 领域模型转为传输模型
func ConvertReaderDTO(thing *vo_model.ReaderVO) *sc_subscription_api.ReaderDTO {
	if thing == nil {
		return nil
	}

	return &sc_subscription_api.ReaderDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Username:   thing.Username,
	}
}

// 列表转换
func ConvertReaderDTOs(things []*vo_model.ReaderVO) []*sc_subscription_api.ReaderDTO {
	if things == nil {
		return nil
	}
	var resultList []*sc_subscription_api.ReaderDTO
	for _, item := range things {
		resultList = append(resultList, ConvertReaderDTO(item))
	}
	return resultList
}
