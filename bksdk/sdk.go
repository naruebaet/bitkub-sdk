package bksdk

import (
	"net/url"

	"github.com/naruebaet/bitkub-sdk/bksdk/response"
	"github.com/parnurzeal/gorequest"
)

type SDK struct {
	apiHost   *url.URL
	apiKey    string
	apiSecret string
	req       *gorequest.SuperAgent
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
	MyOpenOrder(sym string) (response.MyOpenOrder, error)
	MyOrderHistory(sym string, page, limit, start, end int) (response.MyOrderHistory, error)
	OrderInfo(sym, orderId, side string) (response.OrderInfo, error)
	OrderInfoByHash(hash string) (response.OrderInfo, error)
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
