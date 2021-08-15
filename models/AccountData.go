package model

import (
	"log"

	"github.com/google/uuid"
)

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

// Attributes
func (a *AccountData) setAttributes(attributes AccountAttributes){
	a.Attributes = &attributes
}

func (a *AccountData) getAttributes() AccountAttributes {
	return *a.Attributes
}

//ID
func (a *AccountData) setID(id string) AccountData {
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

func (a *AccountData) getID() string {
	return a.ID
}

// OrganisationID
func (a *AccountData) setOrganisationID(organisationID string) {
	a.OrganisationID = organisationID
}

func (a *AccountData) getOrganisationID() string {
	return a.OrganisationID
}

// Type
func (a *AccountData) setType(myType string) {
	a.Type = myType
}

func (a *AccountData) getType() string {
	return a.Type
}

// Version
func (a *AccountData) setVersion(version int64) {
	a.Version = &version
}

func (a *AccountData) getVersion() int64 {
	return *a.Version
}