package main

import (
	"log"
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
