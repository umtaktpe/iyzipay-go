package model

import "github.com/umtaktpe/iyzipay-go/utils"

type Buyer struct {
	id                  string
	name                string
	surname             string
	identityNumber      string
	email               string
	gsmNumber           string
	registrationDate    string
	lastLoginDate       string
	registrationAddress string
	city                string
	country             string
	zipcode             string
	ip                  string
}

func NewBuyer() *Buyer {
	return &Buyer{}
}

func (b *Buyer) GetId() string {
	return b.id
}

func (b *Buyer) SetId(id string) {
	b.id = id
}

func (b *Buyer) GetName() string {
	return b.name
}

func (b *Buyer) SetName(name string) {
	b.name = name
}

func (b *Buyer) GetSurname() string {
	return b.surname
}

func (b *Buyer) SetSurname(surname string) {
	b.surname = surname
}

func (b *Buyer) GetIdentityNumber() string {
	return b.identityNumber
}

func (b *Buyer) SetIdentityNumber(identityNumber string) {
	b.identityNumber = identityNumber
}

func (b *Buyer) GetEmail() string {
	return b.email
}

func (b *Buyer) SetEmail(email string) {
	b.email = email
}

func (b *Buyer) GetGsmNumber() string {
	return b.gsmNumber
}

func (b *Buyer) SetGsmNumber(gsmNumber string) {
	b.gsmNumber = gsmNumber
}

func (b *Buyer) GetRegistrationDate() string {
	return b.registrationDate
}

func (b *Buyer) SetRegistrationDate(registrationDate string) {
	b.registrationDate = registrationDate
}

func (b *Buyer) GetLastLoginDate() string {
	return b.lastLoginDate
}

func (b *Buyer) SetLastLoginDate(lastLoginDate string) {
	b.lastLoginDate = lastLoginDate
}

func (b *Buyer) GetRegistrationAddress() string {
	return b.registrationAddress
}

func (b *Buyer) SetRegistrationAddress(registrationAddress string) {
	b.registrationAddress = registrationAddress
}

func (b *Buyer) GetCity() string {
	return b.city
}

func (b *Buyer) SetCity(city string) {
	b.city = city
}

func (b *Buyer) GetCountry() string {
	return b.country
}

func (b *Buyer) SetCountry(country string) {
	b.country = country
}

func (b *Buyer) GetZipcode() string {
	return b.zipcode
}

func (b *Buyer) SetZipcode(zipcode string) {
	b.zipcode = zipcode
}

func (b *Buyer) GetIp() string {
	return b.ip
}

func (b *Buyer) SetIp(ip string) {
	b.ip = ip
}

func (b *Buyer) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("id", b.GetId())
	jsonBuilder.Add("name", b.GetName())
	jsonBuilder.Add("surname", b.GetSurname())
	jsonBuilder.Add("identityNumber", b.GetIdentityNumber())
	jsonBuilder.Add("email", b.GetEmail())
	jsonBuilder.Add("gsmNumber", b.GetGsmNumber())
	jsonBuilder.Add("registrationDate", b.GetRegistrationDate())
	jsonBuilder.Add("lastLoginDate", b.GetLastLoginDate())
	jsonBuilder.Add("registrationAddress", b.GetRegistrationAddress())
	jsonBuilder.Add("city", b.GetCity())
	jsonBuilder.Add("country", b.GetCountry())
	jsonBuilder.Add("zipCode", b.GetZipcode())
	jsonBuilder.Add("ip", b.GetIp())

	return jsonBuilder.GetObject()
}

func (b *Buyer) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("id", b.GetId())
	requestStringBuilder.Append("name", b.GetName())
	requestStringBuilder.Append("surname", b.GetSurname())
	requestStringBuilder.Append("identityNumber", b.GetIdentityNumber())
	requestStringBuilder.Append("email", b.GetEmail())
	requestStringBuilder.Append("gsmNumber", b.GetGsmNumber())
	requestStringBuilder.Append("registrationDate", b.GetRegistrationDate())
	requestStringBuilder.Append("lastLoginDate", b.GetLastLoginDate())
	requestStringBuilder.Append("registrationAddress", b.GetRegistrationAddress())
	requestStringBuilder.Append("city", b.GetCity())
	requestStringBuilder.Append("country", b.GetCountry())
	requestStringBuilder.Append("zipCode", b.GetZipcode())
	requestStringBuilder.Append("ip", b.GetIp())

	return requestStringBuilder.GetRequestString()
}
