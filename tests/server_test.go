package handler

import (
	"os"
	"strconv"
	"testing"
)

var (
	chatId, _ = strconv.ParseInt(os.Getenv("CHAT_ID_TEST"), 10, 64)
)

func Test_shouldSetMessage(t *testing.T) {
	if isUnitTest() {
		return
	}
	processUpdateMessage(chatId, "/set -tomato")
	compareTestStrings(t, getMessage(chatId), "-tomato")
}

func Test_shouldAddMessage(t *testing.T) {
	if isUnitTest() {
		return
	}
	processUpdateMessage(chatId, "/add -ketchup")
	processUpdateMessage(chatId, "/add -mayo")
	compareTestStrings(t, getMessage(chatId), "-tomato\n-ketchup\n-mayo")
}

func Test_shouldRemoveMessage(t *testing.T) {
	if isUnitTest() {
		return
	}
	processUpdateMessage(chatId, "/remove -ketchup")
	compareTestStrings(t, getMessage(chatId), "-tomato\n-mayo")
}

func Test_shouldGetMessage(t *testing.T) {
	if isUnitTest() {
		return
	}
	processUpdateMessage(chatId, "/get")
}

func isUnitTest() bool {
	return os.Getenv("KVSTORE_TOKEN") == "" || os.Getenv("KVSTORE_COLLECTION_NAME") == "" || os.Getenv("TELEGRAM_BOT_ID") == ""
}
