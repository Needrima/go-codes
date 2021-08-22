package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get(`http://localhost:8080/?key=spiderman`)
	if err != nil {
		log.Println("Resp err:", err)
		return
	}

	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Readall err:", err)
		return
	}

	fmt.Println(string(bs))
}
