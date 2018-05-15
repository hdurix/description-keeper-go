package main

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

func Test_shouldExtractFromRemove(t *testing.T) {
	compareTestStrings(t, extractRemoveText("/remove I remove this sentence"), "I remove this sentence")
}
