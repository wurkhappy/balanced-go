package balanced

import (
	"encoding/json"
	"time"
)

type ApiKey struct {
	CreatedAt time.Time              `json:"created_at,omitempty"`
	Href      string                 `json:"href,omitempty"`
	ID        string                 `json:"id,omitempty"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
	Secret    string                 `json:"secret,omitempty"`
}

type apiKeyResponse struct {
	ApiKeys []*ApiKey `json:"api_keys"`
}

func (a *ApiKey) path() string {
	return "/api_keys"
}

func (a *ApiKey) getID() string {
	return a.ID
}

func (a *ApiKey) singleResponse(data []byte) {
	parsedResponse := new(apiKeyResponse)
	json.Unmarshal(data, &parsedResponse)
	*a = *parsedResponse.ApiKeys[0]
}

//Create a new API key.
func (a *ApiKey) Create() []*BalancedError {
	data, bErrors := apiRequest("POST", nil, "/api_keys")
	if len(bErrors) > 0 {
		return bErrors
	}
	parsedResponse := new(apiKeyResponse)
	json.Unmarshal(data, &parsedResponse)
	*a = *parsedResponse.ApiKeys[0]
	return nil
}

//Get an existing API key.
func (a *ApiKey) Fetch() []*BalancedError {
	data, bErrors := apiRequest("GET", nil, "/api_keys/"+a.ID)
	if len(bErrors) > 0 {
		return bErrors
	}
	parsedResponse := new(apiKeyResponse)
	json.Unmarshal(data, &parsedResponse)
	*a = *parsedResponse.ApiKeys[0]
	return nil
}

//Delete an API key.
func (a *ApiKey) Delete() []*BalancedError {
	_, bErrors := apiRequest("DELETE", nil, "/api_keys/"+a.ID)
	return bErrors
}
