package model

import "github.com/umtaktpe/iyzipay-go/utils"

type PaymentCard struct {
	cardHolderName string
	cardNumber     string
	expireYear     string
	expireMonth    string
	cvc            string
	registerCard   string
	cardAlias      string
	cardToken      string
	cardUserKey    string
}

func NewPaymentCard() *PaymentCard {
	return &PaymentCard{}
}

func (p *PaymentCard) GetCardHolderName() string {
	return p.cardHolderName
}

func (p *PaymentCard) SetCardHolderName(cardHolderName string) {
	p.cardHolderName = cardHolderName
}

func (p *PaymentCard) GetCardNumber() string {
	return p.cardNumber
}

func (p *PaymentCard) SetCardNumber(cardNumber string) {
	p.cardNumber = cardNumber
}

func (p *PaymentCard) GetExpireYear() string {
	return p.expireYear
}

func (p *PaymentCard) SetExpireYear(expireYear string) {
	p.expireYear = expireYear
}

func (p *PaymentCard) GetExpireMonth() string {
	return p.expireMonth
}

func (p *PaymentCard) SetExpireMonth(expireMonth string) {
	p.expireMonth = expireMonth
}

func (p *PaymentCard) GetCvc() string {
	return p.cvc
}

func (p *PaymentCard) SetCvc(cvc string) {
	p.cvc = cvc
}

func (p *PaymentCard) GetRegisterCard() string {
	return p.registerCard
}

func (p *PaymentCard) SetRegisterCard(registerCard string) {
	p.registerCard = registerCard
}

func (p *PaymentCard) GetCardAlias() string {
	return p.cardAlias
}

func (p *PaymentCard) SetCardAlias(cardAlias string) {
	p.cardAlias = cardAlias
}

func (p *PaymentCard) GetCardToken() string {
	return p.cardToken
}

func (p *PaymentCard) SetCardToken(cardToken string) {
	p.cardToken = cardToken
}

func (p *PaymentCard) GetCardUserKey() string {
	return p.cardUserKey
}

func (p *PaymentCard) SetCardUserKey(cardUserKey string) {
	p.cardUserKey = cardUserKey
}

func (p *PaymentCard) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("cardHolderName", p.GetCardHolderName())
	jsonBuilder.Add("cardNumber", p.GetCardNumber())
	jsonBuilder.Add("expireYear", p.GetExpireYear())
	jsonBuilder.Add("expireMonth", p.GetExpireMonth())
	jsonBuilder.Add("cvc", p.GetCvc())
	jsonBuilder.Add("registerCard", p.GetRegisterCard())
	jsonBuilder.Add("cardAlias", p.GetCardAlias())
	jsonBuilder.Add("cardToken", p.GetCardToken())
	jsonBuilder.Add("cardUserKey", p.GetCardUserKey())
	return jsonBuilder.GetObject()
}

func (p *PaymentCard) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("cardHolderName", p.GetCardHolderName())
	requestStringBuilder.Append("cardNumber", p.GetCardNumber())
	requestStringBuilder.Append("expireYear", p.GetExpireYear())
	requestStringBuilder.Append("expireMonth", p.GetExpireMonth())
	requestStringBuilder.Append("cvc", p.GetCvc())
	requestStringBuilder.Append("registerCard", p.GetRegisterCard())
	requestStringBuilder.Append("cardAlias", p.GetCardAlias())
	requestStringBuilder.Append("cardToken", p.GetCardToken())
	requestStringBuilder.Append("cardUserKey", p.GetCardUserKey())

	return requestStringBuilder.GetRequestString()
}
