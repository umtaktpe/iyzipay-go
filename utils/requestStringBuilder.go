package utils

import (
	"fmt"
	"strconv"
)

type requestStringBuilder struct {
	requestString string
}

func NewRequestStringBuilder() *requestStringBuilder {
	return &requestStringBuilder{
		requestString: "",
	}
}

func (r *requestStringBuilder) AppendPrice(key, value string) {
	if len(value) != 0 {
		valueToFloat, _ := strconv.ParseFloat(value, 64)
		stringValue := string(fmt.Sprintf("%.2f", valueToFloat))

		if (stringValue[len(stringValue)-1:]) == "0" {
			stringValue = stringValue[:len(stringValue)-1]
		}

		r.appendKeyValue(key, stringValue)
	}
}

func (r *requestStringBuilder) Append(key, value string) {
	if len(value) != 0 {
		r.appendKeyValue(key, value)
	}
}

func (r *requestStringBuilder) AppendModel(key string, value RequestConvertible) {
	if value != nil {
		r.appendKeyValue(key, value.ToPKIRequestString())
	}
}

func (r *requestStringBuilder) AppendArray(key string, value []string) {
	if value != nil {
		appendedValue := ""
		for _, v := range value {
			appendedValue += v + ", "
		}
		r.appendKeyValueArray(key, appendedValue)
	}
}

func (r *requestStringBuilder) GetRequestString() string {
	r.removeTrailingComma()
	r.appendPrefix()
	return r.requestString
}

func (r *requestStringBuilder) appendKeyValue(key, value string) {
	if len(value) != 0 {
		r.requestString += key + "=" + value + ","
	}
}

func (r *requestStringBuilder) appendKeyValueArray(key string, appendedValue string) {
	appendedValue = appendedValue[:len(appendedValue)-2]
	r.requestString = r.requestString + key + "=[" + appendedValue + "],"
}

func (r *requestStringBuilder) removeTrailingComma() {
	r.requestString = r.requestString[:len(r.requestString)-1]
}

func (r *requestStringBuilder) appendPrefix() {
	r.requestString = "[" + r.requestString + "]"
}
