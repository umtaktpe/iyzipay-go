package request

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

type RefundPayment struct {
	model.BaseModel
	paymentTransactionId string
	price                string
	ip                   string
	currency             string
	reason               string
	description          string
}

func NewRefundPayment() *RefundPayment {
	return &RefundPayment{}
}

func (r *RefundPayment) GetPaymentTransactionId() string {
	return r.paymentTransactionId
}

func (r *RefundPayment) SetPaymentTransactionId(paymentTransactionId string) {
	r.paymentTransactionId = paymentTransactionId
}

func (r *RefundPayment) GetPrice() string {
	return r.price
}

func (r *RefundPayment) SetPrice(price string) {
	r.price = price
}

func (r *RefundPayment) GetIp() string {
	return r.ip
}

func (r *RefundPayment) SetIp(ip string) {
	r.ip = ip
}

func (r *RefundPayment) GetCurrency() string {
	return r.currency
}

func (r *RefundPayment) SetCurrency(currency string) {
	r.currency = currency
}

func (r *RefundPayment) GetReason() string {
	return r.reason
}

func (r *RefundPayment) SetReason(reason string) {
	r.reason = reason
}

func (r *RefundPayment) GetDescription() string {
	return r.description
}

func (r *RefundPayment) SetDescription(description string) {
	r.description = description
}

func (r *RefundPayment) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("paymentTransactionId", r.GetPaymentTransactionId())
	jsonBuilder.Add("price", r.GetPrice())
	jsonBuilder.Add("ip", r.GetIp())
	jsonBuilder.Add("currency", r.GetCurrency())
	jsonBuilder.Add("reason", r.GetReason())
	jsonBuilder.Add("description", r.GetDescription())

	return jsonBuilder.GetObject()
}

func (r *RefundPayment) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("paymentTransactionId", r.GetPaymentTransactionId())
	requestStringBuilder.Append("price", r.GetPrice())
	requestStringBuilder.Append("ip", r.GetIp())
	requestStringBuilder.Append("currency", r.GetCurrency())
	requestStringBuilder.Append("reason", r.GetReason())
	requestStringBuilder.Append("description", r.GetDescription())

	return requestStringBuilder.GetRequestString()
}
