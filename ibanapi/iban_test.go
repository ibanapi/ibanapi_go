package ibanapi_go

import (
	"testing"
)

const APIKEY = "973CC96C0AFB5E8E03BA0152BDF7272D7D8BFCACE07DBB47FE03E580364ECB234"

func TestValidateIBAN(t *testing.T) {
	iban := "VA5900112300001234567"
	resp, err := ValidateIBAN(APIKEY, iban)
	if err != nil {
		t.Fatalf(`ValidateIBAN = %s, failed with error %v`, iban, err)
	}

	if resp.Result != 200 {
		t.Fatalf(`ValidateIBAN = %s, failed with error code %d - %s`, iban, resp.Result, resp.Message)
	}
}

func TestValidateIBANBasic(t *testing.T) {
	iban := "VA59001123000012345678"
	resp, err := ValidateIBANBasic(APIKEY, iban)
	if err != nil {
		t.Fatalf(`ValidateIBANBasic = %s, failed with error %v`, iban, err)
	}

	if resp.Result != 200 {
		t.Fatalf(`ValidateIBANBasic = %s, failed with error code %d - %s`, iban, resp.Result, resp.Message)
	}
}

func TestGetBalance(t *testing.T) {
	resp, err := GetBalance(APIKEY)
	if err != nil {
		t.Fatalf(`GetBalance, failed with error %v`, err)
	}

	if resp.Result != 200 {
		t.Fatalf(`GetBalance, failed with error code %d - %s`, resp.Result, resp.Message)
	}
}
