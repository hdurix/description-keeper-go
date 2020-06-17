package handler

import (
	"testing"
)

func Test_shouldUnmarshallJson(t *testing.T) {
	compareTestStrings(t, extractValueFromJson([]byte(`{"value":"Hello World!","created_at":1518687228.482573,"updated_at":1518691902.889267}`)), "Hello World!")
}

func Test_shouldGetItemUrl(t *testing.T) {
	compareTestStrings(t, getItemUrl(123, "messages"), "https://api.kvstore.io/collections/messages/items/123")
}
