package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/db_model_conv"
	"github.com/eyebluecn/sc-misc/src/converter/model_conv"
	"github.com/eyebluecn/sc-misc/src/model"
	"github.com/eyebluecn/sc-misc/src/repository/query"
	"gorm.io/gen"
	"gorm.io/gorm"
	"time"
)

type OrderRepo struct {
}

func NewOrderRepo() OrderRepo {
	return OrderRepo{}
}

// 新建一个Order
func (receiver OrderRepo) Insert(
	ctx context.Context,
	payment *model.Order,
) (*model.Order, error) {
	table := query.Use(config.DB).OrderDO

	//时间置为当前
	payment.CreateTime = time.Now()
	payment.UpdateTime = time.Now()

	paymentDO := db_model_conv.ConvertOrderDO(payment)

	err := table.WithContext(ctx).Debug().Create(paymentDO)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, err
	}

	return model_conv.ConvertOrder(paymentDO), nil
}

// 按照分页查询 1基
func (receiver OrderRepo) Page(
	ctx context.Context,
	req OrderPageRequest,
) (list []*model.Order, pagination *model.Pagination, err error) {

	table := query.Use(config.DB).OrderDO
	conditions := make([]gen.Condition, 0)

	if req.ReaderId != nil {
		conditions = append(conditions, table.ReaderID.Eq(*req.ReaderId))
	}
	if req.ColumnId != nil {
		conditions = append(conditions, table.ColumnID.Eq(*req.ColumnId))
	}

	if len(req.Statuses) != 0 {
		status := db_model_conv.OrderStatusesToStorage(req.Statuses)
		conditions = append(conditions, table.Status.In(status...))
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

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
	return model_conv.ConvertOrders(pageData), pagination, nil
}

// 根据订单号查询订单 找不到返回nil.
func (receiver OrderRepo) QueryByNo(
	ctx context.Context,
	no string,
) (*model.Order, error) {
	table := query.Use(config.DB).OrderDO
	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.No.Eq(no))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	orderDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return model_conv.ConvertOrder(orderDO), nil
}

// 根据订单号查询订单 找不到返回nil.
func (receiver OrderRepo) CheckByNo(
	ctx context.Context,
	no string,
) (*model.Order, error) {
	order, err := receiver.QueryByNo(ctx, no)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errs.NotFoundErrorf("无法找到订单号 %v 对应的订单", no)
	}

	return order, nil
}

// 根据readerId、columnId和状态查询。
func (receiver OrderRepo) QueryByReaderIdAndColumnIdAndStatuses(
	ctx context.Context,
	readerId int64,
	columnId int64,
	statuses []model.OrderStatus,
) ([]*model.Order, error) {
	table := query.Use(config.DB).OrderDO
	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ReaderID.Eq(readerId))
	conditions = append(conditions, table.ColumnID.Eq(columnId))

	if len(statuses) != 0 {
		status := db_model_conv.OrderStatusesToStorage(statuses)
		conditions = append(conditions, table.Status.In(status...))
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	orderList, err := tableDO.Find()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return model_conv.ConvertOrders(orderList), nil
}

// 更新状态
func (receiver OrderRepo) UpdateStatus(
	ctx context.Context,
	orderId int64,
	orderStatus model.OrderStatus,
) (int64, error) {
	table := query.Use(config.DB).OrderDO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(orderId))

	dbOrderStatus := db_model_conv.OrderStatusToStorage(orderStatus)

	info, err := table.WithContext(ctx).Debug().Where(conditions...).Update(table.Status, dbOrderStatus)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return 0, err
	}

	return info.RowsAffected, nil
}

// 按照id查找，找不到返回nil
func (receiver OrderRepo) FindById(
	ctx context.Context,
	id int64,
) (*model.Order, error) {
	table := query.Use(config.DB).OrderDO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(id))

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}
	orderDO, err := tableDO.First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			// 处理未找到记录的错误...
			return nil, nil
		}
		return nil, err
	}
	return model_conv.ConvertOrder(orderDO), nil
}

// 按照id查找，找不到返回err
func (receiver OrderRepo) CheckById(
	ctx context.Context,
	id int64,
) (*model.Order, error) {
	order, err := receiver.FindById(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errs.CodeErrorf(enums.StatusCodeNotFound, fmt.Sprintf("id=%d对应的订单不存在", id))
	}
	return order, nil
}

// 根据id批量查询
func (receiver OrderRepo) QueryByIds(
	ctx context.Context,
	ids []int64,
) (list []*model.Order, err error) {

	table := query.Use(config.DB).OrderDO
	conditions := make([]gen.Condition, 0)

	if len(ids) > 0 {
		conditions = append(conditions, table.ID.In(ids...))
	} else {
		return nil, errs.CodeErrorf(enums.StatusCodeParamsError, "ids列表不能为空")
	}

	tableDO := table.WithContext(ctx).Debug()
	if len(conditions) > 0 {
		tableDO = tableDO.Where(conditions...)
	}

	listData, err := tableDO.Find()
	if err != nil {
		klog.CtxErrorf(ctx, "FindByIds failed, err=%v", err)
		return nil, err
	}

	return model_conv.ConvertOrders(listData), nil
}
