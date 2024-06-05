package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/do2po"
	"github.com/eyebluecn/sc-misc/src/converter/po2do"
	"github.com/eyebluecn/sc-misc/src/model/do"
	"github.com/eyebluecn/sc-misc/src/model/do/enums"
	"github.com/eyebluecn/sc-misc/src/model/query"
	"github.com/eyebluecn/sc-misc/src/model/result"
	"github.com/eyebluecn/sc-misc/src/repository/config"
	"github.com/eyebluecn/sc-misc/src/repository/dao"
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
	payment *do.OrderDO,
) (*do.OrderDO, error) {
	table := dao.Use(config.DB).OrderPO

	//时间置为当前
	payment.CreateTime = time.Now()
	payment.UpdateTime = time.Now()

	paymentDO := do2po.ConvertOrderPO(payment)

	err := table.WithContext(ctx).Debug().Create(paymentDO)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return nil, err
	}

	return po2do.ConvertOrderDO(paymentDO), nil
}

// 按照分页查询 1基
func (receiver OrderRepo) Page(
	ctx context.Context,
	req query.OrderPageQuery,
) (list []*do.OrderDO, pagination *result.Pagination, err error) {

	table := dao.Use(config.DB).OrderPO
	conditions := make([]gen.Condition, 0)

	if req.ReaderId != nil {
		conditions = append(conditions, table.ReaderID.Eq(*req.ReaderId))
	}
	if req.ColumnId != nil {
		conditions = append(conditions, table.ColumnID.Eq(*req.ColumnId))
	}

	if len(req.Statuses) != 0 {
		status := do2po.ConvertOrderStatuses(req.Statuses)
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

	pagination = &result.Pagination{
		PageNum:    req.PageNum,
		PageSize:   req.PageSize,
		TotalItems: total,
	}
	return po2do.ConvertOrderDOs(pageData), pagination, nil
}

// 根据订单号查询订单 找不到返回nil.
func (receiver OrderRepo) QueryByNo(
	ctx context.Context,
	no string,
) (*do.OrderDO, error) {
	table := dao.Use(config.DB).OrderPO
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
	return po2do.ConvertOrderDO(orderDO), nil
}

// 根据订单号查询订单 找不到返回nil.
func (receiver OrderRepo) CheckByNo(
	ctx context.Context,
	no string,
) (*do.OrderDO, error) {
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
	statuses []enums.OrderStatus,
) ([]*do.OrderDO, error) {
	table := dao.Use(config.DB).OrderPO
	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ReaderID.Eq(readerId))
	conditions = append(conditions, table.ColumnID.Eq(columnId))

	if len(statuses) != 0 {
		status := do2po.ConvertOrderStatuses(statuses)
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
	return po2do.ConvertOrderDOs(orderList), nil
}

// 更新状态
func (receiver OrderRepo) UpdateStatus(
	ctx context.Context,
	orderId int64,
	orderStatus enums.OrderStatus,
) (int64, error) {
	table := dao.Use(config.DB).OrderPO

	conditions := make([]gen.Condition, 0)

	conditions = append(conditions, table.ID.Eq(orderId))

	dbOrderStatus := do2po.ConvertOrderStatus(orderStatus)

	info, err := table.WithContext(ctx).Debug().Where(conditions...).Update(table.Status, dbOrderStatus)
	if err != nil {
		klog.CtxErrorf(ctx, "db repo error %v", err)
		return 0, err
	}

	return info.RowsAffected, nil
}

// 按照id查找，找不到返回nil
func (receiver OrderRepo) QueryById(
	ctx context.Context,
	id int64,
) (*do.OrderDO, error) {
	table := dao.Use(config.DB).OrderPO

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
	return po2do.ConvertOrderDO(orderDO), nil
}

// 按照id查找，找不到返回err
func (receiver OrderRepo) CheckById(
	ctx context.Context,
	id int64,
) (*do.OrderDO, error) {
	order, err := receiver.QueryById(ctx, id)
	if err != nil {
		return nil, err
	}
	if order == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, fmt.Sprintf("id=%d对应的订单不存在", id))
	}
	return order, nil
}

// 根据id批量查询
func (receiver OrderRepo) QueryByIds(
	ctx context.Context,
	ids []int64,
) (list []*do.OrderDO, err error) {

	table := dao.Use(config.DB).OrderPO
	conditions := make([]gen.Condition, 0)

	if len(ids) > 0 {
		conditions = append(conditions, table.ID.In(ids...))
	} else {
		return nil, errs.CodeErrorf(errs.StatusCodeParamsError, "ids列表不能为空")
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

	return po2do.ConvertOrderDOs(listData), nil
}
