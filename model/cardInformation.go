package model

import "github.com/umtaktpe/iyzipay-go/utils"

type CardInformation struct {
	cardAlias      string
	cardNumber     string
	expireYear     string
	expireMonth    string
	cardHolderName string
}

func NewCardInformation() *CardInformation {
	return &CardInformation{}
}

func (c *CardInformation) GetCardAlias() string {
	return c.cardAlias
}

func (c *CardInformation) SetCardAlias(cardAlias string) {
	c.cardAlias = cardAlias
}

func (c *CardInformation) GetCardNumber() string {
	return c.cardNumber
}

func (c *CardInformation) SetCardNumber(cardNumber string) {
	c.cardNumber = cardNumber
}

func (c *CardInformation) GetExpireYear() string {
	return c.expireYear
}

func (c *CardInformation) SetExpireYear(expireYear string) {
	c.expireYear = expireYear
}

func (c *CardInformation) GetExpireMonth() string {
	return c.expireMonth
}

func (c *CardInformation) SetExpireMonth(expireMonth string) {
	c.expireMonth = expireMonth
}

func (c *CardInformation) GetCardHolderName() string {
	return c.cardHolderName
}

func (c *CardInformation) SetCardHolderName(cardHolderName string) {
	c.cardHolderName = cardHolderName
}

func (c *CardInformation) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("cardAlias", c.GetCardAlias())
	jsonBuilder.Add("cardNumber", c.GetCardNumber())
	jsonBuilder.Add("expireYear", c.GetExpireYear())
	jsonBuilder.Add("expireMonth", c.GetExpireMonth())
	jsonBuilder.Add("cardHolderName", c.GetCardHolderName())

	return jsonBuilder.GetObject()
}

func (c *CardInformation) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("cardAlias", c.GetCardAlias())
	requestStringBuilder.Append("cardNumber", c.GetCardNumber())
	requestStringBuilder.Append("expireYear", c.GetExpireYear())
	requestStringBuilder.Append("expireMonth", c.GetExpireMonth())
	requestStringBuilder.Append("cardHolderName", c.GetCardHolderName())

	return requestStringBuilder.GetRequestString()
}
