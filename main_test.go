package main

import (
	"testing"
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

func TestCreateOrganisationAccount(t *testing.T) {
	
    getOrganizationAccounts()

	// req, _ := http.NewRequest("GET", "/products", nil)
    // response := executeRequest(req)

    // checkResponseCode(t, http.StatusOK, response.Code)

    // if body := response.Body.String(); body != "[]" {
    //     t.Errorf("Expected an empty array. Got %s", body)
    // }

}



func TestFetchOrganisationAccount(t *testing.T) {
	
}

func TestDeleteOrganisationAccount(t *testing.T) {
	
}