package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
	"github.com/eyebluecn/sc-misc/src/converter/po2do"
	"github.com/eyebluecn/sc-misc/src/converter/universal2dto"
	"github.com/eyebluecn/sc-misc/src/model/query"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type SubscriptionPage struct{}

func NewSubscriptionPage() *SubscriptionPage {
	return &SubscriptionPage{}
}

// 服务请求入口
func (receiver SubscriptionPage) Handle(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest) (r *sc_subscription_api.SubscriptionPageResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver SubscriptionPage) CheckParam(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver SubscriptionPage) doHandle(ctx context.Context, request sc_subscription_api.SubscriptionPageRequest) (r *sc_subscription_api.SubscriptionPageResponse, err error) {
	req := query.SubscriptionPageQuery{
		Status:    po2do.ConvertSubscriptionStatusPtr(request.Status),
		ReaderId:  request.ReaderId,
		ColumnIds: request.ColumnIds,
		OrderId:   request.OrderId,
		PageNum:   request.PageNum,
		PageSize:  request.PageSize,
	}
	list, pagination, err := repo.NewSubscriptionRepo().Page(ctx, req)
	if err != nil {
		return nil, err
	}

	r = &sc_subscription_api.SubscriptionPageResponse{
		Data:       do2dto.ConvertSubscriptionDTOs(list),
		Pagination: universal2dto.ConvertPagination(pagination),
	}

	return r, nil
}
