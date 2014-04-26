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

var (
	apiKeyTestType       string = "APIKey"
	callbackTestType     string = "Callback"
	cardTestType         string = "Card"
	cardHoldTestType     string = "CardHold"
	customerTestType     string = "Customer"
	debitForCardTestType string = "Debit - Card"
	orderTestType        string = "Order"
	refundTestType       string = "Refund"
	reversalTestType     string = "Reversal"
)

type testCase struct {
	testType string
	resource Resourcer
}

func createCases() []*testCase {
	rand.Seed(time.Now().Unix())
	card := &Card{
		ExpirationMonth: 12,
		CVV:             "123",
		Number:          "5105105105105100",
		ExpirationYear:  2020,
	}
	customer := &Customer{
		Name:     "Henry Ford",
		DobYear:  1963,
		DobMonth: 7,
	}
	debit := &Debit{
		Amount:               5000,
		AppearsOnStatementAs: "Test_Card",
		Description:          "Some descriptive text",
		Owner:                card,
	}
	return []*testCase{
		{
			cardTestType,
			card,
		},
		{
			apiKeyTestType,
			&ApiKey{},
		},
		{
			callbackTestType,
			&Callback{
				Method: "POST",
				URL:    "http://www." + strconv.Itoa(rand.Int()) + ".com/callback",
			},
		},
		{
			cardHoldTestType,
			&CardHold{
				Amount:      5000,
				Description: "Some descriptive text",
				Owner:       card,
			},
		},
		{
			customerTestType,
			customer,
		},
		{
			debitForCardTestType,
			debit,
		},
		{
			orderTestType,
			&Order{
				Description: "Some descriptive text",
				Owner:       customer,
			},
		},
		{
			refundTestType,
			&Refund{
				Amount:      3000,
				Description: "Some descriptive text",
				Owner:       debit,
			},
		},
	}
}

func Test_Create(t *testing.T) {
	cases := createCases()
	for _, c := range cases {
		bErrors := Create(c.resource)
		if len(bErrors) > 0 {
			t.Errorf("Type - %q Error returned:%q", c.testType, bErrors, c.resource.getOwnerPath()+c.resource.path())
		}
		if c.resource.getID() == "" {
			t.Errorf("No id was returned for type:", c.testType)
		}
	}
}

func Test_Fetch(t *testing.T) {
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
		if _, ok := c.resource.(Deleter); !ok {
			continue
		}
		Create(c.resource)
		bErrors := Delete(c.resource.(Deleter))
		if len(bErrors) > 0 {
			t.Errorf("Type - %q Error returned:%q", c.testType, bErrors)
		}
	}
}
