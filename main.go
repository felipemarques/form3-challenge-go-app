package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	model "github.com/felipemarques/form3-challenge-go-app/model"
)

const ENDPOINT string = "http://localhost:8080/v1"

func main() {

	//getOrganizationAccounts()	
	//createOrganisationAccount()
	//fetchOrganisationAccount()
	//deleteOrganisationAccount()
}

func postRequest(resource string, json []byte) {

	resp, err := http.Post(ENDPOINT + resource, "application/json; charset=utf-8", bytes.NewBuffer(json))
	if (err != nil) {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

func getOrganizationAccounts() {

	response, err := http.Get(ENDPOINT + "/organisation/accounts")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))

}

func createOrganisationAccount() {

	attributes := new(model.AccountAttributes)
	attributes.setClassification("Personal")
	//attributes.setAccountNumber("10000004")

	// var alternativeNames [2]string
	// alternativeNames[0] = "My Alternative Name 1"
	// alternativeNames[1] = "My Alternative Name 2"
	//attributes.setAlternativeNames(alternativeNames[:])
	//attributes.setAlternativeNames([]string{"My Alternative name1","name2"})
	// attributes.setAlternativeNames(nil)

	// attributes.setBankID("400302")
	// attributes.setBankIDCode("GBDSC")
	// attributes.setBaseCurrency("GBP")
	// attributes.setBic("NWBKGB42")
	// attributes.setCountry("GB")
	// attributes.setIban("GB28NWBK40030212764204")
	// attributes.setName([]string{"Organisation Name Test"})

	// account := new(AccountData)
	// account.setID("")
	// account.setOrganisationID("eb0bd6f5-c3f5-44b2-b677-acd23cdde73c")
	// account.setType("accounts")
	// account.setAttributes(*attributes)

	// jsonRequest := new(JsonRequest)
	// jsonRequest.setData(*account)

	// jsonRequestMarshal, err := json.Marshal(jsonRequest)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// postRequest("/organisation/accounts", jsonRequestMarshal)
	
}

func fetchOrganisationAccount() {

	id := "9af7feb1-2d94-45ce-8e4d-40e846831448"
	response, err := http.Get(ENDPOINT + "/organisation/accounts/" + id)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(responseData))
	
}

func deleteOrganisationAccount() {

	// Create client
	client := &http.Client{}

	id := "9af7feb1-2d94-45ce-8e4d-40e846831448"
	request, err := http.NewRequest("DELETE", ENDPOINT + "/organisation/accounts/" + id + "?version=0", nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	response, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer response.Body.Close()

	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response Status : ", response.Status)
    //fmt.Println("response Headers : ", response.Header)
	fmt.Println(string(responseBody))
	
}



