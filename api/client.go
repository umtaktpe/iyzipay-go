package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	model "github.com/umtaktpe/iyzipay-go/model"
	utils "github.com/umtaktpe/iyzipay-go/utils"
)

func request(method, url string, request utils.RequestConvertible, option *model.Options, respStruct interface{}) error {
	requestBody, err := json.Marshal(request.GetJsonObject())
	if err != nil {
		return err
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest(method, option.GetBaseUrl()+url, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	headers := getHttpHeaders(request, option)
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(body, &respStruct); err != nil {
		return err
	}

	return nil
}

func getHttpHeaders(request utils.RequestConvertible, option *model.Options) map[string]string {
	header := make(map[string]string)
	header["Accept"] = "application/json"
	header["Content-type"] = "application/json"

	rnd := utils.RandString(8)
	header["Authorization"] = prepareAuthString(option, rnd, request.ToPKIRequestString())
	header["x-iyzi-rnd"] = rnd
	header["x-iyzi-client-version"] = "iyzipay-php-2.0.51"

	return header
}

func prepareAuthString(options *model.Options, rnd, pki string) string {
	hashed := utils.GenerateHash(options.GetApiKey(), options.GetApiSecret(), rnd, pki)
	return formatHeaderString(options.GetApiKey(), hashed)
}

func formatHeaderString(api_key string, hashed string) string {
	return "IYZWS " + api_key + ":" + hashed
}
