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

## Todo

* add usage
* add license
* etc.

