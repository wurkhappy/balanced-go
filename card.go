package balanced

import (
	"encoding/json"
)

type Card struct {
	CardNumber      string                 `json:"card_number,omitempty"`
	ExpirationYear  int                    `json:"expiration_year,omitempty"`
	ExpirationMonth int                    `json:"expiration_month,omitempty"`
	SecurityCode    string                 `json:"security_code,omitempty"`
	Name            string                 `json:"name,omitempty"`
	PhoneNumber     string                 `json:"phone_number,omitempty"`
	City            string                 `json:"city,omitempty"`
	Region          string                 `json:"region,omitempty"`
	State           string                 `json:"state,omitempty"`
	PostalCode      string                 `json:"postal_code,omitempty"`
	StreetAddress   string                 `json:"street_address,omitempty"`
	CountryCode     string                 `json:"country_code,omitempty"`
	Meta            map[string]interface{} `json:"meta,omitempty"`
	Verify          bool                   `json:"verify,omitempty"`
	URI             string                 `json:"uri,omitempty"`
	IsValid         bool                   `json:"is_valid,omitempty"`
	LastFour        string                 `json:"last_four,omitempty"`
}

func (c *Card) Create(marketplaceID string) []*BalancedError {
	jsonData, _ := json.Marshal(c)
	data, bError := apiRequest("POST", jsonData, apiVersion+"/marketplaces/"+marketplaceID+"/cards")
	json.Unmarshal(data, &c)
	return bError
}

func (c *Card) Retrieve() []*BalancedError {
	data, bError := apiRequest("GET", nil, c.URI)
	json.Unmarshal(data, &c)
	return bError
}

func (c *Card) Update() []*BalancedError {
	jsonData, _ := json.Marshal(c)
	data, bError := apiRequest("PUT", jsonData, c.URI)
	json.Unmarshal(data, &c)
	return bError
}

func (c *Card) Delete() []*BalancedError {
	_, bError := apiRequest("DELETE", nil, c.URI)
	return bError
}
