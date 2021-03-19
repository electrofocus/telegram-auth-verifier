# telegram-auth-verifier
Golang package for [Telegram Website Login](https://core.telegram.org/widgets/login#checking-authorization) credentials verification.

## Install
With a [correctly configured](https://golang.org/doc/install#testing) Go toolchain:

```go
go get github.com/electrofocus/telegram-auth-verifier
```

## Example

Let's verify credentials:

```go
import (
	"encoding/json"
	"fmt"

	tgverifier "github.com/electrofocus/telegram-auth-verifier"
)

func main() {
	token := []byte("Your Telegram Bot Token")

	rawCreds := `{
		"id": 111111111,
		"first_name": "John",
		"last_name": "Doe",
		"username": "johndoe",
		"auth_date": 1615974774,
		"hash": "ae1b08443b7bb50295be3961084c106072798cb65e91995a1b49927cd4cc5b0c"
	}`

	creds := tgverifier.Credentials{}
	json.Unmarshal([]byte(rawCreds), &creds)

	if err := creds.Verify(token); err != nil {
		fmt.Println("Credentials are not from Telegram")
		return
	}

	fmt.Println("Credentials are from Telegram")
}
```