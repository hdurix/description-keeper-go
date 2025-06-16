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

	JSONBIN_URL    string = "https://api.jsonbin.io/v3/b"
	JSONBIN_HEADER string = "X-Master-Key"
	binId          string = os.Getenv("JSONBIN_BIN_ID")
	jsonbinToken   string = os.Getenv("JSONBIN_TOKEN")
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
	return textWithout(text, SET_COMMAND)
}

func extractAddText(text string) string {
	return textWithout(text, ADD_COMMAND)
}

func extractRemoveText(text string) string {
	return textWithout(text, REMOVE_COMMAND)
}

func textWithout(text string, partToRemove string) string {
	if (text == partToRemove) {
		return ""
	}

	return text[len(partToRemove)+1:]
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

func getMessage(chatId int64) string {
	client := &http.Client{Timeout: 10 * time.Second}

	req, err := http.NewRequest(http.MethodGet, JSONBIN_URL+"/"+binId+"/latest", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add(JSONBIN_HEADER, jsonbinToken)

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var parsed map[string]interface{}
	if err := json.Unmarshal(body, &parsed); err != nil {
		log.Fatal(err)
	}

	recordRaw, exists := parsed["record"]
	if !exists {
		log.Println("No record found in response")
		return ""
	}

	record, ok := recordRaw.(map[string]interface{})
	if !ok {
		log.Println("Invalid record format")
		return ""
	}

	val, exists := record[strconv.FormatInt(chatId, 10)]
	if !exists {
		return ""
	}

	valStr, ok := val.(string)
	if !ok {
		return ""
	}

	return valStr
}


func putMessage(chatId int64, value string) {
	// Step 1: Get current record
	client := &http.Client{Timeout: 10 * time.Second}
	getReq, _ := http.NewRequest(http.MethodGet, JSONBIN_URL+"/"+binId+"/latest", nil)
	getReq.Header.Add(JSONBIN_HEADER, jsonbinToken)
	getRes, err := client.Do(getReq)
	if err != nil {
		log.Fatal(err)
	}
	defer getRes.Body.Close()

	body, _ := ioutil.ReadAll(getRes.Body)
	var parsed map[string]interface{}
	json.Unmarshal(body, &parsed)
	var record map[string]interface{}
    if recordRaw, ok := parsed["record"]; ok {
        record, _ = recordRaw.(map[string]interface{})
    } else {
        record = make(map[string]interface{})
    }

	// Step 2: Modify in-memory
	record[strconv.FormatInt(chatId, 10)] = value

	// Step 3: Send updated record
	newBody, _ := json.Marshal(record)
	putReq, err := http.NewRequest(http.MethodPut, JSONBIN_URL+"/"+binId, bytes.NewBuffer(newBody))
	if err != nil {
		log.Fatal(err)
	}
	putReq.Header.Add(JSONBIN_HEADER, jsonbinToken)
	putReq.Header.Add("Content-Type", "application/json")
	_, err = client.Do(putReq)
	if err != nil {
		log.Fatal(err)
	}
}