package balanced

import (
	"encoding/json"
)

type Callback struct {
	Href     string `json:"href,omitempty"`
	ID       string `json:"id,omitempty"`
	Method   string `json:"method,omitempty"`
	Revision string `json:"revision,omitempty"`
	URL      string `json:"url,omitempty"`
}

type callbackResponse struct {
	Callbacks []*Callback `json:"callbacks"`
}

func (c *Callback) path() string {
	return "/callbacks"
}

func (c *Callback) getID() string {
	return c.ID
}

func (c *Callback) getOwnerPath() string {
	return ""
}

func (c *Callback) singleResponse(data []byte) {
	parsedResponse := new(callbackResponse)
	json.Unmarshal(data, &parsedResponse)
	*c = *parsedResponse.Callbacks[0]
}

func (c *Callback) canDelete() bool {
	return true
}
