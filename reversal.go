package balanced

import (
	"encoding/json"
)

type Reversal struct {
	Debit
}

type reversalResponse struct {
	Reversals []*Reversal `json:"reversals"`
}

func (r *Reversal) path() string {
	return "/reversals"
}

func (r *Reversal) getID() string {
	return r.ID
}

func (r *Reversal) getOwnerPath() string {
	if r.Owner == nil {
		return ""
	}
	return r.Owner.path() + "/" + r.Owner.getID()
}

func (r *Reversal) singleResponse(data []byte) {
	parsedResponse := new(reversalResponse)
	json.Unmarshal(data, &parsedResponse)
	*r = *parsedResponse.Reversals[0]
}
