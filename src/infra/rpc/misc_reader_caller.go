package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/dto2vo"
	"github.com/eyebluecn/sc-misc/src/infra/rpc/config"
	"github.com/eyebluecn/sc-misc/src/model/vo"
)

// 根据id获取读者，可能为空。
// 如果err==nil，则ReaderVO!=nil
func (receiver MiscCaller) ReaderQueryById(ctx context.Context, readerId int64) (*vo.ReaderVO, error) {
	request := &sc_misc_api.ReaderQueryByIdRequest{
		ReaderId: readerId,
	}
	response, err := config.MiscRpcClient.ReaderQueryById(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ReaderQueryById failed: %v", err)
		return nil, err
	}

	readerVO := dto2vo.ConvertReaderVO(response.Data)

	return readerVO, nil
}

// 根据id获取读者，可能为空。
// 如果err==nil，则ReaderVO!=nil
func (receiver MiscCaller) ReaderQueryByIds(ctx context.Context, readerIds []int64) ([]*vo.ReaderVO, error) {
	request := &sc_misc_api.ReaderQueryByIdsRequest{
		ReaderIds: readerIds,
	}
	response, err := config.MiscRpcClient.ReaderQueryByIds(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "ReaderQueryByIds failed: %v", err)
		return nil, err
	}

	resultList := dto2vo.ConvertReaderVOs(response.Data)

	return resultList, nil
}

// 根据id获取读者，如果为nil，返回报错。
func (receiver MiscCaller) ReaderCheckById(ctx context.Context, readerId int64) (*vo.ReaderVO, error) {
	readerVO, err := receiver.ReaderQueryById(ctx, readerId)
	if err != nil {
		return nil, err
	}

	if readerVO == nil {
		return nil, errs.CodeErrorf(errs.StatusCodeNotFound, "id=%v的记录不存在", readerId)
	}

	return readerVO, nil
}
