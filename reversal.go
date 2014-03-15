package balanced

import (
	"encoding/json"
	"time"
)

type Reversal struct {
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
	Owner                Resource          `json:"-"`
}

type reversalResponse struct {
	Reversals []*Reversal `json:"reversals"`
}

func (r *Reversal) path() string {
	return "/reversals"
}

func (r *Reversal) getID() string {
	return r.ID
}

func (r *Reversal) getOwnerPath() string {
	if r.Owner == nil {
		return ""
	}
	return r.Owner.path() + "/" + r.Owner.getID()
}

func (r *Reversal) singleResponse(data []byte) {
	parsedResponse := new(reversalResponse)
	json.Unmarshal(data, &parsedResponse)
	*r = *parsedResponse.Reversals[0]
}
