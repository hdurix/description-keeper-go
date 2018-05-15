package description_keeper

import (
	"testing"
)

var (
	chatId int64 = 123
)

func Test_shouldUnmarshallJson(t *testing.T) {
	compareTestStrings(t, extractValueFromJson([]byte(`{"value":"Hello World!","created_at":1518687228.482573,"updated_at":1518691902.889267}`)), "Hello World!")
}

func Test_shouldGetItemUrl(t *testing.T) {
	compareTestStrings(t, getItemUrl(123), "https://kvstore.p.mashape.com/collections/messages/items/123")
}

func Test_shouldSetMessage(t *testing.T) {
	processUpdateMessage(chatId, "/set -tomato")
	compareTestStrings(t, getMessage(chatId), "-tomato")
}

func Test_shouldAddMessage(t *testing.T) {
	processUpdateMessage(chatId, "/add -ketchup")
	processUpdateMessage(chatId, "/add -mayo")
	compareTestStrings(t, getMessage(chatId), "-tomato\n-ketchup\n-mayo")
}

func Test_shouldRemoveMessage(t *testing.T) {
	processUpdateMessage(chatId, "/remove -ketchup")
	compareTestStrings(t, getMessage(chatId), "-tomato\n-mayo")
}

func Test_shouldGetMessage(t *testing.T) {
	processUpdateMessage(chatId, "/get")
}
