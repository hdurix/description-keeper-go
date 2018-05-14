package description_keeper

import (
    "testing"

    "gopkg.in/telegram-bot-api.v4"
)

func compareTestStrings(t *testing.T, extractedText string, expectedText string) {
    if extractedText != expectedText {
        t.Errorf("Must be %s but was %s", expectedText, extractedText)
    }
}

func compareTestBooleans(t *testing.T, calculatedBoolean bool, expectedBoolean bool) {
    if calculatedBoolean != expectedBoolean {
        t.Errorf("Must be %t but was %t", calculatedBoolean, expectedBoolean)
    }
}

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
    expected := "https://api.telegram.org/bot/bot123/sendMessage?text=Hello World!&chat_id=123&disable_web_page_preview=true"
    compareTestStrings(t, computeSendMessageUri("Hello World!", 123, "bot123"), expected)
}

/*

type data struct {
	number   int
	msg, out string
}

var testDatas = []data{
	data{1, "foo", "result: 1foo"},
	data{2, "bar", "result: 2bar"},
	data{3, "bir", "result: 3bir"},
	data{4, "bor", "result: 4bor"},
	data{5, "bur", "result: 5bur"},
}

func Test_simple_struct(t *testing.T) {

	for _, test := range testDatas {
		out := Simple(test.number, test.msg)
		if out != test.out {
			t.Errorf("Simple(%v, %v) = %v, want %v", test.number, test.msg, out, test.out)
		}
	}
}

func Test_simple_inner_struct(t *testing.T) {
	for _, test := range []struct {
		number   int
		msg, out string
	}{
		{1, "foo", "result: 1foo"},
		{2, "bar", "result: 2bar"},
		{3, "bir", "result: 3bir"},
		{4, "bor", "result: 4bor"},
		{5, "bur", "result: 5bur"},
	} {
		out := Simple(test.number, test.msg)
		if out != test.out {
			t.Errorf("Simple(%v, %v) = %v, want %v", test.number, test.msg, out, test.out)
		}
	}
}
*/