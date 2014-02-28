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
	r.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(r)
	if err != nil {
		fmt.Printf("Error : %s", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)

	errs := new(errors)
	if resp.StatusCode >= 400 {
		json.Unmarshal(buf.Bytes(), &errs)
	}
	return buf.Bytes(), errs.Errors
}
