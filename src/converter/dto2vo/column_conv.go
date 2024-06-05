package dto2vo

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/model/vo/enums"
)

// 转为枚举
func ConvertColumnStatus(status sc_misc_api.ColumnStatus) enums.ColumnStatus {
	return enums.ColumnStatus(status)
}

// 转为VO
func ConvertColumnVO(thing *sc_misc_api.ColumnDTO) *vo.ColumnVO {
	if thing == nil {
		return nil
	}
	return &vo.ColumnVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		Name:       thing.Name,
		AuthorID:   thing.AuthorId,
		Status:     ConvertColumnStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertColumnVOs(subscriptionDTOS []*sc_misc_api.ColumnDTO) []*vo.ColumnVO {
	if subscriptionDTOS == nil {
		return nil
	}
	var subscriptionVOS []*vo.ColumnVO
	for _, item := range subscriptionDTOS {
		subscriptionVOS = append(subscriptionVOS, ConvertColumnVO(item))
	}
	return subscriptionVOS
}
