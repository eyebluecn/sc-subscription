package domain

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
	"time"
)

type SubscriptionDomainService struct{}

func NewSubscriptionDomainService() *SubscriptionDomainService {
	return &SubscriptionDomainService{}
}

// 创建专栏
func (service *SubscriptionDomainService) Create(ctx context.Context, order *model.Order) (*model.Subscription, error) {
	//创建一个新的订阅关系。
	subscription := &model.Subscription{
		ID:         0,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		ReaderID:   order.ReaderID,
		ColumnID:   order.ColumnID,
		OrderID:    order.ID,
		Status:     model.SubscriptionStatusOk,
	}
	subscription, err := repo.NewSubscriptionRepo().Insert(ctx, subscription)
	if err != nil {
		return nil, err
	}
	klog.CtxInfof(ctx, "专栏订阅成功：readerId=%v columnId=%v", order.ReaderID, order.ColumnID)

	//TODO: 发送专栏订阅成功的领域事件。

	return subscription, nil
}
