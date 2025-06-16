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
		t.Skip("JSONBIN_TOKEN or JSONBIN_BIN_ID not set")
	}
	processUpdateMessage(chatId, "/set -tomato")
	compareTestStrings(t, getMessage(chatId), "-tomato")
}

func Test_shouldAddMessage(t *testing.T) {
	if isUnitTest() {
		t.Skip("JSONBIN_TOKEN or JSONBIN_BIN_ID not set")
	}
	processUpdateMessage(chatId, "/add -ketchup")
	processUpdateMessage(chatId, "/add -mayo")
	expected := "-tomato\n-ketchup\n-mayo"
	compareTestStrings(t, getMessage(chatId), expected)
}

func Test_shouldRemoveMessage(t *testing.T) {
	if isUnitTest() {
		t.Skip("JSONBIN_TOKEN or JSONBIN_BIN_ID not set")
	}
	processUpdateMessage(chatId, "/remove -ketchup")
	expected := "-tomato\n-mayo"
	compareTestStrings(t, getMessage(chatId), expected)
}

func Test_shouldGetMessage(t *testing.T) {
	if isUnitTest() {
		t.Skip("JSONBIN_TOKEN or JSONBIN_BIN_ID not set")
	}
	processUpdateMessage(chatId, "/get")
}

func isUnitTest() bool {
	return os.Getenv("JSONBIN_TOKEN") == "" || os.Getenv("JSONBIN_BIN_ID") == "" || os.Getenv("TELEGRAM_BOT_ID") == ""
}
