package main

import (
	"log"
	"os"
	"net/http"
	"io/ioutil"
	"github.com/Defake/day-assistant/util/env"
	cf "github.com/Defake/day-assistant/util/controlflow"
	db "github.com/Defake/day-assistant/data_access/db"
)

func main() {
	// TODO: read what messages are sent to the bot
	// TODO: send a dummy response
	// TODO: save message to a tasks array
	// TODO: when a message is equal to "/list", reply with the tasks array

	env.ReadDevEnvs()
	token := os.Getenv("token")

	db.ConnectDb();
	
	url := "https://api.telegram.org/bot" + token + "/getMe"
	resp, err := http.Get(url)
	cf.Fatal(err)
    
	body, err := ioutil.ReadAll(resp.Body)
	cf.Fatal(err)
    
	sb := string(body)
	log.Printf(sb)
}
