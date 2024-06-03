package realization

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/infra/mq"
	"github.com/eyebluecn/sc-subscription-idl/kitex_gen/sc_subscription_api"
)

// MQ消息入口本应在consumer中，这里模拟mq，入口从RPC入口过来。
type MqMessageArrive struct{}

func NewMqMessageArrive() *MqMessageArrive {
	return &MqMessageArrive{}
}

// 服务请求入口
func (receiver MqMessageArrive) Handle(ctx context.Context, request *sc_subscription_api.MqMessageArriveRequest) (r *sc_subscription_api.MqMessageArriveResponse, err error) {

	err = receiver.CheckParam(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "参数校验出错：%v", err)
		return nil, err
	}

	response, err := receiver.doHandle(ctx, *request)
	return response, err
}

// 校验参数
func (receiver MqMessageArrive) CheckParam(ctx context.Context, request *sc_subscription_api.MqMessageArriveRequest) error {
	if request == nil {
		return errs.BadRequestErrorf("request 不能为空")
	}
	if request.Topic == "" {
		return errs.BadRequestErrorf("request.Topic 不能为空")
	}
	if request.Body == "" {
		return errs.BadRequestErrorf("request.Body 不能为空")
	}

	return nil
}

// 参数校验后的真实处理
func (receiver MqMessageArrive) doHandle(ctx context.Context, request sc_subscription_api.MqMessageArriveRequest) (r *sc_subscription_api.MqMessageArriveResponse, err error) {
	//在这里翻译成对应的消息，投递到Consumer中去。
	err = mq.DefaultConsumer().Receive(ctx, request.Topic, request.Tags, request.Keys, request.Body)
	if err != nil {
		return nil, err
	}

	resp := &sc_subscription_api.MqMessageArriveResponse{}

	return resp, nil
}
