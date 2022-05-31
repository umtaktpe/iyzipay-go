package utils

type RequestConvertible interface {
	ToPKIRequestString() string
	GetJsonObject() map[string]interface{}
}
