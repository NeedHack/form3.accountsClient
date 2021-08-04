package accountsClient

import (
	"testing"

	"github.com/google/uuid"
)

func Test_Create_should_return_true(t *testing.T) {

	accountData := AccountData{
		ID:             uuid.NewString(),
		OrganisationID: uuid.NewString(),
		Type:           "accounts",
		Attributes: AccountAttributes{
			Country: "GB",
			Name:    []string{"Holly Golightly"},
		},
	}
	response := Create(accountData)
	if !response {
		t.Fail()
	}
}

func Test_Fetch_should_return_AccountData(t *testing.T) {
	accountData := AccountData{
		ID:             uuid.NewString(),
		OrganisationID: uuid.NewString(),
		Type:           "accounts",
		Attributes: AccountAttributes{
			Country: "GB",
			Name:    []string{"Thelonious Monk"},
		},
	}
	Create(accountData)
	response := Fetch(accountData.ID)

	if response.Attributes.Name[0] != "Thelonious Monk" {
		t.Fail()
	}
}
