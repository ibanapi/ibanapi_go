package ibanapi_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const APIURL = "https://api.ibanapi.com/v1/"

func ValidateIBAN(apiKey, iban string) (*IBANResponse, error) {
	return doValidateIBAN(apiKey, iban, "validate")
}

func ValidateIBANBasic(apiKey, iban string) (*IBANResponse, error) {
	return doValidateIBAN(apiKey, iban, "validate-basic")
}

func GetBalance(apiKey string) (*BalanceResponse, error) {
	var balanceResponse = &BalanceResponse{}
	apiURL := fmt.Sprintf("%s%s", APIURL, "balance")
	form := url.Values{}
	form.Add("api_key", apiKey)

	resString, err := post(apiURL, apiKey, &form)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resString, balanceResponse)
	if err != nil {
		return nil, err
	}

	return balanceResponse, nil
}

func doValidateIBAN(apiKey, iban, endpoint string) (*IBANResponse, error) {
	var ibanResponse = &IBANResponse{}
	apiURL := fmt.Sprintf("%s%s", APIURL, endpoint)
	form := url.Values{}
	form.Add("api_key", apiKey)
	form.Add("iban", iban)

	resString, err := post(apiURL, apiKey, &form)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resString, ibanResponse)
	if err != nil {
		return nil, err
	}

	return ibanResponse, nil
}

/**
	Network Methods.
**/
func post(url, apiKey string, data *url.Values) ([]byte, error) {
	return send(url, apiKey, "POST", data)
}

func send(url, apiKey, method string, data *url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader([]byte(data.Encode())))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	return body, nil
}
