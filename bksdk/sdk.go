package bksdk

import (
	"net/url"

	"github.com/naruebaet/bitkub-sdk/bksdk/response"
	"github.com/parnurzeal/gorequest"
)

type SDK struct {
	apiHost   *url.URL
	req       *gorequest.SuperAgent
	apiKey    string
	apiSecret string
}

type SDKEndpoints interface {
	// public endpoints
	GetStatus() (response.Status, error)
	GetServerTime() (string, error)
	GetSymbols() ([]response.MarketSymbolsResult, error)
	GetTicker(sym string) (map[string]response.MarketTickerData, error)
	GetTrade(sym string, limit int) (response.MarketTradesResult, error)
	GetBids(sym string, limit int) (response.MarketResult, error)
	GetAsks(sym string, limit int) (response.MarketResult, error)
	GetBooks(sym string, limit int) (response.MarketBooksResult, error)
	GetDepth(sym string, limit int) (response.MarketDepth, error)
	GetHistory(symbol string, resolution string, from int, to int) (response.TradingviewHistory, error)

	// User secure endpoints
	TradingCredit() (float64, error)
	Limits() (response.LimitsResult, error)

	// Market secure endpoints
	Wallet() (response.WalletResult, error)
	Balances() (response.BalanceResult, error)
	PlaceBid(sym string, amt, rat float64, typ, client_id string) (response.PlaceBidResult, error)
	PlaceAsk(sym string, amt, rat float64, typ, client_id string) (response.PlaceAskResult, error)
	CancelOrder(sym, id, sd, hash string) (response.CancelOrder, error)
	WsToken() (token string, err error)
	MyOpenOrder(sym string) ([]response.MyOpenOrderResult, error)
	MyOrderHistory(sym string, page, limit, start, end int) ([]response.MyOrderHistoryResult, response.BKPaginate, error)
	OrderInfo(sym, orderId, side string) (response.OrderInfoResult, error)
	OrderInfoByHash(hash string) (response.OrderInfoResult, error)

	// Crypto secure endpoints
	CryptoInternalWithdraw(currency string, address string, memo string, amount float64) (response.InternalWithdrawResult, error)
	CryptoAddresses(page, limit int) ([]response.CryptoAddressesResult, response.BKPaginate, error)
	CryptoWithdraw(currency string, address string, memo string, amount float64, network string) (response.CryptoWithdrawResult, error)
	CryptoDepositHistory(page, limit int) ([]response.DepositHistoryResult, response.BKPaginate, error)
	CryptoWithdrawHistory(page, limit int) ([]response.WithdrawHistoryResult, response.BKPaginate, error)
	CryptoGenerateAddress(symbol string) ([]response.CryptoGenerateAddressResult, error)

	// Fiat secure endpoints
	FiatAccounts(page int, limit int) ([]response.FiatAccountsResult, response.BKPaginate, error)
	FiatWithdraw(id string, amt float64) (response.FiatWithdrawResult, error)
	FiatDepositHistory(page, limit int) ([]response.FiatDepositHistoryResult, response.BKPaginate, error)
	FiatWithdrawHistory(page, limit int) ([]response.FiatWithdrawHistoryResult, response.BKPaginate, error)
}

// New creates a new SDK instance with the provided apiKey and apiSecret.
// It initializes the gorequest super agent and sets the API host URL.
func New(apiKey, apiSecret string) SDKEndpoints {
	// Initialize the gorequest super agent
	req := gorequest.New()

	// Set the API host URL
	apiHostURL, _ := url.Parse("https://api.bitkub.com")

	// Create a new SDK instance
	sdk := &SDK{
		apiHost:   apiHostURL,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		req:       req,
	}

	return sdk
}
