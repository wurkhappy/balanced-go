package balanced

type Address struct {
	City        string `json:"city,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
	Line1       string `json:"line1,omitempty"`
	Line2       string `json:"line2,omitempty"`
	State       string `json:"state,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
}
