package request

import "github.com/umtaktpe/iyzipay-go/utils"

type RetrievePayment struct {
	paymentId             string
	paymentConversationId string
}

func NewRetrievePayment() *RetrievePayment {
	return &RetrievePayment{}
}

func (r *RetrievePayment) GetPaymentId() string {
	return r.paymentId
}

func (r *RetrievePayment) SetPaymentId(paymentId string) {
	r.paymentId = paymentId
}

func (r *RetrievePayment) GetPaymentConversationId() string {
	return r.paymentConversationId
}

func (r *RetrievePayment) SetPaymentConversationId(paymentConversationId string) {
	r.paymentConversationId = paymentConversationId
}

func (r *RetrievePayment) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("paymentId", r.GetPaymentId())
	jsonBuilder.Add("paymentConversationId", r.GetPaymentConversationId())

	return jsonBuilder.GetObject()
}

func (r *RetrievePayment) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("paymentId", r.GetPaymentId())
	requestStringBuilder.Append("paymentConversationId", r.GetPaymentConversationId())

	return requestStringBuilder.GetRequestString()
}
