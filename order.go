package balanced

import (
	"encoding/json"
	"time"
)

type Order struct {
	Amount          int               `json:"amount,omitempty"`
	AmountEscrowed  int               `json:"amount_escrowed,omitempty"`
	CreatedAt       time.Time         `json:"created_at,omitempty"`
	Currency        string            `json:"currency,omitempty"`
	DeliveryAddress *Address          `json:"delivery_address,omitempty"`
	Description     string            `json:"description,omitempty"`
	ID              string            `json:"id,omitempty"`
	Meta            map[string]string `json:"meta,omitempty"`
	UpdatedAt       time.Time         `json:"updated_at,omitempty"`
	Owner           Resource
}

type orderResponse struct {
	Orders []*Order `json:"orders"`
}

func (o *Order) path() string {
	return "/orders"
}

func (o *Order) getID() string {
	return o.ID
}

func (o *Order) getOwnerPath() string {
	if o.Owner == nil {
		return ""
	}
	return o.Owner.path() + "/" + o.Owner.getID()
}

func (o *Order) singleResponse(data []byte) {
	parsedResponse := new(orderResponse)
	json.Unmarshal(data, &parsedResponse)
	*o = *parsedResponse.Orders[0]
}
