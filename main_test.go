package main

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	model "github.com/felipemarques/form3-challenge-go-app/models"
	"github.com/google/uuid"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}

// global ID
var TEST_ACCOUNT_ID string

func setupNewOrganisationAccount() {

    myGuid, err := uuid.NewRandom()
    if err != nil {
        log.Fatalln(err)
    }

    attributes := model.NewAccountAttributes()
	attributes.SetClassification("Personal")
	attributes.SetAccountNumber("10000004")
	attributes.SetAlternativeNames(nil)
	attributes.SetBankID("400302")
	attributes.SetBankIDCode("GBDSC")
	attributes.SetBaseCurrency("GBP")
	attributes.SetBic("NWBKGB42")
	attributes.SetCountry("GB")
	attributes.SetIban("GB28NWBK40030212764204")
	attributes.SetName([]string{"Organisation Name Test"})

	account := new(model.AccountData)
	account.SetID(myGuid.String())
	account.SetOrganisationID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	account.SetType("accounts")
	account.SetAttributes(*attributes)

	accountDataRequest := new(model.AccountDataRequest)
	accountDataRequest.SetData(*account)

	jsonRequestMarshal, err := json.Marshal(accountDataRequest)
	if err != nil {
		log.Fatalln(err)
	}

	response := postRequest("/organisation/accounts", jsonRequestMarshal)
    
    if (response.StatusCode == 201) {
        TEST_ACCOUNT_ID = account.GetID()

        fmt.Println("Created ACCOUNT_ID = " + TEST_ACCOUNT_ID)
    }

}

func init() {
    fmt.Println("Starting with tests....")

    setupNewOrganisationAccount();

}

func TestSuccessCreateFakeOrganisationAccount(t *testing.T) {

    attributes := model.NewAccountAttributes()
	attributes.SetClassification("Personal")
	attributes.SetAccountNumber("10000004")
	attributes.SetAlternativeNames(nil)
	attributes.SetBankID("400302")
	attributes.SetBankIDCode("GBDSC")
	attributes.SetBaseCurrency("GBP")
	attributes.SetBic("NWBKGB42")
	attributes.SetCountry("GB")
	attributes.SetIban("GB28NWBK40030212764204")
	attributes.SetName([]string{"Organisation Name Test"})

	account := new(model.AccountData)
	account.SetID("")
	account.SetOrganisationID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	account.SetType("accounts")
	account.SetAttributes(*attributes)

	accountDataRequest := new(model.AccountDataRequest)
	accountDataRequest.SetData(*account)

	jsonRequestMarshal, err := json.Marshal(accountDataRequest)
	if err != nil {
		log.Fatalln(err)
	}

	response := postRequest("/organisation/accounts", jsonRequestMarshal)
    checkResponseCode(t, 201, response.StatusCode)
}

func TestSuccessCreateFakeOrganisationAccountWithAlternativeNames(t *testing.T) {

    attributes := model.NewAccountAttributes()
	attributes.SetClassification("Personal")
	attributes.SetAccountNumber("10000004")

	var alternativeNames [2]string
	alternativeNames[0] = "My Alternative Name 1"
	alternativeNames[1] = "My Alternative Name 2"
	attributes.SetAlternativeNames(alternativeNames[:])

	attributes.SetBankID("400302")
	attributes.SetBankIDCode("GBDSC")
	attributes.SetBaseCurrency("GBP")
	attributes.SetBic("NWBKGB42")
	attributes.SetCountry("GB")
	attributes.SetIban("GB28NWBK40030212764204")
	attributes.SetName([]string{"Organisation Name Test"})

	account := new(model.AccountData)
	account.SetID("")
	account.SetOrganisationID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	account.SetType("accounts")
	account.SetAttributes(*attributes)

	accountDataRequest := new(model.AccountDataRequest)
	accountDataRequest.SetData(*account)

	jsonRequestMarshal, err := json.Marshal(accountDataRequest)
	if err != nil {
		log.Fatalln(err)
	}

    response := postRequest("/organisation/accounts", jsonRequestMarshal)
    checkResponseCode(t, 201, response.StatusCode)
}

func TestFailedCreateFakeOrganisationAccount(t *testing.T) {

    attributes := model.NewAccountAttributes()
	attributes.SetClassification("Personal")
	attributes.SetAccountNumber("10000004")

	attributes.SetBankID("400302")
	attributes.SetBankIDCode("GBDSC")
	attributes.SetBaseCurrency("GBP")
	attributes.SetBic("NWBKGB42")
	
	account := new(model.AccountData)
	account.SetID("")
	account.SetOrganisationID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	account.SetType("accounts")
	account.SetAttributes(*attributes)

	accountDataRequest := new(model.AccountDataRequest)
	accountDataRequest.SetData(*account)

	jsonRequestMarshal, err := json.Marshal(accountDataRequest)
	if err != nil {
		log.Fatalln(err)
	}

    response := postRequest("/organisation/accounts", jsonRequestMarshal)
    checkResponseCode(t, 400, response.StatusCode)
}

func TestFailedCreateFakeOrganisationAccountWithIDExistent(t *testing.T) {

    attributes := model.NewAccountAttributes()
	attributes.SetClassification("Personal")
	attributes.SetAccountNumber("10000004")

	attributes.SetBankID("400302")
	attributes.SetBankIDCode("GBDSC")
	attributes.SetBaseCurrency("GBP")
	attributes.SetBic("NWBKGB42")
	
    attributes.SetCountry("GB")
	attributes.SetIban("GB28NWBK40030212764204")
	attributes.SetName([]string{"Organisation Name Test"})

	account := new(model.AccountData)
	account.SetID(TEST_ACCOUNT_ID)
	account.SetOrganisationID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	account.SetType("accounts")
	account.SetAttributes(*attributes)

	accountDataRequest := new(model.AccountDataRequest)
	accountDataRequest.SetData(*account)

	jsonRequestMarshal, err := json.Marshal(accountDataRequest)
	if err != nil {
		log.Fatalln(err)
	}

    response := postRequest("/organisation/accounts", jsonRequestMarshal)
    checkResponseCode(t, 409, response.StatusCode)
}

func TestFetchOrganisationAccount(t *testing.T) {
	
    response := getRequest("/organisation/accounts/" + TEST_ACCOUNT_ID)

    checkResponseCode(t, 200, response.StatusCode)

}

func TestDeleteOrganisationAccount(t *testing.T) {
	
    response := deleteRequest("/organisation/accounts/" + TEST_ACCOUNT_ID + "?version=0")

    checkResponseCode(t, 204, response.StatusCode)

}