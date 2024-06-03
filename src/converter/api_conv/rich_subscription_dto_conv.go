package api_conv

import (
	api_conv "github.com/eyebluecn/sc-misc/src/converter/vo2api_conv"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 领域模型转为传输模型
func ConvertRichSubscriptionDTO(thing *model.RichSubscription) *sc_subscription_api.RichSubscriptionDTO {
	if thing == nil {
		return nil
	}

	return &sc_subscription_api.RichSubscriptionDTO{
		Subscription: ConvertSubscriptionDTO(thing.Subscription),
		Column:       api_conv.ConvertColumnDTO(thing.Column),
		Reader:       api_conv.ConvertReaderDTO(thing.Reader),
		Order:        ConvertOrderDTO(thing.Order),
	}
}

// 数据库模型转换为领域模型
func ConvertRichSubscriptionDTOs(things []*model.RichSubscription) []*sc_subscription_api.RichSubscriptionDTO {
	if things == nil {
		return nil
	}
	var resultList []*sc_subscription_api.RichSubscriptionDTO
	for _, item := range things {
		resultList = append(resultList, ConvertRichSubscriptionDTO(item))
	}
	return resultList
}
