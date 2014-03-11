package balanced

import (
	"testing"
)

func setUpVerification() *Verification {
	account := setUpAccount()
	verification, _ := account.Verify()
	return verification
}

func Test_Verification_Fetch(t *testing.T) {
	setUp()
	verification := setUpVerification()
	bErrors := verification.Fetch()

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}
}

func Test_Verification_Confirm(t *testing.T) {
	setUp()
	verification := setUpVerification()
	bErrors := verification.Confirm(1, 1)

	if len(bErrors) > 0 {
		t.Errorf("Error returned:%q", bErrors)
	}

	if verification.VerificationStatus != "succeeded" {
		t.Errorf("verification did not succeed. Status is: ", verification.VerificationStatus)
	}
}
