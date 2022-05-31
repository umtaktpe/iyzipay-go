package api

// General response with error messages
type baseResponse struct {
	Status       string `json:"status"`
	Locale       string `json:"locale"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	SystemTime   int64  `json:"systemTime"`
}

// Responses to card transactions
type retriveCards struct {
	baseResponse
	CardUserKey string `json:"cardUserKey"`
	CardDetails []struct {
		CardToken       string `json:"cardToken"`
		CardAlias       string `json:"cardAlias"`
		BinNumber       string `json:"binNumber"`
		LastFourDigits  string `json:"lastFourDigits"`
		CardType        string `json:"cardType"`
		CardAssociation string `json:"cardAssociation"`
		CardFamily      string `json:"cardFamily"`
		CardBankCode    int    `json:"cardBankCode"`
		CardBankName    string `json:"cardBankName"`
	} `json:"cardDetails"`
}

type createCard struct {
	ConversationID  string `json:"conversationId"`
	ExternalID      string `json:"externalId"`
	Email           string `json:"email"`
	CardUserKey     string `json:"cardUserKey"`
	CardToken       string `json:"cardToken"`
	CardAlias       string `json:"cardAlias"`
	BinNumber       string `json:"binNumber"`
	LastFourDigits  string `json:"lastFourDigits"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	CardBankCode    int    `json:"cardBankCode"`
	CardBankName    string `json:"cardBankName"`
}

// Responses to installment check
type installmentInfo struct {
	baseResponse
	InstallmentDetails []struct {
		BinNumber         string  `json:"binNumber"`
		Price             float64 `json:"price"`
		CardType          string  `json:"cardType"`
		CardAssociation   string  `json:"cardAssociation"`
		CardFamilyName    string  `json:"cardFamilyName"`
		Force3Ds          int     `json:"force3ds"`
		BankCode          int     `json:"bankCode"`
		BankName          string  `json:"bankName"`
		ForceCvc          int     `json:"forceCvc"`
		Commercial        int     `json:"commercial"`
		InstallmentPrices []struct {
			InstallmentPrice  float64 `json:"installmentPrice"`
			TotalPrice        float64 `json:"totalPrice"`
			InstallmentNumber int     `json:"installmentNumber"`
		} `json:"installmentPrices"`
	} `json:"installmentDetails"`
}

type binNumber struct {
	baseResponse
	BinNumber       string `json:"binNumber"`
	CardType        string `json:"cardType"`
	CardAssociation string `json:"cardAssociation"`
	CardFamily      string `json:"cardFamily"`
	BankName        string `json:"bankName"`
	BankCode        int    `json:"bankCode"`
	Commercial      int    `json:"commercial"`
}

