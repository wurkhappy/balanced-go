package balanced

import (
	"testing"
)

func setUp() {
	Username = "ak-test-2ADpvITfpgBn8uBzEGsQ2bIgWaftUWiul"
}

func setUpAccount() *BankAccount {
	account := &BankAccount{
		RoutingNumber: "121000358",
		Type:          "checking",
		Name:          "Johan Bernoulli",
		AccountNumber: "9900000001",
	}
	account.Create()
	return account
}

func Test_Create(t *testing.T) {
	setUp()
	account := &BankAccount{
		RoutingNumber: "121000358",
		Type:          "checking",
		Name:          "Johan Bernoulli",
		AccountNumber: "9900000001",
	}
	bErrors := account.Create()
	t.Log(account)
	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}
	if account.ID == "" {
		t.Error("No bank account id was returned")
	}
}

func Test_Retrieve(t *testing.T) {
	setUp()
	account := setUpAccount()
	bErrors := account.Retrieve()

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}
}

func Test_Update(t *testing.T) {
	setUp()
	account := setUpAccount()
	account.Meta = map[string]interface{}{
		"updated": false,
	}
	bErrors := account.Update()

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}
	//Balanced API turns all meta properties into strings
	//Initially, a bool is passed so if a string is returned then the request was succesful
	if account.Meta["updated"].(string) == "false" {
		t.Error("Bank account was not updated")
	}
}

func Test_Delete(t *testing.T) {
	setUp()
	account := setUpAccount()
	bErrors := account.Delete()

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}
}

func Test_Debit(t *testing.T) {
	setUp()
	account := setUpAccount()
	debit := &Debit{
		Amount:               5000,
		AppearsOnStatementAs: "test",
		Description:          "Some descriptive text for the debit in the dashboard",
	}
	verification, _ := account.Verify()
	verification.Confirm(1, 1)
	updatedDebit, bErrors := account.Debit(debit)

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
		return
	}

	if updatedDebit.ID == "" {
		t.Error("Debit was not given an ID")
	}
	if updatedDebit.Status != "succeeded" {
		t.Errorf("Expected succeeded status, instead got: ", updatedDebit.Status)
	}
}

func Test_Verify(t *testing.T) {
	setUp()
	account := setUpAccount()
	verification, bErrors := account.Verify()

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}

	if verification.ID == "" {
		t.Error("Verification was not given an ID")
	}
}
