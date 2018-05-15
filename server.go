package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"gopkg.in/telegram-bot-api.v4"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Telegram World!")
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var update tgbotapi.Update
	err := decoder.Decode(&update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	processUpdate(update)
}

func main() {
	checkEnvironmentVariable("KVSTORE_TOKEN")
	checkEnvironmentVariable("TELEGRAM_BOT_ID")
	checkEnvironmentVariable("KVSTORE_COLLECTION_NAME")
	log.Println("server started")
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/"+os.Getenv("TELEGRAM_BOT_ID")+"/send", sendHandler)
	http.ListenAndServe(":8080", nil)
}

func checkEnvironmentVariable(name string) {
	if os.Getenv(name) == "" {
		panic("Environment variable " + name + " is not present")
	}
}
