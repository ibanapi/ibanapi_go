package ibanapi_go

type BalanceResponse struct {
	Result  int          `json:"result"`
	Message string       `json:"message"`
	Data    *balanceData `json:"data,omitempty"`
}

type balanceData struct {
	BasicBalance int    `json:"basic_balance"`
	BankBalance  int    `json:"bank_balance"`
	ExpiryDate   string `json:"expiry_date"`
}

type IBANResponse struct {
	Result      int           `json:"result"`
	Message     string        `json:"message"`
	Validations []*Validation `json:"validations,omitempty"`
	Expremntal  int           `json:"expremental,omitempty"`
	Data        *Data         `json:"data,omitempty"`
}

type Validation struct {
	Result  int    `json:"result"`
	Message string `json:"message"`
}

type Data struct {
	CountryCode  string `json:"country_code"`
	IsoAlpha3    string `json:"iso_alpha3"`
	CountryName  string `json:"country_name"`
	CurrencyCode string `json:"currency_code"`
	SepaMember   string `json:"sepa_member"`
	Sepa         *Sepa  `json:"sepa,omitempty"`
	BBAN         string `json:"bban"`
	BankAccount  string `json:"bank_account"`
	Bank         *Bank  `json:"bank"`
}

type Sepa struct {
	CreditTransfer string `json:"sepa_credit_transfer"`
	DirectDebit    string `json:"sepa_direct_debit"`
	SDDCore        string `json:"sepa_sdd_core"`
	B2B            string `json:"sepa_b2b"`
	CardClearing   string `json:"sepa_card_clearing"`
}

type Bank struct {
	BankName string `json:"bank_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	BIC      string `json:"bic"`
	City     string `json:"city"`
	State    string `json:"state"`
	Zip      string `json:"zip"`
}
