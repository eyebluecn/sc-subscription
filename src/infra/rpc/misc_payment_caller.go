package rpc

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/eyebluecn/sc-misc-idl/kitex_gen/sc_misc_api"
	"github.com/eyebluecn/sc-misc/src/common/config"
	"github.com/eyebluecn/sc-misc/src/common/enums"
	"github.com/eyebluecn/sc-misc/src/common/errs"
	"github.com/eyebluecn/sc-misc/src/converter/vo_conv"
	"github.com/eyebluecn/sc-misc/src/model/info_model"
	"github.com/eyebluecn/sc-misc/src/model/vo_model"
)

// 根据id获取支付单，可能为空。
// 如果err==nil，则PaymentVO!=nil
func (receiver MiscCaller) PaymentQueryById(ctx context.Context, paymentId int64) (*vo_model.PaymentVO, error) {
	request := &sc_misc_api.PaymentQueryByIdRequest{
		PaymentId: paymentId,
	}
	response, err := config.MiscRpcClient.PaymentQueryById(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "PaymentQueryById failed: %v", err)
		return nil, err
	}

	paymentVO := vo_conv.ConvertPaymentVO(response.Data)

	return paymentVO, nil
}

// 根据id获取支付单，如果为nil，返回报错。
func (receiver MiscCaller) PaymentCheckById(ctx context.Context, paymentId int64) (*vo_model.PaymentVO, error) {
	paymentVO, err := receiver.PaymentQueryById(ctx, paymentId)
	if err != nil {
		return nil, err
	}

	if paymentVO == nil {
		return nil, errs.CodeErrorf(enums.StatusCodeNotFound, "id=%v的记录不存在", paymentId)
	}

	return paymentVO, nil
}

// 准备支付
func (receiver MiscCaller) PaymentPrepare(ctx context.Context, paymentId int64) (*info_model.PaymentPrepareInfo, error) {
	request := &sc_misc_api.PaymentPrepareRequest{
		PaymentId: paymentId,
	}
	response, err := config.MiscRpcClient.PaymentPrepare(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "PaymentPrepare failed: %v", err)
		return nil, err
	}
	if response.Data == nil {
		return nil, errs.BadRequestErrorf("response data is nil")
	}

	paymentVO := vo_conv.ConvertPaymentVO(response.Data.PaymentDTO)

	resp := &info_model.PaymentPrepareInfo{
		PaymentVO:          paymentVO,
		ThirdTransactionNo: response.Data.ThirdTransactionNo,
		NonceStr:           response.Data.NonceStr,
	}

	return resp, nil
}

// 创建一个支付单同时返回支付准备物料等信息
func (receiver MiscCaller) PaymentCreate(ctx context.Context,
	orderNo string,
	method string,
	amount int64) (*info_model.PaymentPrepareInfo, error) {
	request := &sc_misc_api.PaymentCreateRequest{
		OrderNo: orderNo,
		Method:  method,
		Amount:  amount,
	}
	response, err := config.MiscRpcClient.PaymentCreate(ctx, request)
	if err != nil {
		klog.CtxErrorf(ctx, "PaymentCreate failed: %v", err)
		return nil, err
	}
	if response.Data == nil {
		return nil, errs.BadRequestErrorf("response data is nil")
	}

	paymentVO := vo_conv.ConvertPaymentVO(response.Data.PaymentDTO)

	resp := &info_model.PaymentPrepareInfo{
		PaymentVO:          paymentVO,
		ThirdTransactionNo: response.Data.ThirdTransactionNo,
		NonceStr:           response.Data.NonceStr,
	}

	return resp, nil
}
