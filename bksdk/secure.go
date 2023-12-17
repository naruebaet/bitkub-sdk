package bksdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/naruebaet/bitkub-sdk/bksdk/api"
	"github.com/naruebaet/bitkub-sdk/bksdk/bkerr"
	"github.com/naruebaet/bitkub-sdk/bksdk/response"
)

// Endpoint : /api/v3/market/my-open-orders
// Method : GET
// Query : sym string The symbol (e.g. btc_thb)
func (bksdk *SDK) MyOpenOrder(sym string) (response.MyOpenOrder, error) {
	var respBody response.MyOpenOrder

	// join the targetUrl with host and path
	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOpenOrderV3)
	// init query values
	queryValues := url.Values{}
	queryValues.Add("sym", sym)

	resp, body, errs := bksdk.authGet(targetUrl, queryValues).End()
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

// Endpoint : /api/v3/market/my-order-history
// Method : GET
/*
Query
	sym string The symbol (e.g. btc_thb)
	p int Page (optional)
	lmt int Limit (optional)
	start int Start timestamp (optional)
	end int End timestamp (optional)
*/
func (bksdk *SDK) MyOrderHistory(sym string, page, limit, start, end int) (response.MyOrderHistory, error) {
	var respBody response.MyOrderHistory

	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOrderHistoryV3)
	queVal := url.Values{}
	queVal.Add("sym", sym)
	// optional parameter
	if page != 0 {
		queVal.Add("p", strconv.Itoa(page))
	}
	if limit != 0 {
		queVal.Add("lmt", strconv.Itoa(limit))
	}
	if start != 0 {
		queVal.Add("start", strconv.Itoa(start))
	}
	if end != 0 {
		queVal.Add("end", strconv.Itoa(end))
	}

	resp, body, errs := bksdk.authGet(targetUrl, queVal).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint : /api/v3/market/order-info
// Method : GET
/*
Query
	sym string The symbol (e.g. btc_thb)
	id string Order id
	sd string Order side: buy or sell
	hash string Lookup an order with order hash (optional). You don't need to specify sym, id, and sd when you specify order hash.
*/
func (bksdk *SDK) OrderInfo(sym, orderId, side string) (response.OrderInfo, error) {
	var respBody response.OrderInfo

	targetUrl := bksdk.apiHost.JoinPath(api.MarketOrderInfoV3)
	queVal := url.Values{}
	queVal.Add("sym", sym)
	queVal.Add("id", orderId)
	queVal.Add("sd", side)

	resp, body, errs := bksdk.authGet(targetUrl, queVal).End()
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

// Endpoint : /api/v3/market/order-info
// Method : GET
/*
Query
	hash string Lookup an order with order hash.
*/
func (bksdk *SDK) OrderInfoByHash(hash string) (response.OrderInfo, error) {
	var respBody response.OrderInfo

	targetUrl := bksdk.apiHost.JoinPath(api.MarketOrderInfoV3)
	queVal := url.Values{}
	queVal.Add("hash", hash)

	resp, body, errs := bksdk.authGet(targetUrl, queVal).End()
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
