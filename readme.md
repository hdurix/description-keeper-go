# Usage

Bot for storing message in Telegram discussion.

Can receive 4 commands:

- `/get`: Get channel description
- `/set`: Set channel description
- `/add`: Add line to channel description
- `/remove`: Remove line from channel description

All data is stored in [jsonbin.io][].

## Go

### Install Go

    https://golang.org/dl/

### Install packages

    go get gopkg.in/telegram-bot-api.v4

## Test

*For now, you need to move all tests into root folder before running the tests...* 

### Unit test

    go test
    
### End to end test
    
    JSONBIN_TOKEN=xxx JSONBIN_COLLECTION_NAME=collection_name TELEGRAM_BOT_ID=123:ABC CHAT_ID_TEST=1234 go test

## Vercel

### Install vercel

For example:

```sh
npm i -g vercel
```

### Get secret variables

```sh
vercel env --ls
```

### Deploy the app

```sh
vercel --prod
```

## Telegram

### Check connection

```sh
curl -i -X GET https://api.telegram.org/bot<apikey>/getMe
```

### Get Webhook Info

```sh
curl -i -X GET https://api.telegram.org/bot<apikey>/getWebhookInfo
```

### Post Webhook

```sh
DOMAIN=<domain>
APIKEY=<apikey>
curl -F "url=https://$DOMAIN/$APIKEY/send" https://api.telegram.org/bot$APIKEY/setWebhook
```

# Go

## References

- https://golang.org
- Download: https://golang.org/dl
- Documentation: https://golang.org/doc
- https://golang.org/doc/effective_go.html
- https://golang.org/ref/spec
- https://github.com/a8m/go-lang-cheat-sheet

## IDE

- https://github.com/golang/go/wiki/IDEsAndTextEditorPlugins
- Autocomplete dans les IDE : [Gocode](https://github.com/nsf/gocode)

## Articles

- http://decouvrir-golang.net/
- http://golang-examples.tumblr.com
- http://www.golangpatterns.info/object-oriented/operators
- http://devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/
- Un bon article sur la gestion des erreurs en Go. Utile lorsqu'on ne connaît que les exceptions ! : https://justinas.org/best-practices-for-errors-in-go/
- http://miek.nl/go/learninggo.html#beyond-the-basics
- http://etienner.fr/golang/creer-une-api-restfull-sur-go

## Resources

- https://gobyexample.com/
- https://github.com/golang/example
- https://github.com/golang-samples
- https://github.com/gobridge

## Testing

- http://golang.org/pkg/testing
- https://github.com/golang-samples/testing

Les méthodes de test doivent commencer par `Test`.

```go
import "testing"

func TestName(t *testing.T) {
    [...]
    if actual != expected {
    t.Errorf("Must be %s but was %s", expected, actual)
  }
}
```

## Test Tooling

- https://github.com/stretchr/testify
- http://goconvey.co
- https://labix.org/gocheck


[jsonbin.io]: http://www.jsonbin.io/
