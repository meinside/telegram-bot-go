# telegram-bot-go/samples/wasm

Go 1.11+

## How to build

Get wasm helper library,

```bash
$ go get -u github.com/meinside/wasm-helper-go
```

then edit **apiToken** value in `main.go` to yours, and build with following command:

```bash
$ GOOS=js GOARCH=wasm go build -o telegram.wasm main.go
```

## Run

Run any http server on this directory and open `index.html`.

```bash
# for example, start a simple webserver with ruby on port 8888,
$ ruby -rwebrick -e's=WEBrick::HTTPServer.new(Port:8888,DocumentRoot:Dir.pwd);trap("INT"){s.shutdown};s.start'

# and open the index.html file
$ open http://localhost:8888
```
