package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

var (
	KVSTORE_URL string = "https://kvstore.p.mashape.com"
	collectionName = os.Getenv("KVSTORE_COLLECTION_NAME")
)

type KvstoreMessage struct {
	Value string `json:"value"`
}

func getMessage(chatId int64) string {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	req, err := http.NewRequest(http.MethodGet, getItemUrl(chatId, collectionName), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-Mashape-Key", os.Getenv("KVSTORE_TOKEN"))

	res, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	return extractValueFromJson(body)

}

func putMessage(chatId int64, value string) {

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	jsonText := []byte(value)
	req, err := http.NewRequest(http.MethodPut, getItemUrl(chatId, collectionName), bytes.NewBuffer(jsonText))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("X-Mashape-Key", os.Getenv("KVSTORE_TOKEN"))

	_, getErr := client.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
}

func getItemUrl(chatId int64, collectionName string) string {
	return KVSTORE_URL + "/collections/" + collectionName + "/items/" + strconv.FormatInt(chatId, 10)
}

func extractValueFromJson(jsonText []byte) string {
	var message KvstoreMessage
	jsonErr := json.Unmarshal(jsonText, &message)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return message.Value
}
