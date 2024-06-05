package application

import (
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/infra/mq"
)

type BootstrapAppSvc struct{}

// 应用中一些初始化工作。
func InitApplication() {
	//注册mq消费者接收方法。
	mq.DefaultConsumer().Register(sc_misc_api.PaymentMqEvent_PAYMENT_PAID.String(), NewPaymentAppSvc().PaymentPaidCallback)
}
