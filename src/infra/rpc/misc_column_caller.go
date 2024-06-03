package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/vo_conv"
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
)

// 根据id获取专栏，可能为空。
// 如果err==nil，则ColumnVO!=nil
func (receiver MiscCaller) ColumnQueryById(ctx context.Context, columnId int64) (*vo_model.ColumnVO, error) {
	request := &sc_misc_api.ColumnQueryByIdRequest{
		ColumnId: columnId,
	}
	response, err := config.MiscRpcClient.ColumnQueryById(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ColumnQueryById failed: %v", err)
		return nil, err
	}

	columnVO := vo_conv.ConvertColumnVO(response.Data)

	return columnVO, nil
}

// 根据id获取读者，可能为空。
// 如果err==nil，则ColumnVO!=nil
func (receiver MiscCaller) ColumnQueryByIds(ctx context.Context, columnIds []int64) ([]*vo_model.ColumnVO, error) {
	request := &sc_misc_api.ColumnQueryByIdsRequest{
		ColumnIds: columnIds,
	}
	response, err := config.MiscRpcClient.ColumnQueryByIds(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ColumnQueryByIds failed: %v", err)
		return nil, err
	}

	resultList := vo_conv.ConvertColumnVOs(response.Data)

	return resultList, nil
}

// 根据id获取专栏，如果为nil，返回报错。
func (receiver MiscCaller) ColumnCheckById(ctx context.Context, columnId int64) (*vo_model.ColumnVO, error) {
	columnVO, err := receiver.ColumnQueryById(ctx, columnId)
	if err != nil {
		return nil, err
	}

	if columnVO == nil {
		return nil, errs.CodeErrorf(enums.StatusCodeNotFound, "id=%v的记录不存在", columnId)
	}

	return columnVO, nil
}
