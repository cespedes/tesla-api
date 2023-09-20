package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type API struct {
	apiEndPoint string
	apiToken    string
}

func New(apiEndPoint, apiToken string) (*API, error) {
	return &API{apiEndPoint: apiEndPoint, apiToken: apiToken}, nil
}

func (A *API) Request(method, url string, data []byte, dest any) error {
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", A.apiEndPoint, url), bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", A.apiToken))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("api: %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode >= 400 {
		var foo struct {
			Error string
		}
		err = json.Unmarshal(body, &foo)
		if err != nil {
			return fmt.Errorf("%s", resp.Status)
		}
		return fmt.Errorf("%s: %s", resp.Status, foo.Error)
	}
	if dest == nil {
		var foo any
		dest = &foo
	}
	err = json.Unmarshal(body, dest)
	if err != nil {
		return fmt.Errorf("invalid JSON: %q", string(body))
	}
	return nil
}

func (A *API) Get(url string, dest any) error {
	return A.Request("GET", url, nil, dest)
}

func (A *API) Post(url string, data []byte, dest any) error {
	return A.Request("POST", url, data, dest)
}

func (A *API) Put(url string, data []byte, dest any) error {
	return A.Request("PUT", url, data, dest)
}

func (A *API) Delete(url string, dest any) error {
	return A.Request("DELETE", url, nil, dest)
}
