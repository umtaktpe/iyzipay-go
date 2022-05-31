package request

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

type DeleteCard struct {
	model.BaseModel
	cardUserKey string
	cardToken   string
}

func NewDeleteCard() *DeleteCard {
	return &DeleteCard{}
}

func (d *DeleteCard) GetCardUserKey() string {
	return d.cardUserKey
}

func (d *DeleteCard) SetCardUserKey(cardUserKey string) {
	d.cardUserKey = cardUserKey
}

func (d *DeleteCard) GetCardToken() string {
	return d.cardToken
}

func (d *DeleteCard) SetCardToken(cardToken string) {
	d.cardToken = cardToken
}

func (d *DeleteCard) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("cardUserKey", d.GetCardUserKey())
	jsonBuilder.Add("cardToken", d.GetCardToken())

	return jsonBuilder.GetObject()
}

func (d *DeleteCard) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("cardUserKey", d.GetCardUserKey())
	requestStringBuilder.Append("cardToken", d.GetCardToken())

	return requestStringBuilder.GetRequestString()
}
