package request

import "github.com/umtaktpe/iyzipay-go/utils"

type CreateThreedsPayment struct {
	paymentId        string
	conversationData string
}

func NewCreateThreedsPayment() *CreateThreedsPayment {
	return &CreateThreedsPayment{}
}

func (c *CreateThreedsPayment) GetPaymentId() string {
	return c.paymentId
}

func (c *CreateThreedsPayment) SetPaymentId(paymentId string) {
	c.paymentId = paymentId
}

func (c *CreateThreedsPayment) GetConversationData() string {
	return c.conversationData
}

func (c *CreateThreedsPayment) SetConversationData(conversationData string) {
	c.conversationData = conversationData
}

func (c *CreateThreedsPayment) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("paymentId", c.GetPaymentId())
	jsonBuilder.Add("conversationData", c.GetConversationData())

	return jsonBuilder.GetObject()
}

func (c *CreateThreedsPayment) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("paymentId", c.GetPaymentId())
	requestStringBuilder.Append("conversationData", c.GetConversationData())

	return requestStringBuilder.GetRequestString()
}
