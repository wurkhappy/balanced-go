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

func (a *ApiKey) getOwnerPath() string {
	return ""
}

func (a *ApiKey) singleResponse(data []byte) {
	parsedResponse := new(apiKeyResponse)
	json.Unmarshal(data, &parsedResponse)
	*a = *parsedResponse.ApiKeys[0]
}

func (a *ApiKey) canDelete() bool {
	return true
}
