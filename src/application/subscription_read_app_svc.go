package application

import (
	"context"
	"github.com/eyebluecn/sc-misc/src/common/util"
	"github.com/eyebluecn/sc-misc/src/infra/rpc"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/info"
	"github.com/eyebluecn/sc-misc/src/model/query"
	"github.com/eyebluecn/sc-misc/src/model/universal"
	"github.com/eyebluecn/sc-misc/src/model/vo"
	"github.com/eyebluecn/sc-misc/src/repository/repo"
)

type SubscriptionReadAppSvc struct{}

func NewSubscriptionReadAppSvc() *SubscriptionReadAppSvc {
	return &SubscriptionReadAppSvc{}
}

// 获取某位读者查看到的订阅列表。
func (receiver SubscriptionReadAppSvc) RichSubscriptionPage(ctx context.Context, repoRequest query.SubscriptionPageQuery) ([]*info.RichSubscription, *universal.Pagination, error) {

	subscriptions, pagination, err := repo.NewSubscriptionRepo().Page(ctx, repoRequest)
	if err != nil {
		return nil, nil, err
	}

	var richSubscriptions []*info.RichSubscription
	//装填专栏骨架
	for _, subscription := range subscriptions {
		richSubscriptions = append(richSubscriptions, &info.RichSubscription{
			Subscription: subscription,
		})
	}

	//依次装填. 这里可以做成并行填充。
	err = receiver.PopulateColumn(ctx, richSubscriptions)
	if err != nil {
		return nil, nil, err
	}

	err = receiver.PopulateReader(ctx, richSubscriptions)
	if err != nil {
		return nil, nil, err
	}

	err = receiver.PopulateOrder(ctx, richSubscriptions)
	if err != nil {
		return nil, nil, err
	}

	return richSubscriptions, pagination, nil
}

// 填充专栏信息
func (receiver SubscriptionReadAppSvc) PopulateColumn(ctx context.Context, richSubscriptions []*info.RichSubscription) error {

	var columnIds []int64
	for _, richSubscription := range richSubscriptions {
		columnIds = util.UniqueAddInt64(columnIds, richSubscription.Subscription.ColumnID)
	}

	if len(columnIds) == 0 {
		return nil
	}

	list, err := rpc.NewMiscCaller().ColumnQueryByIds(ctx, columnIds)
	if err != nil {
		return err
	}

	//从list中找到合适的Author.
	for _, richSubscription := range richSubscriptions {
		column := receiver.findColumn(ctx, list, richSubscription.Subscription.ColumnID)
		if column != nil {
			richSubscription.Column = column
		}
	}

	return nil

}

// 从列表中找到对应的读者
func (receiver SubscriptionReadAppSvc) findColumn(ctx context.Context, columnList []*vo.ColumnVO, columnId int64) *vo.ColumnVO {
	for _, column := range columnList {
		if column.ID == columnId {
			return column
		}
	}
	return nil
}

// 填充读者信息信息
func (receiver SubscriptionReadAppSvc) PopulateReader(ctx context.Context, richSubscriptions []*info.RichSubscription) error {

	var readerIds []int64
	for _, richSubscription := range richSubscriptions {
		readerIds = util.UniqueAddInt64(readerIds, richSubscription.Subscription.ReaderID)
	}

	if len(readerIds) == 0 {
		return nil
	}

	list, err := rpc.NewMiscCaller().ReaderQueryByIds(ctx, readerIds)
	if err != nil {
		return err
	}

	//从list中找到合适的Author.
	for _, richSubscription := range richSubscriptions {
		reader := receiver.findReader(ctx, list, richSubscription.Subscription.ReaderID)
		if reader != nil {
			richSubscription.Reader = reader
		}
	}

	return nil

}

// 从列表中找到对应的读者
func (receiver SubscriptionReadAppSvc) findReader(ctx context.Context, readerList []*vo.ReaderVO, readerId int64) *vo.ReaderVO {
	for _, reader := range readerList {
		if reader.ID == readerId {
			return reader
		}
	}
	return nil
}

// 填充读者信息信息
func (receiver SubscriptionReadAppSvc) PopulateOrder(ctx context.Context, richSubscriptions []*info.RichSubscription) error {

	var orderIds []int64
	for _, richSubscription := range richSubscriptions {
		orderIds = util.UniqueAddInt64(orderIds, richSubscription.Subscription.OrderID)
	}

	if len(orderIds) == 0 {
		return nil
	}

	list, err := repo.NewOrderRepo().QueryByIds(ctx, orderIds)
	if err != nil {
		return err
	}

	//从list中找到合适的Author.
	for _, richSubscription := range richSubscriptions {
		order := receiver.findOrder(ctx, list, richSubscription.Subscription.OrderID)
		if order != nil {
			richSubscription.Order = order
		}
	}

	return nil

}

// 从列表中找到对应的读者
func (receiver SubscriptionReadAppSvc) findOrder(ctx context.Context, orderList []*do.OrderDO, orderId int64) *do.OrderDO {
	for _, order := range orderList {
		if order.ID == orderId {
			return order
		}
	}
	return nil
}
