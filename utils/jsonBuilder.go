package utils

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type jsonBuilder struct {
	json map[string]interface{}
}

func NewJsonBuilder() *jsonBuilder {
	return &jsonBuilder{
		json: make(map[string]interface{}),
	}
}

func (j *jsonBuilder) AddPrice(key, value string) {
	if len(value) != 0 {
		valueToFloat, _ := strconv.ParseFloat(value, 64)
		stringValue := string(fmt.Sprintf("%.2f", valueToFloat))

		if (stringValue[len(stringValue)-1:]) == "0" {
			stringValue = stringValue[:len(stringValue)-1]
		}

		j.Add(key, stringValue)
	}
}

func (j *jsonBuilder) Add(key, value string) {
	if len(value) != 0 {
		j.json[key] = value
	}
}

func (j *jsonBuilder) AddModel(key string, value RequestConvertible) {
	if value != nil {
		j.json[key] = value.GetJsonObject()
	}
}

func (j *jsonBuilder) AddArray(key string, value []interface{}) {
	if value != nil {
		j.json[key] = value
	}
}

func (j *jsonBuilder) GetObject() map[string]interface{} {
	return j.json
}

func JsonEncode(jsonObject map[string]interface{}) (string, error) {
	result, err := json.Marshal(jsonObject)
	if err != nil {
		return "", err
	}
	return string(result), nil
}
