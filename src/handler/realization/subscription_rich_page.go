package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/api_conv"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type SubscriptionRichPage struct{}

func NewSubscriptionRichPage() *SubscriptionRichPage {
	return &SubscriptionRichPage{}
}

// 服务请求入口
func (receiver SubscriptionRichPage) Handle(ctx context.Context, request *sc_subscription_api.SubscriptionRichPageRequest) (r *sc_subscription_api.SubscriptionRichPageResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver SubscriptionRichPage) CheckParam(ctx context.Context, request *sc_subscription_api.SubscriptionRichPageRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver SubscriptionRichPage) doHandle(ctx context.Context, request sc_subscription_api.SubscriptionRichPageRequest) (r *sc_subscription_api.SubscriptionRichPageResponse, err error) {

	repoRequest := repo.SubscriptionPageRequest{
		ReaderId: request.ReaderId,
		PageNum:  request.PageNum,
		PageSize: request.PageSize,
	}
	richSubscriptions, pagination, err := application.NewSubscriptionReadApp().RichSubscriptionPage(ctx, repoRequest)
	if err != nil {
		return nil, err
	}

	r = &sc_subscription_api.SubscriptionRichPageResponse{
		Data:       api_conv.ConvertRichSubscriptionDTOs(richSubscriptions),
		Pagination: api_conv.ConvertPagination(pagination),
	}

	return r, nil
}
