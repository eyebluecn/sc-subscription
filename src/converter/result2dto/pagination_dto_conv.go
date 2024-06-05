package result2dto

import (
	"github.com/eyebluecn/sc-misc/src/model/result"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_base"
)

// 领域模型转为传输模型
func ConvertPagination(pagination *result.Pagination) *sc_subscription_base.Pagination {
	if pagination == nil {
		return nil
	}

	return &sc_subscription_base.Pagination{
		PageNum:    pagination.PageNum,
		PageSize:   pagination.PageSize,
		TotalItems: pagination.TotalItems,
	}
}
