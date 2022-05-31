package request

import (
	model "github.com/umtaktpe/iyzipay-go/model"
	utils "github.com/umtaktpe/iyzipay-go/utils"
)

type CreatePayment struct {
	model.BaseModel
	price           string
	paidPrice       string
	installment     string
	paymentChannel  string
	basketId        string
	paymentGroup    string
	paymentCard     *model.PaymentCard
	buyer           *model.Buyer
	shippingAddress *model.Address
	billingAddress  *model.Address
	basketItems     []*model.BasketItem
	paymentSource   string
	currency        string
	posOrderId      string
	connectorName   string
	callbackUrl     string
}

func NewCreatePayment() *CreatePayment {
	return &CreatePayment{
		installment: utils.SINGLE_INSTALLMENT,
	}
}

func (p *CreatePayment) GetPrice() string {
	return p.price
}

func (p *CreatePayment) SetPrice(price string) {
	p.price = price
}

func (p *CreatePayment) GetPaidPrice() string {
	return p.paidPrice
}

func (p *CreatePayment) SetPaidPrice(paidPrice string) {
	p.paidPrice = paidPrice
}

func (p *CreatePayment) GetInstallment() string {
	return p.installment
}

func (p *CreatePayment) SetInstallment(installment string) {
	p.installment = installment
}

func (p *CreatePayment) GetPaymentChannel() string {
	return p.paymentChannel
}

func (p *CreatePayment) SetPaymentChannel(paymentChannel string) {
	p.paymentChannel = paymentChannel
}

func (p *CreatePayment) GetBasketId() string {
	return p.basketId
}

func (p *CreatePayment) SetBasketId(basketId string) {
	p.basketId = basketId
}

func (p *CreatePayment) GetPaymentGroup() string {
	return p.paymentGroup
}

func (p *CreatePayment) SetPaymentGroup(paymentGroup string) {
	p.paymentGroup = paymentGroup
}

func (p *CreatePayment) GetPaymentCard() *model.PaymentCard {
	return p.paymentCard
}

func (p *CreatePayment) SetPaymentCard(paymentCard *model.PaymentCard) {
	p.paymentCard = paymentCard
}

func (p *CreatePayment) GetBuyer() *model.Buyer {
	return p.buyer
}

func (p *CreatePayment) SetBuyer(buyer *model.Buyer) {
	p.buyer = buyer
}

func (p *CreatePayment) GetShippingAddress() *model.Address {
	return p.shippingAddress
}

func (p *CreatePayment) SetShippingAddress(shippingAddress *model.Address) {
	p.shippingAddress = shippingAddress
}

func (p *CreatePayment) GetBillingAddress() *model.Address {
	return p.billingAddress
}

func (p *CreatePayment) SetBillingAddress(billingAddress *model.Address) {
	p.billingAddress = billingAddress
}

func (p *CreatePayment) GetBasketItems() []*model.BasketItem {
	return p.basketItems
}

func (p *CreatePayment) SetBasketItems(basketItems []*model.BasketItem) {
	p.basketItems = basketItems
}

func (p *CreatePayment) GetPaymentSource() string {
	return p.paymentSource
}

func (p *CreatePayment) SetPaymentSource(paymentSource string) {
	p.paymentSource = paymentSource
}

func (p *CreatePayment) GetCurrency() string {
	return p.currency
}

func (p *CreatePayment) SetCurrency(currency string) {
	p.currency = currency
}

func (p *CreatePayment) GetPosOrderId() string {
	return p.posOrderId
}

func (p *CreatePayment) SetPosOrderId(posOrderId string) {
	p.posOrderId = posOrderId
}

func (p *CreatePayment) GetConnectorName() string {
	return p.connectorName
}

func (p *CreatePayment) SetConnectorName(connectorName string) {
	p.connectorName = connectorName
}

func (p *CreatePayment) GetCallbackUrl() string {
	return p.callbackUrl
}

func (p *CreatePayment) SetCallbackUrl(callbackUrl string) {
	p.callbackUrl = callbackUrl
}

func (p *CreatePayment) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("locale", p.GetLocale())
	jsonBuilder.Add("conversationId", p.GetConversationId())
	jsonBuilder.AddPrice("price", p.GetPrice())
	jsonBuilder.AddPrice("paidPrice", p.GetPaidPrice())
	jsonBuilder.Add("installment", p.GetInstallment())
	jsonBuilder.Add("paymentChannel", p.GetPaymentChannel())
	jsonBuilder.Add("basketId", p.GetBasketId())
	jsonBuilder.Add("paymentGroup", p.GetPaymentGroup())
	jsonBuilder.AddModel("paymentCard", p.GetPaymentCard())
	jsonBuilder.AddModel("buyer", p.GetBuyer())
	jsonBuilder.AddModel("shippingAddress", p.GetShippingAddress())
	jsonBuilder.AddModel("billingAddress", p.GetBillingAddress())

	var jsonArray []interface{}
	for _, v := range p.GetBasketItems() {
		jsonArray = append(jsonArray, v.GetJsonObject())
	}
	jsonBuilder.AddArray("basketItems", jsonArray)

	jsonBuilder.Add("paymentSource", p.GetPaymentSource())
	jsonBuilder.Add("currency", p.GetCurrency())
	jsonBuilder.Add("posOrderId", p.GetPosOrderId())
	jsonBuilder.Add("connectorName", p.GetConnectorName())
	jsonBuilder.Add("callbackUrl", p.GetCallbackUrl())

	return jsonBuilder.GetObject()
}

func (p *CreatePayment) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("locale", p.GetLocale())
	requestStringBuilder.Append("conversationId", p.GetConversationId())
	requestStringBuilder.AppendPrice("price", p.GetPrice())
	requestStringBuilder.AppendPrice("paidPrice", p.GetPaidPrice())
	requestStringBuilder.Append("installment", p.GetInstallment())
	requestStringBuilder.Append("paymentChannel", p.GetPaymentChannel())
	requestStringBuilder.Append("basketId", p.GetBasketId())
	requestStringBuilder.Append("paymentGroup", p.GetPaymentGroup())
	requestStringBuilder.AppendModel("paymentCard", p.GetPaymentCard())
	requestStringBuilder.AppendModel("buyer", p.GetBuyer())
	requestStringBuilder.AppendModel("shippingAddress", p.GetShippingAddress())
	requestStringBuilder.AppendModel("billingAddress", p.GetBillingAddress())

	var pkiArray []string
	for _, v := range p.GetBasketItems() {
		pkiArray = append(pkiArray, v.ToPKIRequestString())
	}
	requestStringBuilder.AppendArray("basketItems", pkiArray)

	requestStringBuilder.Append("paymentSource", p.GetPaymentSource())
	requestStringBuilder.Append("currency", p.GetCurrency())
	requestStringBuilder.Append("posOrderId", p.GetPosOrderId())
	requestStringBuilder.Append("connectorName", p.GetConnectorName())
	requestStringBuilder.Append("callbackUrl", p.GetCallbackUrl())

	return requestStringBuilder.GetRequestString()
}
