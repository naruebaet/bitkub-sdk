package api

// Non-secure endpoints
// Method GET Only
const (
	Status             = "/api/status"
	Servertime         = "/api/servertime"
	MarketSymbol       = "/api/market/symbols"
	MarketTicker       = "/api/market/ticker"
	MarketTrade        = "/api/market/trades"
	MarketBids         = "/api/market/bids"
	MarketAsks         = "/api/market/asks"
	MarketBooks        = "/api/market/books"
	MarketDepth        = "/api/market/depth"
	TradingviewHistory = "/tradingview/history"
	ServertimeV3       = "/api/v3/servertime"
)

// Secure endpoints v1, v2
// Method POST Only
// With credential api signature
const (
	MarketWallet           = "/api/market/wallet"
	MarketBalances         = "/api/market/balances"
	MarketPlaceBid         = "/api/market/place-bid"
	MarketPlaceAsk         = "/api/market/place-ask"
	MarketPlaceBidTest     = "/api/market/place-bid/test"
	MarketPlaceAskTest     = "/api/market/place-ask/test"
	MarketPlaceAskByFiat   = "/api/market/place-ask-by-fiat"
	MarketCancelOrder      = "/api/market/cancel-order"
	MarketMyOpenOrder      = "/api/market/my-open-orders"
	MarketMyOrderHistory   = "/api/market/my-order-history"
	MarketOrderInfo        = "/api/market/order-info"
	CryptoAddresses        = "/api/crypto/addresses"
	CryptoWithdraw         = "/api/crypto/withdraw"
	CryptoInternalWithdraw = "/api/crypto/internal-withdraw"
	CryptoDepositHistory   = "/api/crypto/deposit-history"
	CryptoWithdrawHistory  = "/api/crypto/withdraw-history"
	CryptoGenerateAddress  = "/api/crypto/generate-address"
	FiatAccounts           = "/api/fiat/accounts"
	FiatWithdraw           = "/api/fiat/withdraw"
	FiatDepositHistory     = "/api/fiat/deposit-history"
	FiatWithdrawHistory    = "/api/fiat/withdraw-history"
	MarketWstoken          = "/api/market/wstoken"
	UserLimits             = "/api/user/limits"
	UserTradingCredits     = "/api/user/trading-credits"
	MarketPlaceBidV2       = "/api/market/v2/place-bid"
	MarketPlaceAskV2       = "/api/market/v2/place-ask"
	MarketCancelOrderV2    = "/api/market/v2/cancel-order"
)

// Secure endpoints v3
// Method POST and GET
// With credential api signature
const (
	MarketWalletV3           = "/api/v3/market/wallet"
	UserTradingCreditsV3     = "/api/v3/user/trading-credits"
	MarketPlaceBidV3         = "/api/v3/market/place-bid"
	MarketPlaceAskV3         = "/api/v3/market/place-ask"
	MarketCancelOrderV3      = "/api/v3/market/cancel-order"
	MarketBalancesV3         = "/api/v3/market/balances"
	MarketMyOpenOrderV3      = "/api/v3/market/my-open-orders"   // GET
	MarketMyOrderHistoryV3   = "/api/v3/market/my-order-history" // GET
	MarketOrderInfoV3        = "/api/v3/market/order-info"       // GET
	CryptoAddressesV3        = "/api/v3/crypto/addresses"
	CryptoWithdrawV3         = "/api/v3/crypto/withdraw"
	CryptoInternalWithdrawV3 = "/api/v3/crypto/internal-withdraw"
	CryptoDepositHistoryV3   = "/api/v3/crypto/deposit-history"
	CryptoWithdrawHistoryV3  = "/api/v3/crypto/withdraw-history"
	CryptoGenerateAddressV3  = "/api/v3/crypto/generate-address"
	FiatAccountsV3           = "/api/v3/fiat/accounts"
	FiatWithdrawV3           = "/api/v3/fiat/withdraw"
	FiatDepositHistoryV3     = "/api/v3/fiat/deposit-history"
	FiatWithdrawHistoryV3    = "/api/v3/fiat/withdraw-history"
	MarketWstokenV3          = "/api/v3/market/wstoken"
	UserLimitsV3             = "/api/v3/user/limits"
)
