package balanced

import (
	"encoding/json"
	"time"
)

type Verification struct {
	Attempts           int       `json:"attempts"`
	AttemptsRemaining  int       `json:"attempts_remaining"`
	CreatedAt          time.Time `json:"created_at"`
	DepositStatus      string    `json:"deposit_status"`
	ID                 string    `json:"id"`
	UpdatedAt          time.Time `json:"updated_at"`
	VerificationStatus string    `json:"verification_status"`
}

type verificationResponse struct {
	Verifications []*Verification `json:"bank_account_verifications"`
}

//Fetches the verification for a bank account.
func (v *Verification) Fetch() []*BalancedError {
	data, bErrors := apiRequest("GET", nil, "/verifications/"+v.ID)
	if len(bErrors) > 0 {
		return bErrors
	}

	var parsedResponse verificationResponse
	json.Unmarshal(data, &parsedResponse)
	*v = *parsedResponse.Verifications[0]
	return nil
}

//Confirm the trial deposit amounts that were sent to the bank account.
//Upon seeing the verification amounts on their bank account statement, the customer should return to a web form and enter the amounts.
//The amounts entered are compared to the amounts sent to assert valid ownership of the bank account.
//NOTE: If you're sending money to a bank account, known as issuing a credit, you do NOT need to verify the bank account
func (v *Verification) Confirm(amount1, amount2 float64) []*BalancedError {
	m := map[string]int{
		"amount_1": int(round(amount1)),
		"amount_2": int(round(amount2)),
	}
	jsonData, _ := json.Marshal(m)
	data, bErrors := apiRequest("PUT", jsonData, "/verifications/"+v.ID)
	if len(bErrors) > 0 {
		return bErrors
	}

	var parsedResponse verificationResponse
	json.Unmarshal(data, &parsedResponse)
	*v = *parsedResponse.Verifications[0]
	return nil
}
