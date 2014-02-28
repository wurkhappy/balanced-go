package balanced

import (
	"testing"
)

func setUp() {
	Username = "ak-test-2redrYtUIi5rOqHOukWRCbyg8S2jBG5nr"
}

var globalVerificationURI string

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

func setUpVerification() *BankAccount {
	account := setUpAccount()
	if globalVerificationURI == "" {
		_, _ = account.Verify()
		globalVerificationURI = account.VerificationURI
	}
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
	bError := account.Create()
	if bError != nil {
		t.Errorf("Error returned:%q", bError)
	}
	if account.URI == "" {
		t.Error("No bank account id was returned")
	}
}

func Test_Retrieve(t *testing.T) {
	setUp()
	account := setUpAccount()
	bError := account.Retrieve()

	if bError != nil {
		t.Errorf("Error returned:%q", bError)
	}
}

func Test_Delete(t *testing.T) {
	setUp()
	account := setUpAccount()
	bError := account.Delete()

	if bError != nil {
		t.Errorf("Error returned:%q", bError)
	}
}

func Test_Verify(t *testing.T) {
	setUp()
	account := setUpAccount()
	_, bError := account.Verify()
	globalVerificationURI = account.VerificationURI

	if bError != nil {
		t.Errorf("Error returned:%q", bError)
	}
}
