package api_conv

import (
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/model/vo/enums"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 转为枚举
func ConvertColumnStatus(status enums.ColumnStatus) sc_subscription_api.ColumnStatus {
	return sc_subscription_api.ColumnStatus(status)
}

// 领域模型转为传输模型
func ConvertColumnDTO(thing *vo.ColumnVO) *sc_subscription_api.ColumnDTO {
	if thing == nil {
		return nil
	}

	return &sc_subscription_api.ColumnDTO{
		Id:         thing.ID,
		CreateTime: util.Timestamp(thing.CreateTime),
		UpdateTime: util.Timestamp(thing.UpdateTime),
		Name:       thing.Name,
		AuthorId:   thing.AuthorID,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 数据库模型转换为领域模型
func ConvertColumnDTOs(things []*vo.ColumnVO) []*sc_subscription_api.ColumnDTO {
	if things == nil {
		return nil
	}
	var resultList []*sc_subscription_api.ColumnDTO
	for _, item := range things {
		resultList = append(resultList, ConvertColumnDTO(item))
	}
	return resultList
}
