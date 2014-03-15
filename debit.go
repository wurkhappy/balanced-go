package balanced

import (
	"encoding/json"
	"time"
)

type Debit struct {
	Amount               int               `json:"amount,omitempty"`
	AppearsOnStatementAs string            `json:"appears_on_statement_as,omitempty"`
	CreatedAt            time.Time         `json:"created_at,omitempty"`
	Currency             string            `json:"currency,omitempty"`
	Description          string            `json:"description,omitempty"`
	ExpiresAt            time.Time         `json:"expires_at,omitempty"`
	FailureReason        string            `json:"failure_reason,omitempty"`
	FailureReasonCode    int               `json:"failure_reason_code,omitempty"`
	ID                   string            `json:"id,omitempty"`
	Meta                 map[string]string `json:"meta,omitempty"`
	Status               string            `json:"status,omitempty"`
	TransactionNumber    string            `json:"transaction_number,omitempty"`
	UpdatedAt            time.Time         `json:"updated_at,omitempty"`
	Owner                Resource
}

type debitResponse struct {
	Debits []*Debit `json:"debits"`
}

func (d *Debit) path() string {
	return "/debits"
}

func (d *Debit) getID() string {
	return d.ID
}

func (d *Debit) getOwnerPath() string {
	if d.Owner == nil {
		return ""
	}
	return d.Owner.path() + "/" + d.Owner.getID()
}

func (d *Debit) singleResponse(data []byte) {
	parsedResponse := new(debitResponse)
	json.Unmarshal(data, &parsedResponse)
	*d = *parsedResponse.Debits[0]
}
