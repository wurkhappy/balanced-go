package balanced

import (
	"encoding/json"
)

type Credit struct {
	Debit
}

type creditResponse struct {
	Credits []*Credit `json:"credits"`
}

func (c *Credit) path() string {
	return "/credits"
}

func (c *Credit) getID() string {
	return c.ID
}

func (c *Credit) getOwnerPath() string {
	if c.Owner == nil {
		return ""
	}
	return c.Owner.path() + "/" + c.Owner.getID()
}

func (c *Credit) singleResponse(data []byte) {
	parsedResponse := new(creditResponse)
	json.Unmarshal(data, &parsedResponse)
	*c = *parsedResponse.Credits[0]
}
