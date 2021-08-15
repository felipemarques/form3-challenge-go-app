package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const ENDPOINT string = "http://localhost:8080/v1"

func main() {

}

type Response struct {
	ResponseBody 	string 	`json:"responseBody"`
	StatusCode 		int	`json:"statusCode"` 
}

func postRequest(resource string, json []byte) Response {

	response, err := http.Post(ENDPOINT + resource, "application/json; charset=utf-8", bytes.NewBuffer(json))
	if (err != nil) {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	bodyString := string(responseData)

	defaultResponse := &Response{
		ResponseBody: bodyString,
		StatusCode: int(response.StatusCode),
	}

	return *defaultResponse
}

func getRequest(resource string) Response {

	response, err := http.Get(ENDPOINT + resource)
	if (err != nil) {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	bodyString := string(responseData)

	defaultResponse := &Response{
		ResponseBody: bodyString,
		StatusCode: int(response.StatusCode),
	}

	return *defaultResponse
}

func deleteRequest(resource string) Response {

	client := &http.Client{}

	request, err := http.NewRequest("DELETE", ENDPOINT + resource, nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	response, err := client.Do(request)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	bodyString := string(responseData)

	defaultResponse := &Response{
		ResponseBody: bodyString,
		StatusCode: int(response.StatusCode),
	}

	return *defaultResponse
}