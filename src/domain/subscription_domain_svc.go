package domain

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"time"
)

type SubscriptionDomainSvc struct{}

func NewSubscriptionDomainSvc() *SubscriptionDomainSvc {
	return &SubscriptionDomainSvc{}
}

// 创建专栏
func (service *SubscriptionDomainSvc) Create(ctx context.Context, order *do.OrderDO) (*do.SubscriptionDO, error) {
	//创建一个新的订阅关系。
	subscription := &do.SubscriptionDO{
		ID:         0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		ReaderID:   order.ReaderID,
		ColumnID:   order.ColumnID,
		OrderID:    order.ID,
		Status:     enums.SubscriptionStatusOk,
	}
	subscription, err := repo.NewSubscriptionRepo().Insert(ctx, subscription)
	if err != nil {
		return nil, err
	}
	klog.CtxInfof(ctx, "专栏订阅成功：readerId=%v columnId=%v", order.ReaderID, order.ColumnID)

	//TODO: 发送专栏订阅成功的领域事件。

	return subscription, nil
}
