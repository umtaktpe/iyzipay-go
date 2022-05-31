package api

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

func InstallmentInfo(req utils.RequestConvertible, opt *model.Options) (*installmentInfo, error) {
	resp := &installmentInfo{}
	if err := request("POST", "/payment/iyzipos/installment", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func BinNumber(req utils.RequestConvertible, opt *model.Options) (*binNumber, error) {
	resp := &binNumber{}
	if err := request("POST", "/payment/bin/check", req, opt, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
