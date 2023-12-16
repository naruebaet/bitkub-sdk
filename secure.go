package bitkubsdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/naruebaet/bitkub-sdk/api"
	"github.com/naruebaet/bitkub-sdk/bkError"
	"github.com/naruebaet/bitkub-sdk/response"
	"github.com/parnurzeal/gorequest"
)

// this is the function for generate signature pattern for bitkub
// Generate the signature from the timestamp, the request method, API path, query parameter, and JSON payload using HMAC SHA-256.
// Use the API Secret as the secret key for generating the HMAC variant of JSON payload. The signature is in hex format.
// The user has to attach the signature via the Request Header You must get a new timestamp in millisecond from /api/v3/servertime. The old one is in second.
/*
	//Example for Get Method
	1699381086593GET/api/v3/market/my-order-history?sym=BTC_THB

	// Example for Post Method
	1699376552354POST/api/v3/market/place-bid{"sym":"thb_btc","amt": 1000,"rat": 10,"typ": "limit"}
*/
func (bksdk *Bitkubsdk) generateSignature(timestamp, method, endpoint, bodyPayload string) (signature string) {
	// secret key for hmac sha256.
	secret := bksdk.apiSecret

	// init encryptor with secret key.
	// create a new HMAC by defining the hash type and the key
	encryptor := hmac.New(sha256.New, []byte(secret))

	// compute the HMAC
	encryptor.Write([]byte(timestamp + method + endpoint + bodyPayload))
	dataHmac := encryptor.Sum(nil)

	// to hex
	hexHmac := hex.EncodeToString(dataHmac)

	return hexHmac
}

// authGet is function for wrap request before send to secure endpoints
func (bksdk *Bitkubsdk) authGet(targetUrl *url.URL, queryValues url.Values) *gorequest.SuperAgent {
	// Step
	// 1.get servertime before generate signature
	// 2.then generate signature with timestamp, method, endpoint and payloads
	// 3.then set header with generated key and api key

	ts, _ := bksdk.GetServerTime()

	sig := bksdk.generateSignature(ts, "GET", "/"+targetUrl.Path, "?"+queryValues.Encode())

	return bksdk.req.Get(targetUrl.String()).
		Set("X-BTK-TIMESTAMP", ts).
		Set("X-BTK-APIKEY", bksdk.apiKey).
		Set("X-BTK-SIGN", sig).
		Set("Content-Type", "application/json").Query(queryValues.Encode())
}

// authPost is function for wrap request before send to secure endpoints
func (bksdk *Bitkubsdk) authPost(targetUrl *url.URL, jsonPayload string) *gorequest.SuperAgent {
	// Step
	// 1.get servertime before generate signature
	// 2.then generate signature with timestamp, method, endpoint and payloads
	// 3.then set header with generated key and api key

	ts, _ := bksdk.GetServerTime()
	sig := bksdk.generateSignature(ts, "POST", "/"+targetUrl.Path, jsonPayload)

	return bksdk.req.Post(targetUrl.String()).
		Set("X-BTK-TIMESTAMP", ts).
		Set("X-BTK-APIKEY", bksdk.apiKey).
		Set("X-BTK-SIGN", sig).
		Set("Content-Type", "application/json").
		Send(jsonPayload)
}

// Endpoint : /api/v3/market/my-open-orders
// Method : GET
// Query : sym string The symbol (e.g. btc_thb)
func (bksdk *Bitkubsdk) MyOpenOrder(sym string) (response.MyOpenOrder, error) {
	var dataResp response.MyOpenOrder

	// join the targetUrl with host and path
	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOpenOrderV3)
	// init query values
	queryValues := url.Values{}
	queryValues.Add("sym", sym)

	resp, body, errs := bksdk.authGet(targetUrl, queryValues).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return dataResp, errs[0]
	}

	err := json.Unmarshal([]byte(body), &dataResp)
	if err != nil {
		return dataResp, err
	}

	if dataResp.Error != 0 {
		return dataResp, errors.New(bkError.ErrorText(dataResp.Error))
	}

	return dataResp, nil
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
func (bksdk *Bitkubsdk) MyOrderHistory(sym string, page, limit, start, end int) (response.MyOrderHistory, error) {
	var dataResp response.MyOrderHistory

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
		return dataResp, errs[0]
	}

	err := json.Unmarshal([]byte(body), &dataResp)
	if err != nil {
		return dataResp, err
	}

	return dataResp, nil
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
func OrderInfo(sym, orderId, side, hash string) (response.OrderInfo, error) {
	return response.OrderInfo{}, nil
}
