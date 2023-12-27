# Bitkub-SDK 🚀
To connect to the Bitkub public API and use the secure endpoint using an api key and api secret, all you need is the Bitkub SDK.

## How to use.
```Go
package main

import (
    "log"
    "fmt"
    "github.com/naruebaet/bitkub-sdk/bksdk"
)

// Example init. the bksdk
func main() {
    sdk := bksdk.New("apiKey","apiSecret")
    
    // Call sample functions using the non-secure endpoint, such GetStatus(). 
    resp, err := sdk.GetStatus(); 
    if err != nil {
        log.Println(err.Error())
    }

    // show to response from GetStatus();
    fmt.Println(resp) 
}
```

## Check error description with function.
you can use this function below for get error description with error code from bitkub public api
```Go
package main

import (
    "fmt"
    "github.com/naruebaet/bitkub-sdk/bksdk/bkerr"
)

func main() {
    // example : if error code is 2
    fmt.Println(bkerr.ErrorText(2)) // Missing X-BTK-APIKEY
}
```

## Functions
### Non-secure endpoints
All non-secure endpoints do not need authentication and use the method GET.
* ✅GetStatus();
* ✅GetServertime();
* ✅GetSymbols();
* ✅GetTicker();
* ✅GetTrades();
* ✅GetBids();
* ✅GetAsks();
* ✅GetBooks();
* ✅GetDepth();
* ✅GetHistory();

### Secure endpoints v3
All secure endpoints require authentication and use the method POST. These are old endpoints. We suspended the creation of old-version API keys using with the old secure endpoints. Please use the new secure endpoints V3 instead.

#### User endpoints
* ✅TradingCredits();
* ✅Limits();

#### Market endpoints
* ✅Wallet();
* ✅Balances();
* ✅PlaceBid();
* ✅PlaceAsk();
* ✅CancelOrder();
* ✅Wstoken();
* ✅MyOpenOrders();
* ✅MyOrderHistory();
* ✅OrderInfo();
* ✅OrderInfoByHash();

#### Crypto endpoints
* ✅CryptoInternalWithdraw();
* ✅CryptoAddresses();
* ✅CryptoWithdraw();
* ✅CryptoDepositHistory();
* ✅CryptoWithdrawHistory();
* ✅CryptoGenerateAddress();

#### Fiat endpoints
* ✅FiatAccounts();
* ✅FiatWithdraw();
* ✅FiatDepositHistory();
* ✅FiatWithdrawHistory();

#### Websocket channel
In this project, you can connect a websocket to a Bitkub websocket and read data from the websocket via the Golang channel, for [Example](examples/ws), here!
``` golang
// for example connection
// package...
// import...

func main(){
    // Create a context with a timeout of 1 minute.
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)

    // create channel reader
    reader := make(chan string)
	defer close(reader)

    // create stream name
    streamName := fmt.Sprintf(bksdk.WS_TICKER_STREAM,"thb_btc")

    bksdk.CreateWsConnection(streamName, reader, ctx)

    for {
        msg := <-reader
        if msg == "Connection closed" {
            return
        }
        fmt.Println(msg)
    }
}

```
 
#### Error codes
Refer to the following descriptions:

| Code | Description                                                        |
| ---- | ------------------------------------------------------------------ |
| 0    | No error                                                           |
| 1    | Invalid JSON payload                                               |
| 2    | Missing X-BTK-APIKEY                                               |
| 3    | Invalid API key                                                    |
| 4    | API pending for activation                                         |
| 5    | IP not allowed                                                     |
| 6    | Missing / invalid signature                                        |
| 7    | Missing timestamp                                                  |
| 8    | Invalid timestamp                                                  |
| 9    | Invalid user                                                       |
| 10   | Invalid parameter                                                  |
| 11   | Invalid symbol                                                     |
| 12   | Invalid amount                                                     |
| 13   | Invalid rate                                                       |
| 14   | Improper rate                                                      |
| 15   | Amount too low                                                     |
| 16   | Failed to get balance                                              |
| 17   | Wallet is empty                                                    |
| 18   | Insufficient balance                                               |
| 19   | Failed to insert order into db                                     |
| 20   | Failed to deduct balance                                           |
| 21   | Invalid order for cancellation (Unable to find OrderID or Symbol.) |
| 22   | Invalid side                                                       |
| 23   | Failed to update order status                                      |
| 24   | Invalid order for lookup                                           |
| 25   | KYC level 1 is required to proceed                                 |
| 30   | Limit exceeds                                                      |
| 40   | Pending withdrawal exists                                          |
| 41   | Invalid currency for withdrawal                                    |
| 42   | Address is not in whitelist                                        |
| 43   | Failed to deduct crypto                                            |
| 44   | Failed to create withdrawal record                                 |
| 45   | Nonce has to be numeric                                            |
| 46   | Invalid nonce                                                      |
| 47   | Withdrawal limit exceeds                                           |
| 48   | Invalid bank account                                               |
| 49   | Bank limit exceeds                                                 |
| 50   | Pending withdrawal exists                                          |
| 51   | Withdrawal is under maintenance                                    |
| 52   | Invalid permission                                                 |
| 53   | Invalid internal address                                           |
| 54   | Address has been deprecated                                        |
| 55   | Cancel only mode                                                   |
| 56   | User has been suspended from purchasing                            |
| 57   | User has been suspended from selling                               |
| 90   | Server error (please contact support)                              |

## Next in BKSDK v2
- Functional programming optimized
- Code refactor
- SDK file size optimized
- Improve function name

## Referrence
- Please follow the link if you need to read more content. [Official bitkub public api documents](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md)
