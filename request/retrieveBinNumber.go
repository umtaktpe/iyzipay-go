package request

import (
	"github.com/umtaktpe/iyzipay-go/model"
	"github.com/umtaktpe/iyzipay-go/utils"
)

type RetrieveBinNumber struct {
	model.BaseModel
	binNumber string
}

func NewRetrieveBinNumber() *RetrieveBinNumber {
	return &RetrieveBinNumber{}
}

func (r *RetrieveBinNumber) GetBinNumber() string {
	return r.binNumber
}

func (r *RetrieveBinNumber) SetBinNumber(binNumber string) {
	r.binNumber = binNumber
}

func (r *RetrieveBinNumber) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("binNumber", r.GetBinNumber())

	return jsonBuilder.GetObject()
}

func (r *RetrieveBinNumber) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("binNumber", r.GetBinNumber())

	return requestStringBuilder.GetRequestString()
}
