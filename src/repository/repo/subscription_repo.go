package repo

import (
	"context"
	"errors"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/db_model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/model_conv"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type SubscriptionRepo struct {
}

func NewSubscriptionRepo() SubscriptionRepo {
	return SubscriptionRepo{}
}

// 新建一个Subscription
func (receiver SubscriptionRepo) Insert(
	ctx context.Context,
	payment *model.Subscription,
) (*model.Subscription, error) {
	table := dao.Use(config.DB).SubscriptionDO

	//时间置为当前
	payment.CreateTime = time.Now()
	payment.UpdateTime = time.Now()

	paymentDO := db_model_conv.ConvertSubscriptionDO(payment)

	err := table.WithContext(ctx).Debug().Create(paymentDO)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, err
	}

	return model_conv.ConvertSubscription(paymentDO), nil
}

// 按照分页查询 1基
func (receiver SubscriptionRepo) Page(
	ctx context.Context,
	req SubscriptionPageRequest,
) (list []*model.Subscription, pagination *model.Pagination, err error) {

	table := dao.Use(config.DB).SubscriptionDO
	conditions := make([]gen.Condition, 0)

	if !req.CreateTimeGte.IsZero() {
		conditions = append(conditions, table.CreateTime.Gte(req.CreateTimeGte))
	}

	if req.ReaderId != nil {
		conditions = append(conditions, table.ReaderID.Eq(*req.ReaderId))
	}

	if len(req.ColumnIds) != 0 {
		conditions = append(conditions, table.ColumnID.In(req.ColumnIds...))
	}

	if req.OrderId != nil {
		conditions = append(conditions, table.OrderID.Eq(*req.OrderId))
	}

	if req.Status != nil {
		status := db_model_conv.SubscriptionStatusToStorage(*req.Status)
		conditions = append(conditions, table.Status.Eq(status))
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	//默认按照创建时间倒序排列
	orderExpr := table.CreateTime.Desc()
	if req.OrderBy == OrderByCreateTimeAsc {
		orderExpr = table.CreateTime.Asc()
	} else if req.OrderBy == OrderByIdDesc {
		orderExpr = table.CreateTime.Desc()
	} else if req.OrderBy == OrderByIdAsc {
		orderExpr = table.CreateTime.Asc()
	}
	tableDO = tableDO.Order(orderExpr)

	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = 10
	}
	if req.PageSize > 100 {
		return nil, nil, errs.BadRequestErrorf("PageSize不能大于100")
	}

	offset := (req.PageNum - 1) * req.PageSize
	limit := req.PageSize
	pageData, total, err := tableDO.FindByPage(int(offset), int(limit))
	if err != nil {
		klog.CtxErrorf(ctx, "PageQuery failed, err=%v", err)
		return nil, nil, err
	}

	pagination = &model.Pagination{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		TotalItems: total,
	}
	return model_conv.ConvertSubscriptions(pageData), pagination, nil
}

// 查询某个读者和某个专栏的订阅关系。 找不到返回nil.
func (receiver SubscriptionRepo) QueryByReaderIdAndColumnId(
	ctx context.Context,
	readerId int64,
	columnId int64,
) (*model.Subscription, error) {
	table := dao.Use(config.DB).SubscriptionDO
	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ReaderID.Eq(readerId))
	conditions = append(conditions, table.ColumnID.Eq(columnId))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	subscriptionDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return model_conv.ConvertSubscription(subscriptionDO), nil
}
