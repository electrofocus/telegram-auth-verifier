# telegram-auth-verifier

[![Go Reference](https://pkg.go.dev/badge/github.com/electrofocus/telegram-auth-verifier.svg)](https://pkg.go.dev/github.com/electrofocus/telegram-auth-verifier)

## About

Here's the Golang package for [Telegram Website Login](https://core.telegram.org/widgets/login#checking-authorization) credentials verification.

This package uses [semantic versioning v2.0.0](https://semver.org/spec/v2.0.0.html), so you can be sure of backward compatibility of API.


## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain run:

```
go get github.com/electrofocus/telegram-auth-verifier
```

## Example

Let's verify credentials:

```go
package main

import (
	"encoding/json"
	"fmt"

	tgverifier "github.com/electrofocus/telegram-auth-verifier"
)

func main() {
	var (
		token = []byte("Your Telegram Bot Token")
		creds tgverifier.Credentials
	)

	rawCreds := `{
		"id": 111111111,
		"first_name": "John",
		"last_name": "Doe",
		"username": "johndoe",
		"auth_date": 1615974774,
		"hash": "ae1b08443b7bb50295be3961084c106072798cb65e91995a1b49927cd4cc5b0c"
	}`

	json.Unmarshal([]byte(rawCreds), &creds)

	if err := creds.Verify(token); err != nil {
		fmt.Println("Credentials are not from Telegram")
		return
	}

	fmt.Println("Credentials are from Telegram")
}
```

Open above example in [The Go Playground](https://play.golang.org/p/Pw2v8h4gbo1).
