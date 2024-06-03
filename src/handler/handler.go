package handler

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/handler/realization"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// 定义一个变量，方便快速查看待实现的方法
var _ sc_subscription_api.SubscriptionService

// 定义一个服务实现类
type SubscriptionServiceImpl struct{}

// 请求订阅列表
func (receiver *SubscriptionServiceImpl) SubscriptionPage(ctx context.Context, request *sc_subscription_api.SubscriptionPageRequest) (r *sc_subscription_api.SubscriptionPageResponse, err error) {

	klog.Infof("SubscriptionPage request: %v", request)
	r, err = realization.NewSubscriptionPage().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_subscription_api.SubscriptionPageResponse{}, err).(*sc_subscription_api.SubscriptionPageResponse), nil
}
func (receiver *SubscriptionServiceImpl) SubscriptionRichPage(ctx context.Context, request *sc_subscription_api.SubscriptionRichPageRequest) (r *sc_subscription_api.SubscriptionRichPageResponse, err error) {

	klog.Infof("SubscriptionRichPage request: %v", request)
	r, err = realization.NewSubscriptionRichPage().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_subscription_api.SubscriptionRichPageResponse{}, err).(*sc_subscription_api.SubscriptionRichPageResponse), nil
}

// 请求订阅列表
func (receiver *SubscriptionServiceImpl) SubscriptionPrepare(ctx context.Context, request *sc_subscription_api.SubscriptionPrepareRequest) (r *sc_subscription_api.SubscriptionPrepareResponse, err error) {

	klog.Infof("SubscriptionPrepare request: %v", request)
	r, err = realization.NewSubscriptionPrepare().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_subscription_api.SubscriptionPrepareResponse{}, err).(*sc_subscription_api.SubscriptionPrepareResponse), nil
}

// MQ消息入口本应在consumer中，这里模拟mq，入口从RPC入口过来。
func (receiver *SubscriptionServiceImpl) MqMessageArrive(ctx context.Context, request *sc_subscription_api.MqMessageArriveRequest) (r *sc_subscription_api.MqMessageArriveResponse, err error) {

	klog.Infof("MqMessageArrive request: %v", request)
	r, err = realization.NewMqMessageArrive().Handle(ctx, request)

	//统一包装后返回
	return FinalWrapResponse(ctx, r, &sc_subscription_api.MqMessageArriveResponse{}, err).(*sc_subscription_api.MqMessageArriveResponse), nil
}
