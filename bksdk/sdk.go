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

type SDKFunction interface {
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

	// secure endpoints
	MyOpenOrder(sym string) (response.MyOpenOrder, error)
	MyOrderHistory(sym string, page, limit, start, end int) (response.MyOrderHistory, error)
	OrderInfo(sym, orderId, side string) (response.OrderInfo, error)
	OrderInfoByHash(hash string) (response.OrderInfo, error)
}

func New(apiKey, apiSecret string) SDKFunction {
	// init. gorequest super agent
	req := gorequest.New()

	// validate that is the url
	apiHostUrl, _ := url.Parse("https://api.bitkub.com")

	return &SDK{
		apiHost:   apiHostUrl,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		req:       req,
	}
}
