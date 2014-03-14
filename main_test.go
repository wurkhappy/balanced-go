package balanced

import (
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func init() {
	Username = "ak-test-2ADpvITfpgBn8uBzEGsQ2bIgWaftUWiul"
}

type testCase struct {
	testType string
	resource Resource
}

func createCases() []*testCase {
	rand.Seed(time.Now().Unix())
	return []*testCase{
		{
			"Bank Account",
			&BankAccount{
				RoutingNumber: "121000358",
				Type:          "checking",
				Name:          "Johan Bernoulli",
				AccountNumber: "9900000001",
			},
		},
		{
			"APi Key",
			&ApiKey{},
		},
		{
			"Callback",
			&Callback{
				Method: "POST",
				URL:    "http://www." + strconv.Itoa(rand.Int()) + ".com/callback",
			},
		},
		{
			"Card",
			&Card{
				ExpirationMonth: 12,
				CVV:             "123",
				Number:          "5105105105105100",
				ExpirationYear:  2020,
			},
		},
	}
}

func Test_Create(t *testing.T) {
	cases := createCases()
	for _, c := range cases {
		bErrors := Create(c.resource)
		if len(bErrors) > 0 {
			t.Errorf("Type - %q Error returned:%q", c.testType, bErrors)
		}
		if c.resource.getID() == "" {
			t.Errorf("No id was returned for type:", c.testType)
		}
	}
}

func Test_Retrieve(t *testing.T) {
	cases := createCases()
	for _, c := range cases {
		Create(c.resource)
		bErrors := Fetch(c.resource)
		if len(bErrors) > 0 {
			t.Errorf("Type - %q Error returned:%q", c.testType, bErrors)
		}
	}
}

func Test_Delete(t *testing.T) {
	cases := createCases()
	for _, c := range cases {
		Create(c.resource)
		bErrors := Delete(c.resource)
		if len(bErrors) > 0 {
			t.Errorf("Type - %q Error returned:%q", c.testType, bErrors)
		}
	}
}

func Test_Debit(t *testing.T) {
	cases := []struct {
		testType string
		resource Instrument
	}{
		// {
		// 	"Bank Account",
		// 	&BankAccount{
		// 		RoutingNumber: "121000358",
		// 		Type:          "checking",
		// 		Name:          "Johan Bernoulli",
		// 		AccountNumber: "9900000001",
		// 	},
		// },
		{
			"Card",
			&Card{
				ExpirationMonth: 12,
				CVV:             "123",
				Number:          "5105105105105100",
				ExpirationYear:  2020,
			},
		},
	}
	for _, c := range cases {
		Create(c.resource)
		d := &Debit{
			Amount:               5000,
			AppearsOnStatementAs: "Test_BankAccount",
			Description:          "Some descriptive text for the debit in the dashboard",
		}
		_, bErrors := Charge(c.resource, d)
		if len(bErrors) > 0 {
			t.Errorf("Type - %q Error returned:%q", c.testType, bErrors)
		}
	}
}
