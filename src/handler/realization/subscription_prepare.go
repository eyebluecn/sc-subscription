package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/application"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2dto"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

type SubscriptionPrepare struct{}

func NewSubscriptionPrepare() *SubscriptionPrepare {
	return &SubscriptionPrepare{}
}

// 服务请求入口
func (receiver SubscriptionPrepare) Handle(ctx context.Context, request *sc_subscription_api.SubscriptionPrepareRequest) (r *sc_subscription_api.SubscriptionPrepareResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver SubscriptionPrepare) CheckParam(ctx context.Context, request *sc_subscription_api.SubscriptionPrepareRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.ColumnId == 0 {
		return errs.BadRequestErrorf("request.ColumnId 不能为空")
	}
	if request.PayMethod == "" {
		return errs.BadRequestErrorf("request.PayMethod 不能为空")
	}
	if request.ReaderId == 0 {
		return errs.BadRequestErrorf("request.ReaderId 不能为空")
	}
	return nil
}

// 参数校验后的真实处理
func (receiver SubscriptionPrepare) doHandle(ctx context.Context, request sc_subscription_api.SubscriptionPrepareRequest) (r *sc_subscription_api.SubscriptionPrepareResponse, err error) {

	prepareSubscribeInfo, err := application.NewSubscriptionWriteAppSvc().PrepareSubscribe(ctx, request.ReaderId, request.ColumnId, request.PayMethod)
	if err != nil {
		return nil, err
	}

	r = &sc_subscription_api.SubscriptionPrepareResponse{
		Data: &sc_subscription_api.SubscriptionPrepareData{
			OrderDTO:           do2dto.ConvertOrderDTO(prepareSubscribeInfo.Order),
			ThirdTransactionNo: prepareSubscribeInfo.ThirdTransactionNo,
			NonceStr:           prepareSubscribeInfo.NonceStr,
		},
	}

	return r, nil
}
