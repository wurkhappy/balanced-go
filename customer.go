package balanced

import (
	"encoding/json"
	"time"
)

type Customer struct {
	Address        *Address               `json:"address,omitempty"`
	BusinessName   string                 `json:"business_name,omitempty"`
	CreatedAt      time.Time              `json:"created_at,omitempty"`
	DobMonth       int                    `json:"dob_month,omitempty"`
	DobYear        int                    `json:"dob_year,omitempty"`
	Ein            string                 `json:"ein,omitempty"`
	Email          string                 `json:"email,omitempty"`
	ID             string                 `json:"id,omitempty"`
	Meta           map[string]interface{} `json:"meta,omitempty"`
	Name           string                 `json:"name,omitempty"`
	Phone          string                 `json:"phone,omitempty"`
	SSNLast4       string                 `json:"ssn_last4,omitempty"`
	MerchantStatus string                 `json:"merchant_status,omitempty"`
}

type customerResponse struct {
	Customers []*Customer `json:"customers"`
}

func (c *Customer) path() string {
	return "/customers"
}

func (c *Customer) getID() string {
	return c.ID
}

func (c *Customer) getOwnerPath() string {
	return ""
}

func (c *Customer) singleResponse(data []byte) {
	parsedResponse := new(customerResponse)
	json.Unmarshal(data, &parsedResponse)
	*c = *parsedResponse.Customers[0]
}

func (c *Customer) canDelete() bool {
	return true
}

func (c *Customer) IsVerified() bool {
	return c.MerchantStatus == "underwritten"
}
