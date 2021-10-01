package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	type validEmail struct{
		SmtpCheck bool `json:"smtp_check"`
		Score float64 `json:"score"`
	}

	access_key := os.Getenv("emailValidator_access_key")
	email := "oyebodeamirdeen@gmail.com"
	resp, err := http.Get(fmt.Sprintf("https://apilayer.net/api/check?access_key=%s&email=%s&smtp=1&format=1", access_key, email))
	if err != nil {
		log.Fatal("Resp:", err)
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll:", err)
	}

	fmt.Println(string(bs), "\n")

	m := validEmail{}

	err = json.Unmarshal(bs, &m)
	if err != nil {
		log.Fatal("Unmarshal:", err)
	}

	if !m.SmtpCheck || m.Score < 0.5 {
		fmt.Println("Unregistered Email Address")
	}else {
		fmt.Println(m)
	}
}
