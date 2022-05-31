package request

import (
	model "github.com/umtaktpe/iyzipay-go/model"
	utils "github.com/umtaktpe/iyzipay-go/utils"
)

type CreateCard struct {
	model.BaseModel
	externalId  string
	email       string
	cardUserKey string
	card        *model.CardInformation
}

func NewCreateCard() *CreateCard {
	return &CreateCard{}
}

func (c *CreateCard) GetExternalId() string {
	return c.externalId
}

func (c *CreateCard) SetExternalId(externalId string) {
	c.externalId = externalId
}

func (c *CreateCard) GetEmail() string {
	return c.email
}

func (c *CreateCard) SetEmail(email string) {
	c.email = email
}

func (c *CreateCard) GetCardUserKey() string {
	return c.cardUserKey
}

func (c *CreateCard) SetCardUserKey(cardUserKey string) {
	c.cardUserKey = cardUserKey
}

func (c *CreateCard) GetCard() *model.CardInformation {
	return c.card
}

func (c *CreateCard) SetCard(card *model.CardInformation) {
	c.card = card
}

func (c *CreateCard) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("locale", c.GetLocale())
	jsonBuilder.Add("conversationId", c.GetConversationId())
	jsonBuilder.Add("externalId", c.GetExternalId())
	jsonBuilder.Add("email", c.GetEmail())
	jsonBuilder.Add("cardUserKey", c.GetCardUserKey())
	jsonBuilder.AddModel("card", c.GetCard())

	return jsonBuilder.GetObject()
}

func (c *CreateCard) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("locale", c.GetLocale())
	requestStringBuilder.Append("conversationId", c.GetConversationId())
	requestStringBuilder.Append("externalId", c.GetExternalId())
	requestStringBuilder.Append("email", c.GetEmail())
	requestStringBuilder.Append("cardUserKey", c.GetCardUserKey())
	requestStringBuilder.AppendModel("card", c.GetCard())

	return requestStringBuilder.GetRequestString()
}
