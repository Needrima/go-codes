package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func checkIfEmailIsRegistered(email string) error {
	//{"debounce":{"email":"oyebodeamirdeen@gmail.com","code":"5","role":"false","free_email":"true","result":"Safe to Send","reason":"Deliverable","send_transactional":"1","did_you_mean":""},"success":"1","balance":"88"}
	type Debounce struct {
		Debounce struct {
			Email             string `json:"email"`
			Code              string `json:"code"`
			Role              string `json:"role"`
			FreeEmail         string `json:"free_email"`
			Result            string `json:"result"`
			Reason            string `json:"reason"`
			SendTransactional string `json:"send_transactional"`
			DidYouMean        string `json:"did_you_mean"`
		} `json:"debounce"`
		Success string `json:"success"`
		Balance string `json:"balance"`
	}

	access_key := os.Getenv("emailValidator_access_key")
	fmt.Println("access key:", access_key)
	resp, err := http.Get(fmt.Sprintf("https://api.debounce.io/v1/?api=%s&email=%s", access_key, email))
	if err != nil {
		log.Println("Sending a GET on email validator api, may be due to poor internet connection")
		return errors.New("something went wrong, check your internet connection and try again later")
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("could not read email validator response body")
		return errors.New("something went wrong, check your internet connection and try again later")
	}

	//fmt.Println(string(bs))

	m := Debounce{}

	err = json.Unmarshal(bs, &m)
	if err != nil {
		log.Println("Email response body unmarshal:", err)
		return errors.New("something went wrong, check your internet connection and try again later")
	}

	fmt.Printf("Result: %v, Reason: %v\n", m.Debounce.Result, m.Debounce.Reason)

	// result is usually "Safe to Send" and Reason is usually "Deliverable" for registered/reachable emails
	if m.Debounce.Result != "Safe to Send" || m.Debounce.Reason != "Deliverable" {
		return errors.New("unregistered")
	}

	return nil
}

func main() {
	if err := checkIfEmailIsRegistered("oyebodeamirdeen@gmail.com"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mail deliverable")
}
