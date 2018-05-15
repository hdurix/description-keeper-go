package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"gopkg.in/telegram-bot-api.v4"
)

var (
	MESSAGE_SEPARATOR string = "\n"
	SET_COMMAND       string = "/set"
	GET_COMMAND       string = "/get"
	ADD_COMMAND       string = "/add"
	REMOVE_COMMAND    string = "/remove"
)

func processUpdate(update tgbotapi.Update) {
	if !isUpdateContainingMessage(update) {
		fmt.Printf("update does not contain any text message")
		return
	}

	processUpdateMessage(update.Message.Chat.ID, update.Message.Text)
}

func processUpdateMessage(chatId int64, text string) {
	fmt.Printf("[Chat Id: %d] Treat message %s\n", chatId, text)

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

func isUpdateContainingMessage(update tgbotapi.Update) bool {
	return update.Message != nil && update.Message.IsCommand() && update.Message.Text != ""
}

func extractSetText(text string) string {
	return text[len(SET_COMMAND)+1:]
}

func extractAddText(text string) string {
	return text[len(ADD_COMMAND)+1:]
}

func addText(previousText string, textToAdd string) string {
	if previousText == "" {
		return textToAdd
	}
	return previousText + MESSAGE_SEPARATOR + textToAdd
}

func extractRemoveText(text string) string {
	return text[len(REMOVE_COMMAND)+1:]
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
