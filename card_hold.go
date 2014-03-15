package balanced

import (
	"encoding/json"
)

type CardHold struct {
	Debit
	IsVoid bool `json:"is_void,omitempty"`
}

type cardHoldResponse struct {
	CardHolds []*CardHold `json:"card_holds"`
}

func (c *CardHold) path() string {
	return "/cards"
}

func (c *CardHold) getID() string {
	return c.ID
}

func (c *CardHold) getOwnerPath() string {
	if c.Owner == nil {
		return ""
	}
	return c.Owner.path() + "/" + c.Owner.getID()
}

func (c *CardHold) singleResponse(data []byte) {
	parsedResponse := new(cardHoldResponse)
	json.Unmarshal(data, &parsedResponse)
	*c = *parsedResponse.CardHolds[0]
}
