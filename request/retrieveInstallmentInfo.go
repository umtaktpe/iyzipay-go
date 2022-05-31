package request

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

type RetrieveInstallmentInfo struct {
	model.BaseModel
	binNumber string
	price     string
	currency  string
}

func NewRetrieveInstallmentInfo() *RetrieveInstallmentInfo {
	return &RetrieveInstallmentInfo{}
}

func (r *RetrieveInstallmentInfo) GetBinNumber() string {
	return r.binNumber
}

func (r *RetrieveInstallmentInfo) SetBinNumber(binNumber string) {
	r.binNumber = binNumber
}

func (r *RetrieveInstallmentInfo) GetPrice() string {
	return r.price
}

func (r *RetrieveInstallmentInfo) SetPrice(price string) {
	r.price = price
}

func (r *RetrieveInstallmentInfo) GetCurrency() string {
	return r.currency
}

func (r *RetrieveInstallmentInfo) SetCurrency(currency string) {
	r.currency = currency
}

func (r *RetrieveInstallmentInfo) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("binNumber", r.GetBinNumber())
	jsonBuilder.AddPrice("price", r.GetPrice())
	jsonBuilder.Add("currency", r.GetCurrency())

	return jsonBuilder.GetObject()
}

func (r *RetrieveInstallmentInfo) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("binNumber", r.GetBinNumber())
	requestStringBuilder.AppendPrice("price", r.GetPrice())
	requestStringBuilder.Append("currency", r.GetCurrency())

	return requestStringBuilder.GetRequestString()
}
