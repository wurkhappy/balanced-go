package balanced

// import (
// 	"testing"
// )

func setUpAccount() *BankAccount {
	account := &BankAccount{
		RoutingNumber: "121000358",
		Type:          "checking",
		Name:          "Johan Bernoulli",
		AccountNumber: "9900000001",
	}
	Create(account)
	return account
}

// func Test_BankAccount_Verify(t *testing.T) {
// 	account := setUpAccount()
// 	verification, bErrors := account.Verify()

// 	if len(bErrors) > 0 {
// 		t.Errorf("Error returned:%q", bErrors)
// 	}

// 	if verification.ID == "" {
// 		t.Error("Verification was not given an ID")
// 	}
// }
