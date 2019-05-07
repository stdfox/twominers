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

	account, err := client.GetAccounts(twominers.XMR, "WALLET_ID")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashrate: %d, Payments Total: %d\n", account.Hashrate, account.PaymentsTotal)

	blocks, err := client.GetBlocks(twominers.XMR)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Matured Total: %d\n", blocks.MaturedTotal)

	miners, err := client.GetMiners(twominers.XMR)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hashrate: %f, Miners Total: %d\n", miners.Hashrate, miners.MinersTotal)

	payments, err := client.GetPayments(twominers.XMR)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Payments Total: %d\n", payments.PaymentsTotal)

	stats, err := client.GetStats(twominers.XMR)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Candidates Total: %d\n", stats.CandidatesTotal)
}
```
```shell
$ go run main.go
Hashrate: 117528, Round Shares: 2835, Payments Total: 0
Matured Total: 134, Orphan Rate: 0.000000
Hashrate: 510126.000000, Miners Total: 82
Payments Total: 691
Candidates Total: 0
```

## License
[MIT License](https://github.com/stdfox/twominers/blob/master/LICENSE.md)

## Donation
If this project help you, you can give me a cup of coffee :)

XMR: `49Nz5mVA9sQjKrT65TgdEneiZo1oCp3n8bCtjA3qaCoa5cuPKxqWBcZfD1f1iv6ASjCQUK55m3r4iho7ivMcNvsLDnP3sqX`

BTC: `1Kyxw87175msSRkXez7iKrzZLBaPb7rp1H`
