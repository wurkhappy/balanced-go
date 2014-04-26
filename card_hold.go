package balanced

import (
	"encoding/json"
	"time"
)

type CardHold struct {
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
	Owner                Resourcer         `json:"-"`
	IsVoid               bool              `json:"is_void,omitempty"`
}

type cardHoldResponse struct {
	CardHolds []*CardHold `json:"card_holds"`
}

func (c *CardHold) path() string {
	return "/card_holds"
}

func (c *CardHold) getID() string {
	return c.ID
}

func (c *CardHold) getOwnerPath() string {
	if c.Owner == nil {
		return ""
	}
	return c.Owner.path() + "/" + c.Owner.getID()
}

func (c *CardHold) singleResponse(data []byte) {
	parsedResponse := new(cardHoldResponse)
	json.Unmarshal(data, &parsedResponse)
	*c = *parsedResponse.CardHolds[0]
}
