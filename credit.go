package balanced

import (
	"encoding/json"
)

type Credit struct {
	Amount               int               `json:"amount,omitempty"`
	Description          string            `json:"description,omitempty"`
	AppearsOnStatementAs string            `json:"appears_on_statement_as,omitempty"`
	ReversalsURI         string            `json:"reversals_uri,omitempty"`
	Status               string            `json:"status,omitempty"`
	URI                  string            `json:"uri,omitempty"`
	Meta                 map[string]string `json:"meta,omitempty"`
}

func (c *Credit) Retrieve() []*BalancedError {
	data, bError := apiRequest("GET", nil, c.URI)
	json.Unmarshal(data, &c)
	return bError
}
