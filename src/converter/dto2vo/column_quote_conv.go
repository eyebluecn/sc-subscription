package dto2vo

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/model/vo/enums"
)

// 转为枚举
func ConvertColumnQuoteStatus(status sc_misc_api.ColumnQuoteStatus) enums.ColumnQuoteStatus {
	return enums.ColumnQuoteStatus(status)
}

// 转为VO
func ConvertColumnQuoteVO(thing *sc_misc_api.ColumnQuoteDTO) *vo.ColumnQuoteVO {
	if thing == nil {
		return nil
	}
	return &vo.ColumnQuoteVO{
		ID:         thing.Id,
		CreateTime: util.ParseTimestamp(thing.CreateTime),
		UpdateTime: util.ParseTimestamp(thing.UpdateTime),
		ColumnID:   thing.ColumnId,
		EditorID:   thing.EditorId,
		Price:      thing.Price,
		Status:     ConvertColumnQuoteStatus(thing.Status),
	}
}

// 转为VO数组
func ConvertColumnQuoteVOs(things []*sc_misc_api.ColumnQuoteDTO) []*vo.ColumnQuoteVO {
	if things == nil {
		return nil
	}
	var resultList []*vo.ColumnQuoteVO
	for _, item := range things {
		resultList = append(resultList, ConvertColumnQuoteVO(item))
	}
	return resultList
}
