package balanced

import (
	"time"
)

type Debit struct {
	Amount               int               `json:"amount,omitempty"`
	AppearsOnStatementAs string            `json:"appears_on_statement_as,omitempty"`
	CreatedAt            time.Time         `json:"created_at,omitempty"`
	Currency             string            `json:"currency,omitempty"`
	Description          string            `json:"description,omitempty"`
	ID                   string            `json:"id,omitempty"`
	Meta                 map[string]string `json:"meta,omitempty"`
	Status               string            `json:"status,omitempty"`
	TransactionNumber    string            `json:"transaction_number,omitempty"`
	UpdatedAt            time.Time         `json:"updated_at,omitempty"`
}
