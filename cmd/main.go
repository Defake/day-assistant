package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	token := ""
	url := "https://api.telegram.org/bot" + token + "/getMe"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
    
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
    
	sb := string(body)
	log.Printf(sb)
}
