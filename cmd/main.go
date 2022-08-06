package main

import (
	"log"
	"net/http"
	"io/ioutil"
)

func main() {
	// TODO: read what messages are sent to the bot
	// TODO: send a dummy response
	// TODO: save message to a tasks array
	// TODO: when a message is equal to "/list", reply with the tasks array

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
