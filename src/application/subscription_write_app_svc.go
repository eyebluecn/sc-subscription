package application

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/domain"
	"github.com/eyebluecn/sc-misc/src/infra/rpc"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/info"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type SubscriptionWriteAppSvc struct{}

func NewSubscriptionWriteAppSvc() *SubscriptionWriteAppSvc {
	return &SubscriptionWriteAppSvc{}
}

// 准备订阅
func (receiver SubscriptionWriteAppSvc) PrepareSubscribe(ctx context.Context, readerId int64, columnId int64, payMethod string) (*info.PrepareSubscribeInfo, error) {

	//1. 查询读者信息
	readerVO, err := rpc.NewMiscCaller().ReaderCheckById(ctx, readerId)
	if err != nil {
		return nil, err
	}

	//2. 查询专栏信息
	columnVO, err := rpc.NewMiscCaller().ColumnCheckById(ctx, columnId)
	if err != nil {
		return nil, err
	}

	//3. 下发一个准备单
	prepareSubscribeInfo, err := domain.NewOrderDomainSvc().CreateAndPrepare(ctx, *columnVO, *readerVO, payMethod)
	if err != nil {
		return nil, err
	}

	return prepareSubscribeInfo, nil
}

// 支付单已完成支付。
func (receiver SubscriptionWriteAppSvc) PaymentPaid(ctx context.Context, paymentId int64) (*do.SubscriptionDO, error) {
	klog.CtxInfof(ctx, "收到支付完成的信息了，准备去确认订单和创建订阅关系。 paymentId=%v", paymentId)

	//反查支付单
	paymentVO, err := rpc.NewMiscCaller().PaymentCheckById(ctx, paymentId)
	if err != nil {
		return nil, err
	}

	//查询订单
	order, err := repo.NewOrderRepo().CheckByNo(ctx, paymentVO.OrderNo)
	if err != nil {
		return nil, err
	}

	//查询读者
	readerVO, err := rpc.NewMiscCaller().ReaderCheckById(ctx, order.ReaderID)
	if err != nil {
		return nil, err
	}

	//查询专栏
	columnVO, err := rpc.NewMiscCaller().ColumnCheckById(ctx, order.ColumnID)
	if err != nil {
		return nil, err
	}

	//查询订阅关系是否已经建立了。
	subscription, err := repo.NewSubscriptionRepo().QueryByReaderIdAndColumnId(ctx, readerVO.ID, columnVO.ID)
	if err != nil {
		return nil, err
	}
	if subscription != nil {
		//说明订阅关系已经建立了，这里是重复投递。
		klog.CtxInfof(ctx, "说明订阅关系已经建立了，这里是重复投递。")
		return subscription, nil
	}

	//手动开启事务。
	{
		//订单标记为已支付。
		order, err = domain.NewOrderDomainSvc().OrderPaid(ctx, order)
		if err != nil {
			return nil, err
		}

		//创建订阅关系。
		subscription, err = domain.NewSubscriptionDomainSvc().Create(ctx, order)
		if err != nil {
			return nil, err
		}

	}
	return subscription, nil
}
