package balanced

import (
	"encoding/json"
	"time"
)

type BankAccount struct {
	ID            string                 `json:"id,omitempty"`
	AccountNumber string                 `json:"account_number,omitempty"`
	Type          string                 `json:"account_type,omitempty"`
	BankName      string                 `json:"bank_name,omitempty"`
	CanDebit      bool                   `json:"can_debit,omitempty"`
	CanCredit     bool                   `json:"can_credit,omitempty"`
	CreatedAt     time.Time              `json:"created_at,omitempty"`
	UpdatedAt     time.Time              `json:"updated_at,omitempty"`
	URI           string                 `json:"uri,omitempty"`
	Name          string                 `json:"name,omitempty"`
	RoutingNumber string                 `json:"routing_number,omitempty"`
	Meta          map[string]interface{} `json:"meta,omitempty"`
	Customer      string                 `json:"customer,omitempty"`
}

type bankAccountResponse struct {
	BankAccounts []*BankAccount `json:"bank_accounts"`
}

func (b *BankAccount) path() string {
	return "/bank_accounts"
}

func (b *BankAccount) getID() string {
	return b.ID
}

func (b *BankAccount) getOwnerPath() string {
	return ""
}

func (b *BankAccount) singleResponse(data []byte) {
	parsedResponse := new(bankAccountResponse)
	json.Unmarshal(data, &parsedResponse)
	*b = *parsedResponse.BankAccounts[0]
}

func (b *BankAccount) AssociateWithCustomer(customer *Customer) []*BalancedError {
	b.Customer = customer.path() + "/" + customer.getID()
	return Update(b)
}

func (b *BankAccount) Verify() (*Verification, []*BalancedError) {
	verification := new(Verification)
	verification.Owner = b
	bErrors := Create(verification)
	return verification, bErrors
}