// Responses to payment transactions
type retrievePayment struct {
	Status                       string  `json:"status"`
	Locale                       string  `json:"locale"`
	SystemTime                   int64   `json:"systemTime"`
	Price                        float64 `json:"price"`
	PaidPrice                    float64 `json:"paidPrice"`
	Installment                  int     `json:"installment"`
	PaymentID                    string  `json:"paymentId"`
	FraudStatus                  int     `json:"fraudStatus"`
	MerchantCommissionRate       float64 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	BinNumber                    string  `json:"binNumber"`
	LastFourDigits               string  `json:"lastFourDigits"`
	BasketID                     string  `json:"basketId"`
	Currency                     string  `json:"currency"`
	ItemTransactions             []struct {
		ItemID                        string  `json:"itemId"`
		PaymentTransactionID          string  `json:"paymentTransactionId"`
		TransactionStatus             int     `json:"transactionStatus"`
		Price                         float64 `json:"price"`
		PaidPrice                     float64 `json:"paidPrice"`
		MerchantCommissionRate        float64 `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRate                  float64 `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		SubMerchantPrice              float64 `json:"subMerchantPrice"`
		SubMerchantPayoutRate         float64 `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		ConvertedPayout               struct {
			PaidPrice                     float64 `json:"paidPrice"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
			SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			IyziConversionRate            float64 `json:"iyziConversionRate"`
			IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
			Currency                      string  `json:"currency"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
	AuthCode      string `json:"authCode"`
	Phase         string `json:"phase"`
	HostReference string `json:"hostReference"`
	PaymentStatus string `json:"paymentStatus"`
}

type createPayment struct {
	baseResponse
	ConversationID               string  `json:"conversationId"`
	Price                        float64 `json:"price"`
	PaidPrice                    float64 `json:"paidPrice"`
	Installment                  int     `json:"installment"`
	PaymentID                    string  `json:"paymentId"`
	FraudStatus                  int     `json:"fraudStatus"`
	MerchantCommissionRate       float64 `json:"merchantCommissionRate"`
	MerchantCommissionRateAmount float64 `json:"merchantCommissionRateAmount"`
	IyziCommissionRateAmount     float64 `json:"iyziCommissionRateAmount"`
	IyziCommissionFee            float64 `json:"iyziCommissionFee"`
	CardType                     string  `json:"cardType"`
	CardAssociation              string  `json:"cardAssociation"`
	CardFamily                   string  `json:"cardFamily"`
	BinNumber                    string  `json:"binNumber"`
	LastFourDigits               string  `json:"lastFourDigits"`
	BasketID                     string  `json:"basketId"`
	Currency                     string  `json:"currency"`
	ItemTransactions             []struct {
		ItemID                        string  `json:"itemId"`
		PaymentTransactionID          string  `json:"paymentTransactionId"`
		TransactionStatus             int     `json:"transactionStatus"`
		Price                         float64 `json:"price"`
		PaidPrice                     float64 `json:"paidPrice"`
		MerchantCommissionRate        float64 `json:"merchantCommissionRate"`
		MerchantCommissionRateAmount  float64 `json:"merchantCommissionRateAmount"`
		IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
		IyziCommissionFee             float64 `json:"iyziCommissionFee"`
		BlockageRate                  float64 `json:"blockageRate"`
		BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
		BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
		BlockageResolvedDate          string  `json:"blockageResolvedDate"`
		SubMerchantPrice              float64 `json:"subMerchantPrice"`
		SubMerchantPayoutRate         float64 `json:"subMerchantPayoutRate"`
		SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
		MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
		ConvertedPayout               struct {
			PaidPrice                     float64 `json:"paidPrice"`
			IyziCommissionRateAmount      float64 `json:"iyziCommissionRateAmount"`
			IyziCommissionFee             float64 `json:"iyziCommissionFee"`
			BlockageRateAmountMerchant    float64 `json:"blockageRateAmountMerchant"`
			BlockageRateAmountSubMerchant float64 `json:"blockageRateAmountSubMerchant"`
			SubMerchantPayoutAmount       float64 `json:"subMerchantPayoutAmount"`
			MerchantPayoutAmount          float64 `json:"merchantPayoutAmount"`
			IyziConversionRate            float64 `json:"iyziConversionRate"`
			IyziConversionRateAmount      float64 `json:"iyziConversionRateAmount"`
			Currency                      string  `json:"currency"`
		} `json:"convertedPayout"`
	} `json:"itemTransactions"`
	AuthCode      string `json:"authCode"`
	Phase         string `json:"phase"`
	HostReference string `json:"hostReference"`
}

type initializeThreedsPayment struct {
	baseResponse
	ConversationID     string `json:"conversationId"`
	ThreeDSHTMLContent string `json:"threeDSHtmlContent"`
}

type cancelPayment struct {
	baseResponse
	PaymentID     string  `json:"paymentId"`
	Price         float64 `json:"price"`
	Currency      string  `json:"currency"`
	AuthCode      string  `json:"authCode"`
	HostReference string  `json:"hostReference"`
}

type refundPayment struct {
	baseResponse
	PaymentID            string  `json:"paymentId"`
	PaymentTransactionID string  `json:"paymentTransactionId"`
	Price                float64 `json:"price"`
	Currency             string  `json:"currency"`
	AuthCode             string  `json:"authCode"`
	HostReference        string  `json:"hostReference"`
	Retryable            bool    `json:"retryable"`
}
