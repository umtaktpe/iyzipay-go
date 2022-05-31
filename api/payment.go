package api

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

func RetrievePayment(req utils.RequestConvertible, opt *model.Options) (*retrievePayment, error) {
	resp := &retrievePayment{}
	if err := request("POST", "/payment/detail", req, opt, resp); err != nil {
		return nil, err
	}

	return resp, nil
}

func CreatePayment(req utils.RequestConvertible, opt *model.Options) (*createPayment, error) {
	resp := &createPayment{}
	if err := request("POST", "/payment/auth", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func InitializeThreedsPayment(req utils.RequestConvertible, opt *model.Options) (*initializeThreedsPayment, error) {
	resp := &initializeThreedsPayment{}
	if err := request("POST", "/payment/3dsecure/initialize", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func CreateThreedsPayment(req utils.RequestConvertible, opt *model.Options) (*createPayment, error) {
	resp := &createPayment{}
	if err := request("POST", "/payment/3dsecure/auth", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func CancelPayment(req utils.RequestConvertible, opt *model.Options) (*cancelPayment, error) {
	resp := &cancelPayment{}
	if err := request("POST", "/payment/cancel", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func RefundPayment(req utils.RequestConvertible, opt *model.Options) (*refundPayment, error) {
	resp := &refundPayment{}
	if err := request("POST", "/payment/refund", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
