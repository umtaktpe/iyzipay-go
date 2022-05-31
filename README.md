# iyzipay-go

# Installation
```
go get github.com/umtaktpe/iyzipay-go
```

# Usage
```go
options := model.NewOptions()
options.SetApiKey("your api key")
options.SetApiSecret("your secret key")
options.SetBaseUrl("https://sandbox-api.iyzipay.com")

request := request.NewCreatePayment()
request.SetLocale(utils.TR)
request.SetConversationId("123456789")
request.SetPrice("1")
request.SetPaidPrice("1.2")
request.SetCurrency(utils.TL)
request.SetInstallment("1")
request.SetBasketId("B67832")
request.SetPaymentChannel(utils.WEB)
request.SetPaymentGroup(utils.PRODUCT)

paymentCard := model.NewPaymentCard()
paymentCard.SetCardHolderName("John Doe")
paymentCard.SetCardNumber("5528790000000008")
paymentCard.SetExpireMonth("12")
paymentCard.SetExpireYear("2030")
paymentCard.SetCvc("123")
paymentCard.SetRegisterCard("0")
request.SetPaymentCard(paymentCard)

buyer := model.NewBuyer()
buyer.SetId("BY789")
buyer.SetName("John")
buyer.SetSurname("Doe")
buyer.SetGsmNumber("+905350000000")
buyer.SetEmail("email@email.com")
buyer.SetIdentityNumber("74300864791")
buyer.SetLastLoginDate("2015-10-05 12:43:35")
buyer.SetRegistrationDate("2013-04-21 15:12:09")
buyer.SetRegistrationAddress("Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1")
buyer.SetIp("85.34.78.112")
buyer.SetCity("Istanbul")
buyer.SetCountry("Turkey")
buyer.SetZipcode("34732")
request.SetBuyer(buyer)

shippingAddress := model.NewAddress()
shippingAddress.SetContactName("Jane Doe")
shippingAddress.SetCity("Istanbul")
shippingAddress.SetCountry("Turkey")
shippingAddress.SetAddress("Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1")
shippingAddress.SetZipCode("34742")
request.SetShippingAddress(shippingAddress)

billingAddress := model.NewAddress()
billingAddress.SetContactName("Jane Doe")
billingAddress.SetCity("Istanbul")
billingAddress.SetCountry("Turkey")
billingAddress.SetAddress("Nidakule Göztepe, Merdivenköy Mah. Bora Sok. No:1")
billingAddress.SetZipCode("34742")
request.SetBillingAddress(billingAddress)

var basketItems []*model.BasketItem
firstBasketItem := model.NewBasketItem()
firstBasketItem.SetId("BI101")
firstBasketItem.SetName("Binocular")
firstBasketItem.SetCategory1("Collectibles")
firstBasketItem.SetCategory2("Accessories")
firstBasketItem.SetItemType(utils.PHYSICAL)
firstBasketItem.SetPrice("0.3")
basketItems = append(basketItems, firstBasketItem)

secondBasketItem := model.NewBasketItem()
secondBasketItem.SetId("BI102")
secondBasketItem.SetName("Game code")
secondBasketItem.SetCategory1("Game")
secondBasketItem.SetCategory2("Online Game Items")
secondBasketItem.SetItemType(utils.VIRTUAL)
secondBasketItem.SetPrice("0.5")
basketItems = append(basketItems, secondBasketItem)

thirdBasketItem := model.NewBasketItem()
thirdBasketItem.SetId("BI103")
thirdBasketItem.SetName("Usb")
thirdBasketItem.SetCategory1("Electronics")
thirdBasketItem.SetCategory2("Usb / Cable")
thirdBasketItem.SetItemType(utils.PHYSICAL)
thirdBasketItem.SetPrice("0.2")
basketItems = append(basketItems, thirdBasketItem)
request.SetBasketItems(basketItems)

response, err := api.CreatePayment(request, options)
if err != nil {
    panic(err)
}

fmt.Println(response)
```

# Development
### Mock test cards

Test cards that can be used to simulate a *successful* payment:

Card Number      | Bank                       | Card Type
-----------      | ----                       | ---------
5890040000000016 | Akbank                     | Master Card (Debit)  
5526080000000006 | Akbank                     | Master Card (Credit)  
4766620000000001 | Denizbank                  | Visa (Debit)  
4603450000000000 | Denizbank                  | Visa (Credit)
4729150000000005 | Denizbank Bonus            | Visa (Credit)  
4987490000000002 | Finansbank                 | Visa (Debit)  
5311570000000005 | Finansbank                 | Master Card (Credit)  
9792020000000001 | Finansbank                 | Troy (Debit)  
9792030000000000 | Finansbank                 | Troy (Credit)  
5170410000000004 | Garanti Bankası            | Master Card (Debit)  
5400360000000003 | Garanti Bankası            | Master Card (Credit)  
374427000000003  | Garanti Bankası            | American Express  
4475050000000003 | Halkbank                   | Visa (Debit)  
5528790000000008 | Halkbank                   | Master Card (Credit)  
4059030000000009 | HSBC Bank                  | Visa (Debit)  
5504720000000003 | HSBC Bank                  | Master Card (Credit)  
5892830000000000 | Türkiye İş Bankası         | Master Card (Debit)  
4543590000000006 | Türkiye İş Bankası         | Visa (Credit)  
4910050000000006 | Vakıfbank                  | Visa (Debit)  
4157920000000002 | Vakıfbank                  | Visa (Credit)  
5168880000000002 | Yapı ve Kredi Bankası      | Master Card (Debit)  
5451030000000000 | Yapı ve Kredi Bankası      | Master Card (Credit)  

*Cross border* test cards:

Card Number      | Country
-----------      | -------
4054180000000007 | Non-Turkish (Debit)
5400010000000004 | Non-Turkish (Credit)  
6221060000000004 | Iran  

Test cards to get specific *error* codes:

Card Number       | Description
-----------       | -----------
5406670000000009  | Success but cannot be cancelled, refund or post auth
4111111111111129  | Not sufficient funds
4129111111111111  | Do not honour
4128111111111112  | Invalid transaction
4127111111111113  | Lost card
4126111111111114  | Stolen card
4125111111111115  | Expired card
4124111111111116  | Invalid cvc2
4123111111111117  | Not permitted to card holder
4122111111111118  | Not permitted to terminal
4121111111111119  | Fraud suspect
4120111111111110  | Pickup card
4130111111111118  | General error
4131111111111117  | Success but mdStatus is 0
4141111111111115  | Success but mdStatus is 4
4151111111111112  | 3dsecure initialize failed

### Mock APM Accounts

Mock APM Accounts that can be used to simulate a payment with alternative payment method:

Account Holder Name     | Description
-------------------     | -----------
success                 | Succeeded payment after succeeded initialize
fail-after-init         | Failed payment after succeeded initialize
error                   | Failed initialize