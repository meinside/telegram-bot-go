# telegram-bot-go/samples/wasm

**Note**: currently not working on browsers (with error: `dial tcp: lookup api.telegram.org: Protocol not available`)

[Wasm](https://webassembly.org/) sample application for telegram-bot-go.

## How to build

Build with following command:

```bash
$ GOOS=js GOARCH=wasm go build -o telegram.wasm main.go
```

## Configure

Edit `token` value in `index.html` file. Yours can be obtained from the [@BotFather](https://telegram.me/BotFather) bot.

## Run

Run any http server on this directory and open `index.html`.

```bash
# for example, start a simple webserver with ruby on port 8888,
$ ruby -rwebrick -e's=WEBrick::HTTPServer.new(Port:8888,DocumentRoot:Dir.pwd);trap("INT"){s.shutdown};s.start'

# and open the index.html file
$ open http://localhost:8888
```
