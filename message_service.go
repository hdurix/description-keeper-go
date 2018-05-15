package main

import (
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)


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
