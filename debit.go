package balanced

import (
	"encoding/json"
)

type Debit struct {
	OnBehalfOfUri        string            `json:"on_behalf_of_uri,omitempty"`
	Amount               int               `json:"amount,omitempty"`
	AppearsOnStatementAs string            `json:"appears_on_statement_as,omitempty"`
	Meta                 map[string]string `json:"meta,omitempty"`
	Description          string            `json:"description,omitempty"`
	AccountUri           string            `json:"account_uri,omitempty"`
	CustomerUri          string            `json:"customer_uri,omitempty"`
	HoldUri              string            `json:"hold_uri,omitempty"`
	SourceUri            string            `json:"source_uri,omitempty"`
	URI                  string            `json:"uri,omitempty"`
	RefundsURI           string            `json:"refunds_uri,omitempty"`
	Status               string            `json:"status,omitempty"`
}

func (d *Debit) Retrieve() []*BalancedError {
	data, bError := apiRequest("GET", nil, d.URI)
	json.Unmarshal(data, &d)
	return bError
}

func (d *Debit) Refunds() []*BalancedError {
	_, bError := apiRequest("POST", nil, d.RefundsURI)
	return bError
}
