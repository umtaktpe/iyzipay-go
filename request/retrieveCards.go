package request

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

type RetrieveCards struct {
	model.BaseModel
	cardUserKey string
}

func NewRetrieveCards() *RetrieveCards {
	return &RetrieveCards{}
}

func (r *RetrieveCards) GetCardUserKey() string {
	return r.cardUserKey
}

func (r *RetrieveCards) SetCardUserKey(cardUserKey string) {
	r.cardUserKey = cardUserKey
}

func (r *RetrieveCards) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("cardUserKey", r.GetCardUserKey())

	return jsonBuilder.GetObject()
}

func (r *RetrieveCards) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("cardUserKey", r.GetCardUserKey())

	return requestStringBuilder.GetRequestString()
}
