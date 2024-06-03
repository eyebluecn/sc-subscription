package application

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/util"
)

type PaymentApp struct{}

func NewPaymentApp() *PaymentApp {
	return &PaymentApp{}
}

// 支付单支付成功。
func (receiver PaymentApp) PaymentPaidCallback(ctx context.Context, keys string, body string) error {

	klog.CtxInfof(ctx, "Payment Paid Callback keys: %s, body: %s", keys, body)

	//反序列化。
	var payload sc_misc_api.PaymentMqPayload
	_, err := util.ParseJSON(body, &payload)
	if err != nil {
		return errs.BadRequestErrorf("无法识别payload %v", err.Error())
	}

	if payload.Tags == sc_misc_api.PaymentMqEvent_PAYMENT_PAID.String() {

		subscription, err := NewSubscriptionWriteApp().PaymentPaid(ctx, payload.PaymentDTO.Id)
		if err != nil {
			return err
		}

		klog.CtxErrorf(ctx, "订阅关系已成功创建： %v", subscription.ID)

	} else {
		klog.CtxErrorf(ctx, "收到的消息不是 PaymentMqEvent_PAYMENT_PAID 放弃处理")
	}

	return nil
}
