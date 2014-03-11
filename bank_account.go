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
}

type bankAccountResponse struct {
	BankAccounts []*BankAccount `json:"bank_accounts"`
}

//Creates a new BankAccount resource that represents a bank account funding instrument.
//NOTE: This method is not recommended for production environments. Please use balanced.js for creating bank accounts.
func (b *BankAccount) Create() []*BalancedError {
	jsonData, _ := json.Marshal(b)
	data, bErrors := apiRequest("POST", jsonData, "/bank_accounts")
	if len(bErrors) > 0 {
		return bErrors
	}
	parsedResponse := new(bankAccountResponse)
	json.Unmarshal(data, &parsedResponse)
	*b = *parsedResponse.BankAccounts[0]
	return nil
}

//Fetches the details of a previously created bank account.
//In order to retrieve, make an instance of BankAccount with the ID of the resource you wish to retrieve.
func (b *BankAccount) Retrieve() []*BalancedError {
	data, bErrors := apiRequest("GET", nil, "/bank_accounts/"+b.ID)
	if len(bErrors) > 0 {
		return bErrors
	}
	parsedResponse := new(bankAccountResponse)
	json.Unmarshal(data, &parsedResponse)
	*b = *parsedResponse.BankAccounts[0]
	return nil
}

//Update information for a BankAccount resource.
//NOTE: Once a bank account has been associated to a customer, it cannot be associated to another customer.
func (b *BankAccount) Update() []*BalancedError {
	jsonData, _ := json.Marshal(b)
	data, bErrors := apiRequest("PUT", jsonData, "/bank_accounts/"+b.ID)
	if len(bErrors) > 0 {
		return bErrors
	}
	parsedResponse := new(bankAccountResponse)
	json.Unmarshal(data, &parsedResponse)
	*b = *parsedResponse.BankAccounts[0]
	return nil
}

//Permanently delete a bank account. It cannot be undone. All associated credits with a deleted bank account will not be affected.
func (b *BankAccount) Delete() []*BalancedError {
	_, bErrors := apiRequest("DELETE", nil, "/bank_accounts/"+b.ID)
	return bErrors
}

//Charge a bank account.
//NOTE: A bank account must be verified with micro deposits before it can be debited. See Bank Account Verifications.
func (b *BankAccount) Debit(debit *Debit) (*Debit, []*BalancedError) {
	jsonData, _ := json.Marshal(debit)
	data, bErrors := apiRequest("POST", jsonData, "/bank_accounts/"+b.ID+"/debits")
	if len(bErrors) > 0 {
		return nil, bErrors
	}
	var response struct {
		Debits []*Debit `json:"debits"`
	}
	json.Unmarshal(data, &response)
	return response.Debits[0], bErrors
}

//Create a new bank account verification.
//This initiates the process of sending micro deposits to the bank account
//which will be used to verify bank account ownership when supplied during Confirm a Bank Account Verification.
//NOTE: If you're sending money to a bank account, known as issuing a credit, you do NOT need to verify the bank account
func (b *BankAccount) Verify() (*Verification, []*BalancedError) {
	data, bErrors := apiRequest("POST", nil, "/bank_accounts/"+b.ID+"/verifications")
	if len(bErrors) > 0 {
		return nil, bErrors
	}
	var response struct {
		Verifications []*Verification `json:"bank_account_verifications"`
	}
	json.Unmarshal(data, &response)
	return response.Verifications[0], bErrors
}

func round(value float64) int64 {
	if value < 0.0 {
		value -= 0.5
	} else {
		value += 0.5
	}
	return int64(value)
}
