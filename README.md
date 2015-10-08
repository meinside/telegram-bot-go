# Telegram Bot API helper for Golang

This package is for building a simple Telegram Bot with webhook interface.

## Install

```
$ go get github.com/meinside/telegram-bot-go
```

## Generate a self-signed certificate

Generate self-signed certificate & private key with following script:

```bash
#!/bin/bash
#
# Generate self-signed certificates for Telegram Bot API
# (https://core.telegram.org/bots/self-signed)
# 
# created on : 2015.10.08.
# last update: 2015.10.08.
# 
# by meinside@gmail.com

DOMAIN="some.where.in.the.galaxy"	# XXX - edit
EXPIRE_IN=3650	# XXX - edit

NUM_BITS=2048	# 2048 bits
C="US"
ST="New York"
L="Brooklyn"
O="Example Company"

PRIVATE_KEY="cert.key"
CERT_PEM="cert.pem"

openssl req -newkey rsa:$NUM_BITS -sha256 -nodes -keyout $PRIVATE_KEY -x509 -days $EXPIRE_IN -out $CERT_PEM -subj "/C=$C/ST=$ST/L=$L/O=$O/CN=$DOMAIN"

# finished
echo "> finished - certificate: $CERT_PEM / private key: $PRIVATE_KEY"
```

Generated *cert.key* and *cert.pem* file will be used in **telegrambot.StartWebhookServerAndWait()** function.

## Usage

```go
package main

import (
	"fmt"
	bot "github.com/meinside/telegram-bot-go"
	"net/http"
)

const (
	ApiToken     = "01234567:abcdefghijklmn_ABCDEFGHIJKLMNOPQRST"
	WebhookHost  = "my.host.com"
	WebhookPort  = 8443
	CertFilename = "cert.pem"
	KeyFilename  = "cert.key"
)

func main() {
	botapi := bot.NewClient(ApiToken)
	botapi.Verbose = true

	// set webhook url
	botapi.SetWebhookUrl(WebhookHost, WebhookPort, CertFilename, func(success bool, err error, description *string) {
		if success {
			fmt.Printf("SetWebhookUrl was successful: %s\n", *description)

			// on success, start webhook server
			botapi.StartWebhookServerAndWait(CertFilename, KeyFilename, func(success bool, err error, writer http.ResponseWriter, reqBody string) {
				if success {
					fmt.Printf(">>> %s\n", reqBody)
				}
			})
		} else {
			panic(fmt.Sprintf("failed to set webhook url (%s)\n", err.Error()))
		}
	})

	botapi.Wait()
}
```

## License

Copyright (c) 2015 Sungjin Han

MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

