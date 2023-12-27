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
	GetSymbols() (response.MarketSymbols, error)
	GetTicker(sym string) (response.MarketTicker, error)
	GetTrade(sym string, limit int) (response.MarketTrades, error)
	GetBids(sym string, limit int) (response.MarketBids, error)
	GetAsks(sym string, limit int) (response.MarketAsks, error)
	GetBooks(sym string, limit int) (response.MarketBooks, error)
	GetDepth(sym string, limit int) (response.MarketDepth, error)
	GetHistory(symbol string, resolution string, from int, to int) (response.TradingviewHistory, error)

	// User secure endpoints
	TradingCredit() (response.TradingCredit, error)
	Limits() (response.Limits, error)

	// Market secure endpoints
	Wallet() (response.Wallet, error)
	Balances() (response.Balances, error)
	PlaceBid(sym string, amt, rat float64, typ, client_id string) (response.PlaceBid, error)
	PlaceAsk(sym string, amt, rat float64, typ, client_id string) (response.PlaceAsk, error)
	CancelOrder(sym, id, sd, hash string) (response.CancelOrder, error)
	WsToken() (response.WsToken, error)
	MyOpenOrder(sym string) (response.MyOpenOrder, error)
	MyOrderHistory(sym string, page, limit, start, end int) (response.MyOrderHistory, error)
	OrderInfo(sym, orderId, side string) (response.OrderInfo, error)
	OrderInfoByHash(hash string) (response.OrderInfo, error)

	// Crypto secure endpoints
	CryptoInternalWithdraw(currency string, address string, memo string, amount float64) (response.InternalWithdraw, error)
	CryptoAddresses(page, limit int) (response.CryptoAddresses, error)
	CryptoWithdraw(currency string, address string, memo string, amount float64, network string) (response.CryptoWithdraw, error)
	CryptoDepositHistory(page, limit int) (response.DepositHistory, error)
	CryptoWithdrawHistory(page, limit int) (response.WithdrawHistory, error)
	CryptoGenerateAddress(symbol string) (response.CryptoGenerateAddress, error)

	// Fiat secure endpoints
	FiatAccounts(page int, limit int) (response.FiatAccounts, error)
	FiatWithdraw(id string, amt float64) (response.FiatWithdraw, error)
	FiatDepositHistory(page, limit int) (response.FiatDepositHistory, error)
	FiatWithdrawHistory(page, limit int) (response.FiatWithdrawHistory, error)
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
