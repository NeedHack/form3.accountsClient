package accountsClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Create(accountData AccountData) bool {
	requestBody := Data{
		AccountData: accountData,
	}

	json, _ := json.Marshal(requestBody)

	response, err := http.Post("http://localhost:8080/v1/organisation/accounts", "application/json", bytes.NewBuffer(json))
	if err != nil {
		panic(err.Error())
	}
	defer response.Body.Close()

	return response.StatusCode == 201
}

func Fetch(accountId string) AccountData {
	response, err := http.Get("http://localhost:8080/v1/organisation/accounts/" + accountId)
	if err != nil {
		fmt.Println(err)
		panic(err.Error())
	}
	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	var data Data
	json.Unmarshal(body, &data)
	return data.AccountData
}

type Data struct {
	AccountData AccountData `json:"data"`
}

type AccountData struct {
	Attributes     AccountAttributes `json:"attributes,omitempty"`
	ID             string            `json:"id,omitempty"`
	OrganisationID string            `json:"organisation_id,omitempty"`
	Type           string            `json:"type,omitempty"`
	Version        *int64            `json:"version,omitempty"`
}

type AccountAttributes struct {
	Country string   `json:"country,omitempty"`
	Name    []string `json:"name,omitempty"`
}
