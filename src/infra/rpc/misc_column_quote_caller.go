package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/vo_conv"
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
)

// 根据id获取读者，可能为空。
// 如果err==nil，则ReaderVO!=nil
func (receiver MiscCaller) ColumnQuoteQueryByColumnId(ctx context.Context, columnQuoteId int64) (*vo_model.ColumnQuoteVO, error) {
	request := &sc_misc_api.ColumnQuoteQueryByColumnIdRequest{
		ColumnId: columnQuoteId,
	}
	response, err := config.MiscRpcClient.ColumnQuoteQueryByColumnId(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ColumnQuoteQueryByColumnId failed: %v", err)
		return nil, err
	}

	columnQuoteVO := vo_conv.ConvertColumnQuoteVO(response.Data)

	return columnQuoteVO, nil
}

// 根据id获取读者，如果为nil，返回报错。
func (receiver MiscCaller) ColumnQuoteCheckByColumnId(ctx context.Context, columnQuoteId int64) (*vo_model.ColumnQuoteVO, error) {
	columnQuoteVO, err := receiver.ColumnQuoteQueryByColumnId(ctx, columnQuoteId)
	if err != nil {
		return nil, err
	}

	if columnQuoteVO == nil {
		return nil, errs.NotFoundErrorf("id=%v的记录不存在", columnQuoteId)
	}

	return columnQuoteVO, nil
}
