package bksdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/naruebaet/bitkub-sdk/bksdk/api"
	"github.com/naruebaet/bitkub-sdk/bksdk/bkerr"
	"github.com/naruebaet/bitkub-sdk/bksdk/request"
	"github.com/naruebaet/bitkub-sdk/bksdk/response"
)

// Endpoint : /api/status
// Method : GET
func (bksdk *SDK) GetStatus() (response.Status, error) {
	// init response body
	var respBody response.Status

	targetUrl := bksdk.apiHost.JoinPath(api.Status)

	resp, body, errs := bksdk.req.Get(targetUrl.String()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint : /api/servertime
// Method : GET
func (bksdk *SDK) GetServerTime() (string, error) {

	targetUrl := bksdk.apiHost.JoinPath(api.ServertimeV3)

	resp, timestamp, errs := bksdk.req.Get(targetUrl.String()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return "0", errs[0]
	}
	return timestamp, nil
}

// Endpoint : /api/market/symbols
// Method : GET
func (bksdk *SDK) GetSymbols() (response.MarketSymbols, error) {
	// init response body
	var respBody response.MarketSymbols

	targetUrl := bksdk.apiHost.JoinPath(api.MarketSymbol)

	resp, body, errs := bksdk.req.Get(targetUrl.String()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Endpoint : /api/market/ticker
// Method : GET
/*
Query
	sym string The symbol (optional)
	e.g. thb_btc
*/
func (bksdk *SDK) GetTicker(sym string) (response.MarketTicker, error) {
	// init response body
	var respBody response.MarketTicker

	targetUrl := bksdk.apiHost.JoinPath(api.MarketTicker)
	queryValues := url.Values{}
	if sym != "" {
		queryValues.Add("sym", sym)
	}

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint : /api/market/trades
// Method : GET
/*
Query
	sym string The symbol (e.g. thb_btc)
	lmt int No. of limit to query recent trades
*/
func (bksdk *SDK) GetTrade(sym string, limit int) (response.MarketTrades, error) {
	// init response body
	var respBody response.MarketTrades

	targetUrl := bksdk.apiHost.JoinPath(api.MarketTrades)
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Endpoint : /api/market/bids
// Method : GET
/*
Query
	sym string The symbol (e.g. thb_btc)
	lmt int No. of limit to query open buy orders
*/
func (bksdk *SDK) GetBids(sym string, limit int) (response.MarketBids, error) {
	// init response body
	var respBody response.MarketBids

	targetUrl := bksdk.apiHost.JoinPath(api.MarketBids)
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Endpoint : /api/market/asks
// Method : GET
/*
Query
	sym string The symbol (e.g. thb_btc)
	lmt int No. of limit to query open buy orders
*/
func (bksdk *SDK) GetAsks(sym string, limit int) (response.MarketAsks, error) {
	// init response body
	var respBody response.MarketAsks

	targetUrl := bksdk.apiHost.JoinPath(api.MarketAsks)
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Endpoint : /api/market/books
// Method : GET
/*
Query
	sym string The symbol (e.g. thb_btc)
	lmt int No. of limit to query open sell orders
*/
func (bksdk *SDK) GetBooks(sym string, limit int) (response.MarketBooks, error) {
	// init response body
	var respBody response.MarketBooks

	targetUrl := bksdk.apiHost.JoinPath(api.MarketBooks)
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Endpoint : /api/market/books
// Method : GET
/*
Query
	sym string The symbol (e.g. thb_btc)
	lmt int No. of limit to query open orders
*/
func (bksdk *SDK) GetDepth(sym string, limit int) (response.MarketDepth, error) {
	// init response body
	var respBody response.MarketDepth

	targetUrl := bksdk.apiHost.JoinPath(api.MarketDepth)
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint : /tradingview/history
// Method : GET
/*
Query
	symbol string The symbol (e.g. BTC_THB)
	resolution string Chart resolution (1, 5, 15, 60, 240, 1D)
	from int Timestamp of the starting time (e.g. 1633424427)
	to int Timestamp of the ending time (e.g. 1633427427)
*/
func (bksdk *SDK) GetHistory(symbol string, resolution string, from int, to int) (response.TradingviewHistory, error) {
	// init response body
	var respBody response.TradingviewHistory

	// validate resolutions
	resl, err := request.ValidateResolution(resolution)
	if err != nil {
		return respBody, err
	}

	targetUrl := bksdk.apiHost.JoinPath(api.TradingviewHistory)
	queryValues := url.Values{}
	queryValues.Add("sym", symbol)
	queryValues.Add("resolution", resl)
	queryValues.Add("from", strconv.Itoa(from))
	queryValues.Add("to", strconv.Itoa(to))

	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
