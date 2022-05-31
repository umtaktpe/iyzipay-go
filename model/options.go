package model

type Options struct {
	apiKey    string
	apiSecret string
	baseUrl   string
}

func NewOptions() *Options {
	return &Options{}
}

func (o *Options) GetApiKey() string {
	return o.apiKey
}

func (o *Options) SetApiKey(apiKey string) {
	o.apiKey = apiKey
}

func (o *Options) GetApiSecret() string {
	return o.apiSecret
}

func (o *Options) SetApiSecret(secretKey string) {
	o.apiSecret = secretKey
}

func (o *Options) GetBaseUrl() string {
	return o.baseUrl
}

func (o *Options) SetBaseUrl(baseUrl string) {
	o.baseUrl = baseUrl
}
