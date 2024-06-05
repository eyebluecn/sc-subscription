package domain

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/infra/rpc"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/info"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"time"
)

type OrderDomainSvc struct{}

func NewOrderDomainSvc() *OrderDomainSvc {
	return &OrderDomainSvc{}
}

// 订单编号生成
func (receiver OrderDomainSvc) GenerateOrderNo(ctx context.Context) string {
	return fmt.Sprintf("SCO%v%v", time.Now().UnixMilli(), util.RandomInt(1000, 9999))
}

// 新建订单关系，并且准备好支付的内容。
func (receiver OrderDomainSvc) CreateAndPrepare(ctx context.Context,
	columnVO vo.ColumnVO,
	readerVO vo.ReaderVO,
	payMethod string) (*info.PrepareSubscribeInfo, error) {

	var paymentPrepareInfo *info.PaymentPrepareInfo

	//查询订单是否已经存在。
	order, err := receiver.QueryUnpaidOrder(ctx, columnVO, readerVO)
	if err != nil {
		return nil, err
	}

	if order == nil {
		//没有待支付的订单。 获取专栏报价。
		columnQuoteVO, err := rpc.NewMiscCaller().ColumnQuoteCheckByColumnId(ctx, columnVO.ID)
		if err != nil {
			return nil, err
		}

		orderNo := NewOrderDomainSvc().GenerateOrderNo(ctx)

		//创建对应支付单。
		paymentPrepareInfo, err = rpc.NewMiscCaller().PaymentCreate(ctx, orderNo, payMethod, columnQuoteVO.Price)
		if err != nil {
			return nil, err
		}

		paymentVO := paymentPrepareInfo.PaymentVO

		//创建订单。
		order = &do.OrderDO{
			ID:            0,
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
			No:            orderNo,
			ReaderID:      readerVO.ID,
			ColumnID:      columnVO.ID,
			ColumnQuoteID: columnQuoteVO.ID,
			PaymentID:     paymentVO.ID,
			Price:         columnQuoteVO.Price,
			Status:        enums.OrderStatusCreated,
		}

		order, err = repo.NewOrderRepo().Insert(ctx, order)
		if err != nil {
			return nil, err
		}

	} else {
		//已有待支付的订单。
		paymentPrepareInfo, err = rpc.NewMiscCaller().PaymentPrepare(ctx, order.PaymentID)
		if err != nil {
			return nil, err
		}
	}

	resp := &info.PrepareSubscribeInfo{
		Order:              order,
		ThirdTransactionNo: paymentPrepareInfo.ThirdTransactionNo,
		NonceStr:           paymentPrepareInfo.NonceStr,
	}

	return resp, nil
}

// 订单标记为已支付
func (receiver OrderDomainSvc) OrderPaid(ctx context.Context, order *do.OrderDO) (*do.OrderDO, error) {

	rowsAffected, err := repo.NewOrderRepo().UpdateStatus(ctx, order.ID, enums.OrderStatusPaid)
	if err != nil {
		return nil, err
	}

	klog.CtxInfof(ctx, "订单id=%v no=%v rowsAffected=%v 状态已更新成PAID", order.ID, order.No, rowsAffected)
	//查询最新的。
	order, err = repo.NewOrderRepo().CheckByNo(ctx, order.No)
	if err != nil {
		return nil, err
	}

	return order, nil
}

// 查询某个用户在某个专栏上，未完成的订单。
func (receiver OrderDomainSvc) QueryUnpaidOrder(ctx context.Context, columnVO vo.ColumnVO, readerVO vo.ReaderVO) (*do.OrderDO, error) {

	var order *do.OrderDO
	//2.查询当前订单是否已经存在了。
	orders, err := repo.NewOrderRepo().QueryByReaderIdAndColumnIdAndStatuses(ctx, readerVO.ID, columnVO.ID, []enums.OrderStatus{enums.OrderStatusCreated})
	if err != nil {
		return nil, err
	}
	if len(orders) > 0 {
		order = orders[0]
	}

	return order, nil
}
