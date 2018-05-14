package description_keeper

import (
    "testing"

    "gopkg.in/telegram-bot-api.v4"
)

func Test_shouldNotContainMessageForEmptyMessage(t *testing.T) {
    update := tgbotapi.Update{}

    compareTestBooleans(t, isUpdateContainingMessage(update), false)
}

func Test_shouldNotContainMessageForEmptyMessageText(t *testing.T) {
    message := tgbotapi.Message{}
    // for `IsCommand()` to return true
    message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 16}}

    update := tgbotapi.Update{Message: &message}

    compareTestBooleans(t, isUpdateContainingMessage(update), false)
}

func Test_shouldNotContainMessageForNotCommand(t *testing.T) {
    message := tgbotapi.Message{Text: "/command@testbot"}

    update := tgbotapi.Update{Message: &message}

    compareTestBooleans(t, isUpdateContainingMessage(update), false)
}

func Test_shouldContainMessage(t *testing.T) {
    message := tgbotapi.Message{Text: "/command@testbot"}
    // for `IsCommand()` to return true
    message.Entities = &[]tgbotapi.MessageEntity{{Type: "bot_command", Offset: 0, Length: 16}}

    update := tgbotapi.Update{Message: &message}

    compareTestBooleans(t, isUpdateContainingMessage(update), true)
}

func Test_shouldExtractFromSet(t *testing.T) {
    compareTestStrings(t, extractSetText("/set Hello world!\nHow are you?"), "Hello world!\nHow are you?")
}

func Test_shouldExtractFromAdd(t *testing.T) {
    compareTestStrings(t, extractAddText("/add I add this sentence"), "I add this sentence")
}

func Test_shouldAddNewLine(t *testing.T) {
    compareTestStrings(t, addText("Hello world!", "I add this sentence"), "Hello world!\nI add this sentence")
}

func Test_shouldNotAddNewLineForEmptyString(t *testing.T) {
    compareTestStrings(t, addText("", "I add this sentence"), "I add this sentence")
}

func Test_shouldExtractFromRemove(t *testing.T) {
    compareTestStrings(t, extractRemoveText("/remove I remove this sentence"), "I remove this sentence")
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
