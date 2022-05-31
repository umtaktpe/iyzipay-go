package request

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

type CancelPayment struct {
	model.BaseModel
	paymentId   string
	ip          string
	reason      string
	description string
}

func NewCancelPayment() *CancelPayment {
	return &CancelPayment{}
}

func (c *CancelPayment) GetPaymentId() string {
	return c.paymentId
}

func (c *CancelPayment) SetPaymentId(paymentId string) {
	c.paymentId = paymentId
}

func (c *CancelPayment) GetIp() string {
	return c.ip
}

func (c *CancelPayment) SetIp(ip string) {
	c.ip = ip
}

func (c *CancelPayment) GetReason() string {
	return c.reason
}

func (c *CancelPayment) SetReason(reason string) {
	c.reason = reason
}

func (c *CancelPayment) GetDescription() string {
	return c.description
}

func (c *CancelPayment) SetDescription(description string) {
	c.description = description
}

func (c *CancelPayment) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("paymentId", c.GetPaymentId())
	jsonBuilder.Add("ip", c.GetIp())
	jsonBuilder.Add("reason", c.GetReason())
	jsonBuilder.Add("description", c.GetDescription())

	return jsonBuilder.GetObject()
}

func (c *CancelPayment) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("paymentId", c.GetPaymentId())
	requestStringBuilder.Append("ip", c.GetIp())
	requestStringBuilder.Append("reason", c.GetReason())
	requestStringBuilder.Append("description", c.GetDescription())

	return requestStringBuilder.GetRequestString()
}
