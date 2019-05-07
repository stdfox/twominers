# twominers
A Golang wrapper for the [2miners](https://2miners.com) API.

## Documentation
Full API Documentation can be found at https://apidoc.2miners.com

All APIs return JSON format. Current API version: 2.0.0

## Installation
```shell
go get github.com/stdfox/twominers
```

## Importing
```go
import (
    "github.com/stdfox/twominers"
)
```

## Example
```go
package main

import (
	"fmt"

	"github.com/stdfox/twominers"
)

func main() {
	client := twominers.New()

	account, err := client.GetAccounts(twominers.SOLO_XMR, "WALLET_ID")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashrate: %d, Current Luck: %s\n", account.Hashrate, account.CurrentLuck)
}
```
