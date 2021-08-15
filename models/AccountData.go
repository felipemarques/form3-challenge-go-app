package model

import (
	"log"

	"github.com/google/uuid"
)

type AccountDataRequest struct {
	Data *AccountData `json:"data"`
}

func (a *AccountDataRequest) SetData(accountData AccountData) {
	a.Data = &accountData
}

// Account represents an account in the form3 org section.
// See https://api-docs.form3.tech/api.html#organisation-accounts for
// more information about fields.
type AccountData struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
}

func New() *AccountData {
	return &AccountData{}
}

// Attributes
func (a *AccountData) SetAttributes(attributes AccountAttributes){
	a.Attributes = &attributes
}

func (a *AccountData) GetAttributes() AccountAttributes {
	return *a.Attributes
}

//ID
func (a *AccountData) SetID(id string) AccountData {
	if (id == "") {

		myGuid, err := uuid.NewRandom()
		if err != nil {
			log.Fatalln(err)
		}

		a.ID = myGuid.String()
	} else {
		a.ID = id
	}

	return *a
}

func (a *AccountData) GetID() string {
	return a.ID
}

// OrganisationID
func (a *AccountData) SetOrganisationID(organisationID string) {
	a.OrganisationID = organisationID
}

func (a *AccountData) GetOrganisationID() string {
	return a.OrganisationID
}

// Type
func (a *AccountData) SetType(myType string) {
	a.Type = myType
}

func (a *AccountData) GetType() string {
	return a.Type
}

// Version
func (a *AccountData) SetVersion(version int64) {
	a.Version = &version
}

func (a *AccountData) GetVersion() int64 {
	return *a.Version
}