# Balanced-go

A Balanced API library in Go

## Installation

go get github.com/wurkhappy/balanced-go

## Usage

```go
import "github.com/wurkhappy/balanced-go"
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
  
