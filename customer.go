package balanced

import (
	"encoding/json"
)

type Customer struct {
	ID               string                 `json:"id,omitempty"`
	URI              string                 `json:"uri,omitempty"`
	DebitsURI        string                 `json:"debits_uri,omitempty"`
	Name             string                 `json:"name,omitempty"`
	Email            string                 `json:"email,omitempty"`
	SSNLast4         string                 `json:"ssn_last4,omitempty"`
	BusinessName     string                 `json:"business_name,omitempty"`
	Address          *Address               `json:"address,omitempty"`
	Meta             map[string]interface{} `json:"meta,omitempty"`
	Phone            string                 `json:"phone,omitempty"`
	Dob              string                 `json:"dob,omitempty"`
	Ein              string                 `json:"ein,omitempty"`
	Facebook         string                 `json:"facebook,omitempty"`
	Twitter          string                 `json:"twitter,omitempty"`
	IdentityVerified bool                   `json:"is_identity_verified,omitempty"`
}

type Address struct {
	Line1       string `json:"line_1"`
	Line2       string `json:"line_2"`
	City        string `json:"city"`
	State       string `json:"state"`
	PostalCode  string `json:"postal_code"`
	CountryCode string `json:"country_code"`
}

func (c *Customer) Create() []*BalancedError {
	jsonData, _ := json.Marshal(c)
	data, bError := apiRequest("POST", jsonData, apiVersion+"/customers")
	json.Unmarshal(data, &c)
	return bError
}

func (c *Customer) Update() []*BalancedError {
	jsonData, _ := json.Marshal(c)
	data, bError := apiRequest("PUT", jsonData, c.URI)
	json.Unmarshal(data, &c)
	return bError
}

func (c *Customer) AddBankAccount(b *BankAccount) []*BalancedError {
	m := map[string]interface{}{
		"bank_account_uri": b.URI,
	}
	jsonData, _ := json.Marshal(m)
	_, bError := apiRequest("PUT", jsonData, c.URI)
	return bError
}

func (c *Customer) AddCreditCard(card *Card) []*BalancedError {
	m := map[string]interface{}{
		"card_uri": card.URI,
	}
	jsonData, _ := json.Marshal(m)
	_, bError := apiRequest("PUT", jsonData, c.URI)
	return bError
}

func (c *Customer) Delete() []*BalancedError {
	_, bError := apiRequest("DELETE", nil, c.URI)
	return bError
}

func (c *Customer) Debit(debit *Debit) []*BalancedError {
	jsonData, _ := json.Marshal(debit)
	data, bError := apiRequest("POST", jsonData, c.DebitsURI)
	json.Unmarshal(data, &debit)
	return bError
}
