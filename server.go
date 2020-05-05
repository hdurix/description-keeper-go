package handler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	MESSAGE_SEPARATOR string = "\n"
	SET_COMMAND       string = "/set"
	GET_COMMAND       string = "/get"
	ADD_COMMAND       string = "/add"
	REMOVE_COMMAND    string = "/remove"

	KVSTORE_URL    string = "https://kvstore.p.rapidapi.com"
	collectionName string = os.Getenv("KVSTORE_COLLECTION_NAME")
	kvStoreToken   string = os.Getenv("KVSTORE_TOKEN")
)

func Handler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var update tgbotapi.Update
	err := decoder.Decode(&update)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	processUpdate(update)
}

// CONTROLLER

func processUpdate(update tgbotapi.Update) {
	if !isUpdateContainingMessage(update) {
		log.Printf("update does not contain any text message")
		return
	}

	processUpdateMessage(update.Message.Chat.ID, update.Message.Text)
}

func processUpdateMessage(chatId int64, text string) {
	log.Printf("[Chat Id: %d] Treat message %s\n", chatId, text)

	if strings.HasPrefix(text, SET_COMMAND) {
		setDescription(chatId, extractSetText(text))
	} else if strings.HasPrefix(text, ADD_COMMAND) {
		addDescription(chatId, extractAddText(text))
	} else if strings.HasPrefix(text, REMOVE_COMMAND) {
		removeDescription(chatId, extractRemoveText(text))
	} else if strings.HasPrefix(text, GET_COMMAND) {
		getDescription(chatId)
	}
}


func isUpdateContainingMessage(update tgbotapi.Update) bool {
	return update.Message != nil && update.Message.IsCommand() && update.Message.Text != ""
}

func extractSetText(text string) string {
	return text[len(SET_COMMAND)+1:]
}

func extractAddText(text string) string {
	return text[len(ADD_COMMAND)+1:]
}

func extractRemoveText(text string) string {
	return text[len(REMOVE_COMMAND)+1:]
}


// SERVICE

func getDescription(chatId int64) {
	message := getMessage(chatId)
	uri := computeSendMessageUri(message, chatId, os.Getenv("TELEGRAM_BOT_ID"))
	http.Get(uri)
}

func setDescription(chatId int64, text string) {
	newMessage := text
	putMessage(chatId, newMessage)
}

func addDescription(chatId int64, text string) {
	newMessage := addText(getMessage(chatId), text)
	putMessage(chatId, newMessage)
}

func removeDescription(chatId int64, text string) {
	newMessage := removeText(getMessage(chatId), text)
	putMessage(chatId, newMessage)
}

func addText(previousText string, textToAdd string) string {
	if previousText == "" {
		return textToAdd
	}
	return previousText + MESSAGE_SEPARATOR + textToAdd
}

func removeText(previousText string, textToRemove string) string {
	if textToRemove != "" && strings.Contains(previousText, MESSAGE_SEPARATOR+textToRemove) {
		textToRemove = MESSAGE_SEPARATOR + textToRemove
	}
	replacer := strings.NewReplacer(textToRemove, "")
	return replacer.Replace(previousText)
}

func computeSendMessageUri(message string, chatId int64, botId string) string {
	u := &url.URL{
		Scheme: "https",
		Host:   "api.telegram.org",
		Path:   "/bot" + botId + "/sendMessage",
	}
	q := u.Query()
	q.Set("text", message)
	q.Add("chat_id", strconv.FormatInt(chatId, 10))
	q.Add("disable_web_page_preview", "true")
	u.RawQuery = q.Encode()
	return u.String()
}

// REPOSITORY
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

	req.Header.Add("X-Mashape-Key", kvStoreToken)

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

	req.Header.Add("X-Mashape-Key", kvStoreToken)

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
