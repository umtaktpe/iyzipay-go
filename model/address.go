package model

import "github.com/umtaktpe/iyzipay-go/utils"

type Address struct {
	contactName string
	city        string
	country     string
	address     string
	zipCode     string
}

func NewAddress() *Address {
	return &Address{}
}

func (a *Address) GetAddress() string {
	return a.address
}

func (a *Address) SetAddress(address string) {
	a.address = address
}

func (a *Address) GetZipCode() string {
	return a.zipCode
}

func (a *Address) SetZipCode(zipCode string) {
	a.zipCode = zipCode
}

func (a *Address) GetContactName() string {
	return a.contactName
}

func (a *Address) SetContactName(contactName string) {
	a.contactName = contactName
}

func (a *Address) GetCity() string {
	return a.city
}

func (a *Address) SetCity(city string) {
	a.city = city
}

func (a *Address) GetCountry() string {
	return a.country
}

func (a *Address) SetCountry(country string) {
	a.country = country
}

func (a *Address) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("address", a.GetAddress())
	jsonBuilder.Add("zipCode", a.GetZipCode())
	jsonBuilder.Add("contactName", a.GetContactName())
	jsonBuilder.Add("city", a.GetCity())
	jsonBuilder.Add("country", a.GetCountry())

	return jsonBuilder.GetObject()
}

func (a *Address) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("address", a.GetAddress())
	requestStringBuilder.Append("zipCode", a.GetZipCode())
	requestStringBuilder.Append("contactName", a.GetContactName())
	requestStringBuilder.Append("city", a.GetCity())
	requestStringBuilder.Append("country", a.GetCountry())

	return requestStringBuilder.GetRequestString()
}
