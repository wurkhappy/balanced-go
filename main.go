package balanced

import (
	"encoding/json"
)

var (
	apiURL     string = "https://api.balancedpayments.com"
	apiVersion string = "/v1"
	Username   string
)

type Resource interface {
	path() string
	getID() string
	getOwnerPath() string
	singleResponse([]byte)
}

//Creates a resource.
//Verifications, CardHolds, Credits, Debits, Orders, Refunds and Reversals, have an Owner field which must point to the correct Resource in order to be created.
func Create(resource Resource) []*BalancedError {
	jsonData, _ := json.Marshal(resource)
	data, bErrors := apiRequest("POST", jsonData, resource.getOwnerPath()+resource.path())
	if len(bErrors) > 0 {
		return bErrors
	}
	resource.singleResponse(data)
	return nil
}

//Fetches a single resource.
func Fetch(resource Resource) []*BalancedError {
	data, bErrors := apiRequest("GET", nil, resource.path()+"/"+resource.getID())
	if len(bErrors) > 0 {
		return bErrors
	}
	resource.singleResponse(data)
	return nil
}

//Updates a resource.
func Update(resource Resource) []*BalancedError {
	jsonData, _ := json.Marshal(resource)
	data, bErrors := apiRequest("PUT", jsonData, resource.path()+"/"+resource.getID())
	if len(bErrors) > 0 {
		return bErrors
	}
	resource.singleResponse(data)
	return nil
}

//Deletes a resource.
//Please note that not all resources can be deleted.
//The resources that cannot be deleted are: Verification, CardHold, Credit, Debit, Order, Refund, Reversal
func Delete(resource Resource) []*BalancedError {
	_, bErrors := apiRequest("DELETE", nil, resource.path()+"/"+resource.getID())
	return bErrors
}
