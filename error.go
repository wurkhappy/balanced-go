package balanced

type BalancedError struct {
	Additional   string                 `json:"additional"`
	CategoryCode string                 `json:"category_code"`
	CategoryType string                 `json:"category_type"`
	Description  string                 `json:"description"`
	Extras       map[string]interface{} `json:"extras"`
	RequestId    string                 `json:"request_id"`
	Status       string                 `json:"status"`
	StatusCode   float64                `json:"status_code"`
}

func (b *BalancedError) Error() string {
	return b.Description
}
