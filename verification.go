package balanced

import (
	"encoding/json"
	"time"
)

type Verification struct {
	Attempts           int       `json:"attempts,omitempty"`
	AttemptsRemaining  int       `json:"attempts_remaining,omitempty"`
	CreatedAt          time.Time `json:"created_at,omitempty"`
	DepositStatus      string    `json:"deposit_status,omitempty"`
	ID                 string    `json:"id,omitempty"`
	UpdatedAt          time.Time `json:"updated_at,omitempty"`
	VerificationStatus string    `json:"verification_status,omitempty"`
}

type verificationResponse struct {
	Verifications []*Verification `json:"bank_account_verifications"`
}

func (v *Verification) path() string {
	return "/verifications"
}

func (v *Verification) getID() string {
	return v.ID
}

func (v *Verification) getOwnerPath() string {
	if v.Owner == nil {
		return ""
	}
	return v.Owner.path() + "/" + v.Owner.getID()
}

func (v *Verification) singleResponse(data []byte) {
	parsedResponse := new(verificationResponse)
	json.Unmarshal(data, &parsedResponse)
	*v = *parsedResponse.Verifications[0]
}
