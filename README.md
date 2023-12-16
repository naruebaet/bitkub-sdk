# Bitkub-SDK
To connect to the Bitkub public API and use the secure endpoint using an api key and api secret, all you need is the Bitkub SDK.

## How to use.
```Go
package main

import (
    "log"
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
## Functions
### Non-secure endpoints
All non-secure endpoints do not need authentication and use the method GET.
* GetStatus();
* GetServertime();
* GetSymbols();
* GetTicker();
* GetTrades();
* GetBids();
* GetAsks();
* GetBooks();
* GetDepth();
* GetHistory();

### Secure endpoints v3
All secure endpoints require authentication and use the method POST. These are old endpoints. We suspended the creation of old-version API keys using with the old secure endpoints. Please use the new secure endpoints V3 instead.

#### User endpoints
* TradingCredits();
* Limits();

#### Market endpoints
* Wallet();
* Balances();
* PlaceBid();
* PlaceAsk();
* CancelOrder();
* Wstoken();
* MyOpenOrders();
* MyOrderHistory();
* OrderInfo();

#### Crypto endpoints
* InternalWithdraw();
* Addresses();
* Withdraw();
* DepositHistory();
* WithdrawHistory();
* GenerateAddress();

#### Fiat endpoints
* Accounts();
* Withdraw();
* DepositHistory();
* WithdrawHistory();

## Referrence
- Please follow the link if you need to read more content. [Official bitkub public api documents](https://github.com/bitkub/bitkub-official-api-docs/blob/master/restful-api.md)
