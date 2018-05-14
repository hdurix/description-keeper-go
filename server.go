package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
    checkEnvironmentVariable("KVSTORE_TOKEN")
    checkEnvironmentVariable("TELEGRAM_BOT_ID")
    checkEnvironmentVariable("KVSTORE_COLLECTION_NAME")
	fmt.Println("server started")
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}

func checkEnvironmentVariable(name string) {
    if (os.Getenv(name) == "") {
        panic("Environment variable " + name + " is not present")
    }
}
