package vo_conv

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
	"github.com/eyebluecn/sc-misc/src/util"
)

// 转为枚举
func ConvertColumnQuoteStatus(status sc_misc_api.ColumnQuoteStatus) vo_model.ColumnQuoteStatus {
	return vo_model.ColumnQuoteStatus(status)
}

// 转为VO
func ConvertColumnQuoteVO(thing *sc_misc_api.ColumnQuoteDTO) *vo_model.ColumnQuoteVO {
	if thing == nil {
		return nil
	}
	return &vo_model.ColumnQuoteVO{
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
func ConvertColumnQuotes(things []*sc_misc_api.ColumnQuoteDTO) []*vo_model.ColumnQuoteVO {
	if things == nil {
		return nil
	}
	var resultList []*vo_model.ColumnQuoteVO
	for _, item := range things {
		resultList = append(resultList, ConvertColumnQuoteVO(item))
	}
	return resultList
}
