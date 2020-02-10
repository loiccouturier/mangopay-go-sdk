package service

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/58-facettes/mangopay-go-sdk/model"
)

var (
	BaseURL   string
	BasicAuth string
)

// DefaultClient is the default Client and is used by Get, Head, and Post.
var DefaultClient = &http.Client{}

func newRequestAndExecute(method, uri string, param interface{}) (int, []byte, error) {
	body, err := json.Marshal(param)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	r, err := http.NewRequest(method, BaseURL+uri, bytes.NewBuffer(body))
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	r.Header.Set("Authorization", BasicAuth)
	r.Header.Set("Content-Type", "application/json")

	client := DefaultClient
	resp, err := client.Do(r)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, nil, parseErrorResponse(data)
	}
	return resp.StatusCode, data, nil
}

func parseErrorResponse(data []byte) error {
	var errResp model.Error
	err := json.Unmarshal(data, &errResp)
	if err != nil {
		return err
	}
	return errResp
}