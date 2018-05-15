package main

import (
	"testing"
)

func Test_shouldAddNewLine(t *testing.T) {
	compareTestStrings(t, addText("Hello world!", "I add this sentence"), "Hello world!\nI add this sentence")
}

func Test_shouldNotAddNewLineForEmptyString(t *testing.T) {
	compareTestStrings(t, addText("", "I add this sentence"), "I add this sentence")
}

func Test_shouldRemoveText(t *testing.T) {
	compareTestStrings(t, removeText("- tomato\n- ketchup", "up"), "- tomato\n- ketch")
}

func Test_shouldRemoveFirstLine(t *testing.T) {
	compareTestStrings(t, removeText("- tomato\n- ketchup", "- tomato"), "\n- ketchup")
}

func Test_shouldRemoveSecondLine(t *testing.T) {
	compareTestStrings(t, removeText("- tomato\n- ketchup", "- ketchup"), "- tomato")
}

func Test_shouldNotRemoveText(t *testing.T) {
	compareTestStrings(t, removeText("- tomato\n- ketchup", "- cucumber"), "- tomato\n- ketchup")
}

func Test_shouldNotRemoveTextForEmptyString(t *testing.T) {
	compareTestStrings(t, removeText("- tomato\n- ketchup", ""), "- tomato\n- ketchup")
}

func Test_shouldComputeSendMessageUri(t *testing.T) {
	expected := "https://api.telegram.org/bot123:ABC/sendMessage?chat_id=123&disable_web_page_preview=true&text=Hello+World%21"
	compareTestStrings(t, computeSendMessageUri("Hello World!", 123, "123:ABC"), expected)
}
