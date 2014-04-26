# Balanced-go

A Balanced API library in Go

## Installation

go get github.com/wurkhappy/balanced-go

## Usage

```go
import "github.com/wurkhappy/balanced-go"
```

###Errors

This package uses a special error type to represent any errors returned by requests to the Balanced API
```go
type BalancedError struct {
	Additional   string                 `json:"additional"`
	CategoryCode string                 `json:"category_code"`
	CategoryType string                 `json:"category_type"`
	Description  string                 `json:"description"`
	Extras       map[string]interface{} `json:"extras"`
	RequestId    string                 `json:"request_id"`
	Status       string                 `json:"status"`
	StatusCode   float64                `json:"status_code"`
}
```

Invalid requests return an array of errors so all functions and methods in this package return
```go
[]*BalancedError
```

### Cards

#### Create a card

NOTE: This method is not recommended for production environments. Please use balanced.js to create cards.

```go
card := &balanced.Card{
		ExpirationMonth: 12,
		CVV:             "123",
		Number:          "5105105105105100",
		ExpirationYear:  2020,
	}
	
balanced.Create(card)
```

#### Fetch a card

Fetches the details of a card that has previously been created. Supply the ID that was returned from your previous request, and the corresponding card information will be returned. The same information is returned when creating the card

```go
card := &balanced.Card{
		ID: "CC2t9628l4ecJics6T8RuLPf",
	}
	
balanced.Fetch(card)
```

#### Update a card

Update information on a previously created card.
NOTE: Once a card has been associated to a customer, it cannot be associated to another customer.

```go
card := &balanced.Card{
		ID: "CC2t9628l4ecJics6T8RuLPf",
		Meta: map[string]interface{}{
		    "facebook.user_id":"0192837465",
		}
	}
	
balanced.Update(card)
```

#### Delete a card

Permanently delete a card. It cannot be undone. All debits associated with a deleted credit card will not be affected.

```go
card := &balanced.Card{
		ID: "CC2t9628l4ecJics6T8RuLPf",
	}
	
balanced.Delete(card)
```


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Write your code **and unit tests**
4. Ensure all tests still pass (`go test`)
5. Commit your changes (`git commit -am 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new pull request
  
