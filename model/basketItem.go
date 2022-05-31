package model

import "github.com/umtaktpe/iyzipay-go/utils"

type BasketItem struct {
	id               string
	price            string
	name             string
	category1        string
	category2        string
	itemType         string
	subMerchantKey   string
	subMerchantPrice string
}

func NewBasketItem() *BasketItem {
	return &BasketItem{}
}

func (b *BasketItem) GetId() string {
	return b.id
}

func (b *BasketItem) SetId(id string) {
	b.id = id
}

func (b *BasketItem) GetPrice() string {
	return b.price
}

func (b *BasketItem) SetPrice(price string) {
	b.price = price
}

func (b *BasketItem) GetName() string {
	return b.name
}

func (b *BasketItem) SetName(name string) {
	b.name = name
}

func (b *BasketItem) GetCategory1() string {
	return b.category1
}

func (b *BasketItem) SetCategory1(category1 string) {
	b.category1 = category1
}

func (b *BasketItem) GetCategory2() string {
	return b.category2
}

func (b *BasketItem) SetCategory2(category2 string) {
	b.category2 = category2
}

func (b *BasketItem) GetItemType() string {
	return b.itemType
}

func (b *BasketItem) SetItemType(itemType string) {
	b.itemType = itemType
}

func (b *BasketItem) GetSubMerchantKey() string {
	return b.subMerchantKey
}

func (b *BasketItem) SetSubMerchantKey(subMerchantKey string) {
	b.subMerchantKey = subMerchantKey
}

func (b *BasketItem) GetSubMerchantPrice() string {
	return b.subMerchantPrice
}

func (b *BasketItem) SetSubMerchantPrice(subMerchantPrice string) {
	b.subMerchantPrice = subMerchantPrice
}

func (b *BasketItem) GetJsonObject() map[string]interface{} {
	jsonBuilder := utils.NewJsonBuilder()
	jsonBuilder.Add("id", b.GetId())
	jsonBuilder.AddPrice("price", b.GetPrice())
	jsonBuilder.Add("name", b.GetName())
	jsonBuilder.Add("category1", b.GetCategory1())
	jsonBuilder.Add("category2", b.GetCategory2())
	jsonBuilder.Add("itemType", b.GetItemType())
	jsonBuilder.Add("subMerchantKey", b.GetSubMerchantKey())
	jsonBuilder.AddPrice("subMerchantPrice", b.GetSubMerchantPrice())

	return jsonBuilder.GetObject()
}

func (b *BasketItem) ToPKIRequestString() string {
	requestStringBuilder := utils.NewRequestStringBuilder()
	requestStringBuilder.Append("id", b.GetId())
	requestStringBuilder.AppendPrice("price", b.GetPrice())
	requestStringBuilder.Append("name", b.GetName())
	requestStringBuilder.Append("category1", b.GetCategory1())
	requestStringBuilder.Append("category2", b.GetCategory2())
	requestStringBuilder.Append("itemType", b.GetItemType())
	requestStringBuilder.Append("subMerchantKey", b.GetSubMerchantKey())
	requestStringBuilder.AppendPrice("subMerchantPrice", b.GetSubMerchantPrice())

	return requestStringBuilder.GetRequestString()
}
