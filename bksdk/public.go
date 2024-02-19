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

// GetStatus retrieves the status from the API.
func (bksdk *SDK) GetStatus() (response.Status, error) {
	// Initialize the response body.
	var respBody response.Status

	// Generate the target URL.
	targetURL := bksdk.apiHost.JoinPath(api.Status)

	// Send a GET request to the target URL and retrieve the response body.
	_, body, errs := bksdk.req.Get(targetURL.String()).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Unmarshal the response body into the respBody variable.
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Return the response body and any error that occurred.
	return respBody, nil
}

// GetServerTime retrieves the server time from the API.
// Endpoint: /api/servertime
// Method: GET
func (bksdk *SDK) GetServerTime() (string, error) {

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.ServertimeV3)

	// Send a GET request to the target URL
	resp, timestamp, errs := bksdk.req.Get(targetUrl.String()).End()

	// Check for errors or a non-OK status code
	if errs != nil || resp.StatusCode != http.StatusOK {
		return "0", errs[0]
	}

	// Return the server time
	return timestamp, nil
}

// GetSymbols retrieves the market symbols from the API.
// Endpoint: /api/market/symbols
// Method: GET
func (bksdk *SDK) GetSymbols() ([]response.MarketSymbolsResult, error) {
	// Initialize the response body
	var respBody response.MarketSymbols

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketSymbol)

	// Send the HTTP GET request
	resp, body, errs := bksdk.req.Get(targetURL.String()).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody.Result, errors.New(body)
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if there is an error in the response body
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// GetTicker retrieves the market ticker for a given symbol.
// It makes a GET request to the "/api/market/ticker" endpoint.
// If a symbol is provided, it is included as a query parameter.
// The response body is unmarshalled into a response.MarketTickerData struct.
// If the request is successful, the response body is returned along with nil error.
// If any error occurs during the request or unmarshalling, the response body is empty
// and the error is returned.
func (bksdk *SDK) GetTicker(sym string) (map[string]response.MarketTickerData, error) {
	// Initialize the response body
	var respBody map[string]response.MarketTickerData

	// Build the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketTicker)

	// Build the query parameters
	queryValues := url.Values{}
	if sym != "" {
		queryValues.Add("sym", sym)
	}

	// Make the GET request
	resp, body, errs := bksdk.req.Get(targetURL.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Unmarshal the response body into the respBody struct
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// GetTrade retrieves recent trades from the API.
// It makes a GET request to the /api/market/trades endpoint.
// Parameters:
// - sym: The symbol (e.g. thb_btc)
// - limit: No. of limit to query recent trades
// Returns:
// - response.MarketTrades: The response body containing the recent trades
// - error: An error if the request fails or if the response body cannot be parsed
func (bksdk *SDK) GetTrade(sym string, limit int) (response.MarketTradesResult, error) {
	// Initialize the response body
	var respBody response.MarketTrades

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketTrades)

	// Create the query parameters
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	// Make the GET request
	resp, body, errs := bksdk.req.Get(targetURL.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody.Result, errors.New(body)
	}

	// Parse the response body
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check for any errors in the response body
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and nil error
	return respBody.Result, nil
}

// GetBids queries the /api/market/bids endpoint with the GET method
// and returns the response body containing the market bids.
//
// Parameters:
// - sym: The symbol (e.g. thb_btc)
// - limit: The number of limit to query open buy orders
//
// Returns:
// - response.MarketBids: The response body containing the market bids
// - error: An error if any occurred during the request or response handling
func (bksdk *SDK) GetBids(sym string, limit int) (response.MarketResult, error) {
	// Initialize the response body
	var respBody response.MarketBids

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketBids)

	// Create the query parameters
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	// Send the GET request and handle the response
	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody.Result, errors.New(body)
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the respBody contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and nil error
	return respBody.Result, nil
}

// GetAsks retrieves a list of open sell orders for a given symbol with a specified limit.
// It makes a GET request to the /api/market/asks endpoint.
// Returns the response body as a MarketAsks struct and any error encountered.
func (bksdk *SDK) GetAsks(sym string, limit int) (response.MarketResult, error) {
	// Initialize the response body
	var respBody response.MarketAsks

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketAsks)

	// Create the query parameters
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	// Send the GET request and retrieve the response
	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Check if the response status code is not OK
	if resp.StatusCode != http.StatusOK {
		return respBody.Result, errors.New(body)
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the respBody.Error field is not zero
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and no error
	return respBody.Result, nil
}

// GetBooks retrieves market books for a given symbol and limit
// Endpoint: /api/market/books
// Method: GET
/*
Query Parameters:
	sym string - The symbol (e.g. thb_btc)
	lmt int - Number of limit to query open sell orders
*/
func (bksdk *SDK) GetBooks(sym string, limit int) (response.MarketBooksResult, error) {
	// Initialize response body
	var respBody response.MarketBooks

	// Construct target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketBooks)

	// Set query parameters
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	// Send GET request to the target URL with query parameters
	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return respBody.Result, errors.New(body)
	}

	// Unmarshal response body into respBody
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if respBody contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return respBody and nil error if everything is successful
	return respBody.Result, nil
}

// GetDepth queries the /api/market/books endpoint using the GET method
// to retrieve market depth information.
//
// Parameters:
// - sym: The symbol (e.g. thb_btc)
// - limit: The number of orders to query
//
// Returns:
// - response.MarketDepth: The market depth response
// - error: Any error that occurred during the request
func (bksdk *SDK) GetDepth(sym string, limit int) (response.MarketDepth, error) {
	// Initialize the response body
	var respBody response.MarketDepth

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketDepth)

	// Create a query string with the sym and lmt parameters
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("lmt", strconv.Itoa(limit))

	// Send the request and retrieve the response
	resp, body, errs := bksdk.req.Get(targetUrl.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if the response status code is not OK
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Return the response body and any error that occurred
	return respBody, nil
}

// GetHistory retrieves trading view history for a given symbol
//
// Endpoint: /tradingview/history
// Method: GET
//
// Parameters:
// - symbol: string - The symbol (e.g. BTC_THB)
// - resolution: string - Chart resolution (1, 5, 15, 60, 240, 1D)
// - from: int - Timestamp of the starting time (e.g. 1633424427)
// - to: int - Timestamp of the ending time (e.g. 1633427427)
//
// Returns:
// - response.TradingviewHistory: the trading view history
// - error: if there was an error retrieving the history
func (bksdk *SDK) GetHistory(symbol string, resolution string, from int, to int) (response.TradingviewHistory, error) {
	// Initialize the response body
	var respBody response.TradingviewHistory

	// Validate the resolution
	resl, err := request.ValidateResolution(resolution)
	if err != nil {
		return respBody, err
	}

	// Build the target URL and query parameters
	targetURL := bksdk.apiHost.JoinPath(api.TradingviewHistory)
	queryValues := url.Values{}
	queryValues.Add("sym", symbol)
	queryValues.Add("resolution", resl)
	queryValues.Add("from", strconv.Itoa(from))
	queryValues.Add("to", strconv.Itoa(to))

	// Send the GET request
	resp, body, errs := bksdk.req.Get(targetURL.String() + "?" + queryValues.Encode()).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Unmarshal the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
