# Telegram Bot API helper for Golang

This package is for building Telegram Bots with or without webhook interface.

View the [documentation here](https://pkg.go.dev/github.com/meinside/telegram-bot-go).

## How to get

```
$ go get -u github.com/meinside/telegram-bot-go
```

## Usage

See codes in [samples/](https://github.com/meinside/telegram-bot-go/tree/master/samples).

## Test

With following environment variables:

```bash
$ export TOKEN="01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"
$ export CHAT_ID="-123456789"

# for verbose output messages
$ export VERBOSE=true
```

run tests with:

```bash
$ go test
```

## Not implemented (yet)

- [ ] [Telegram Passport](https://core.telegram.org/bots/api#telegram-passport)
- [ ] [Seamless Telegram Login](https://telegram.org/blog/privacy-discussions-web-bots#meet-seamless-web-bots)
- [ ] [Payments 2.0](https://core.telegram.org/bots/payments)
- [ ] [Telegram Mini Apps](https://core.telegram.org/bots/webapps)

## Todos

- [ ] (WIP) Add tests for every API method

