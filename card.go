package balanced

import (
	"encoding/json"
	"time"
)

type Card struct {
	Address         *Address               `json:"address,omitempty"`
	CreatedAt       time.Time              `json:"created_at,omitempty"`
	CVV             string                 `json:"cvv,omitempty"`
	CVVMatch        string                 `json:"cvv_match,omitempty"`
	CVVResult       string                 `json:"cvv_result,omitempty"`
	ExpirationYear  int                    `json:"expiration_year,omitempty"`
	ExpirationMonth int                    `json:"expiration_month,omitempty"`
	ID              string                 `json:"id,omitempty"`
	UpdatedAt       time.Time              `json:"updated_at,omitempty"`
	Number          string                 `json:"number,omitempty"`
	Name            string                 `json:"name,omitempty"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
}

type cardResponse struct {
	Cards []*Card `json:"cards"`
}

func (c *Card) path() string {
	return "/cards"
}

func (c *Card) getID() string {
	return c.ID
}

func (c *Card) getOwnerPath() string {
	return ""
}

func (c *Card) singleResponse(data []byte) {
	parsedResponse := new(cardResponse)
	json.Unmarshal(data, &parsedResponse)
	*c = *parsedResponse.Cards[0]
}
