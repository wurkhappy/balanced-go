package balanced

import (
	"encoding/json"
)

type Refund struct {
	Debit
}

type refundResponse struct {
	Refunds []*Refund `json:"refunds"`
}

func (r *Refund) path() string {
	return "/refunds"
}

func (r *Refund) getID() string {
	return r.ID
}

func (r *Refund) getOwnerPath() string {
	if r.Owner == nil {
		return ""
	}
	return r.Owner.path() + "/" + r.Owner.getID()
}

func (r *Refund) singleResponse(data []byte) {
	parsedResponse := new(refundResponse)
	json.Unmarshal(data, &parsedResponse)
	*r = *parsedResponse.Refunds[0]
}
