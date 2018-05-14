package description_keeper

import (
	"fmt"
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

// Simple main function
func Simple(number int, msg string) string {
	fmt.Printf("Kata number: %d, msg: %s\n", number, msg)
	return "result: " + strconv.Itoa(number) + msg
}

func processUpdate(update tgbotapi.Update) {
    if !isUpdateContainingMessage(update) {
        fmt.Printf("update does not contain any text message")
        return
    }

    text := update.Message.Text
    chatId := update.Message.Chat.ID
    fmt.Printf("[Chat Id: %d]Treat message %s", chatId, text)

}

func isUpdateContainingMessage(update tgbotapi.Update) bool {
	return update.Message != nil && update.Message.IsCommand() && update.Message.Text != ""
}

func extractSetText(text string) string {
	return text[len(SET_COMMAND)+1 : len(text)]
}

func extractAddText(text string) string {
	return text[len(ADD_COMMAND)+1 : len(text)]
}

func addText(previousText string, textToAdd string) string {
	if previousText == "" {
		return textToAdd
	}
	return previousText + MESSAGE_SEPARATOR + textToAdd
}

func extractRemoveText(text string) string {
	return text[len(REMOVE_COMMAND)+1 : len(text)]
}

func removeText(previousText string, textToRemove string) string {
	if textToRemove != "" && strings.Contains(previousText, MESSAGE_SEPARATOR+textToRemove) {
		textToRemove = MESSAGE_SEPARATOR + textToRemove
	}
	replacer := strings.NewReplacer(textToRemove, "")
	return replacer.Replace(previousText)
}

func computeSendMessageUri(message string, chatId int, botId string) string {
	return "https://api.telegram.org/bot/" + botId + "/sendMessage" + "?text=" + message + "&chat_id=" + strconv.Itoa(chatId) + "&disable_web_page_preview=true"
}
