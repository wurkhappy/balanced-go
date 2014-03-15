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

func Create(resource Resource) []*BalancedError {
	jsonData, _ := json.Marshal(resource)
	data, bErrors := apiRequest("POST", jsonData, resource.getOwnerPath()+resource.path())
	if len(bErrors) > 0 {
		return bErrors
	}
	resource.singleResponse(data)
	return nil
}

func Fetch(resource Resource) []*BalancedError {
	data, bErrors := apiRequest("GET", nil, resource.path()+"/"+resource.getID())
	if len(bErrors) > 0 {
		return bErrors
	}
	resource.singleResponse(data)
	return nil
}

func Update(resource Resource) []*BalancedError {
	jsonData, _ := json.Marshal(resource)
	data, bErrors := apiRequest("PUT", jsonData, resource.path()+"/"+resource.getID())
	if len(bErrors) > 0 {
		return bErrors
	}
	resource.singleResponse(data)
	return nil
}

func Delete(resource Resource) []*BalancedError {
	_, bErrors := apiRequest("DELETE", nil, resource.path()+"/"+resource.getID())
	return bErrors
}
