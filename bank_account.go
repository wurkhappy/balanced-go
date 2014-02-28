package balanced

import (
	"encoding/json"
	"time"
)

type BankAccount struct {
	Id               string                 `json:"id,omitempty"`
	AccountNumber    string                 `json:"account_number,omitempty"`
	BankName         string                 `json:"bank_name,omitempty"`
	URI              string                 `json:"uri,omitempty"`
	CanDebit         bool                   `json:"can_debit,omitempty"`
	Name             string                 `json:"name,omitempty"`
	BankCode         string                 `json:"bank_code,omitempty"`
	RoutingNumber    string                 `json:"routing_number,omitempty"`
	Type             string                 `json:"type,omitempty"`
	Meta             map[string]interface{} `json:"meta,omitempty"`
	VerificationsURI string                 `json:"verifications_uri,omitempty"`
	VerificationURI  string                 `json:"verification_uri,omitempty"`
	CreditsURI       string                 `json:"credits_uri,omitempty"`
}

type Verification struct {
	Type              string    `json:"_type"`
	Uris              string    `json:"_uris"`
	Attempts          string    `json:"attempts"`
	CreatedAt         time.Time `json:"created_at"`
	Id                string    `json:"id"`
	RemainingAttempts string    `json:"remaining_attempts"`
	State             string    `json:"state"`
	UpdatedAt         time.Time `json:"updated_at"`
	Uri               string    `json:"uri"`
}

func (b *BankAccount) Create() []*BalancedError {
	jsonData, _ := json.Marshal(b)
	data, bError := apiRequest("POST", jsonData, apiVersion+"/bank_accounts")
	json.Unmarshal(data, &b)
	return bError
}

func (b *BankAccount) Retrieve() []*BalancedError {
	data, bError := apiRequest("GET", nil, b.URI)
	json.Unmarshal(data, &b)
	return bError
}

func (b *BankAccount) Delete() []*BalancedError {
	_, bError := apiRequest("DELETE", nil, b.URI)
	return bError
}

func (b *BankAccount) Verify() (*Verification, []*BalancedError) {
	data, bError := apiRequest("POST", nil, b.VerificationsURI)
	if bError != nil {
		return nil, bError
	}
	var verification *Verification
	json.Unmarshal(data, &verification)
	b.VerificationURI = verification.Uri
	return verification, bError
}

func (b *BankAccount) ConfirmVerification(amount1, amount2 float64) (*Verification, []*BalancedError) {
	m := map[string]int{
		"amount_1": int(round(amount1)),
		"amount_2": int(round(amount2)),
	}
	jsonData, _ := json.Marshal(m)
	data, bError := apiRequest("PUT", jsonData, b.VerificationURI)
	var v *Verification
	json.Unmarshal(data, &v)
	return v, bError
}

func round(value float64) int64 {
	if value < 0.0 {
		value -= 0.5
	} else {
		value += 0.5
	}
	return int64(value)
}

func (b *BankAccount) Credit(credit *Credit) []*BalancedError {
	jsonData, _ := json.Marshal(credit)
	data, bError := apiRequest("POST", jsonData, b.CreditsURI)
	json.Unmarshal(data, &credit)
	return bError
}
