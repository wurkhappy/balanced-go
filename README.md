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


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Write your code **and unit tests**
4. Ensure all tests still pass (`go test`)
5. Commit your changes (`git commit -am 'Add some feature'`)
6. Push to the branch (`git push origin my-new-feature`)
7. Create new pull request
  
