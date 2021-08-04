package accountsClient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func Create() {
	requestBody := RequestBody{
		AccountData: AccountData{
			ID:             uuid.NewString(),
			OrganisationID: uuid.NewString(),
			Type:           "accounts",
			Attributes: AccountAttributes{
				Country: "GB",
				Name:    []string{"Holly Golightly"},
			},
		},
	}

	json, _ := json.Marshal(requestBody)

	response, _ := http.Post("http://localhost:8080/v1/organisation/accounts", "application/json", bytes.NewBuffer(json))

	fmt.Println(response.StatusCode)
}

type RequestBody struct {
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
