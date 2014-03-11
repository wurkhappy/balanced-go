package balanced

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func apiRequest(verb string, jsonData []byte, path string) (data []byte, bErrors []*BalancedError) {
	client := &http.Client{}
	url := apiURL + path
	body := bytes.NewReader(jsonData)
	r, _ := http.NewRequest(verb, url, body)
	r.SetBasicAuth(Username, "")
	if jsonData != nil {
		r.Header.Set("Content-Type", "application/json;revision=1.1")
	}
	r.Header.Set("Accept", "application/vnd.api+json;revision=1.1")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	var errs struct {
		Errors []*BalancedError `json:"errors"`
	}
	if resp.StatusCode >= 400 {
		json.Unmarshal(buf.Bytes(), &errs)
	}
	return buf.Bytes(), errs.Errors
}
