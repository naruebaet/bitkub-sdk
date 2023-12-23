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

// MyOpenOrder retrieves a list of all open orders for a given symbol.
// Endpoint: /api/v3/market/my-open-orders
// Method: GET
// Parameters:
// - sym string: The symbol (e.g. btc_thb)
// Returns:
// - response.MyOpenOrder: The response body containing the list of open orders
// - error: Any error that occurred during the API call
func (bksdk *SDK) MyOpenOrder(sym string) (response.MyOpenOrder, error) {
	var respBody response.MyOpenOrder

	// Join the targetUrl with the host and path
	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOpenOrderV3)
	// Initialize query values
	queryValues := url.Values{}
	queryValues.Add("sym", sym)

	// Make the authenticated GET request
	resp, body, errs := bksdk.authGet(targetUrl, queryValues).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// MyOrderHistory retrieves a list of orders that have already matched.
//
// Endpoint: /api/v3/market/my-order-history
// Method: GET
//
// Parameters:
//   - sym: string (required) - The symbol (e.g. btc_thb)
//   - page: int (optional) - Page number
//   - limit: int (optional) - Limit per page
//   - start: int (optional) - Start timestamp
//   - end: int (optional) - End timestamp
//
// Returns:
//   - response.MyOrderHistory: The response body containing the order history
//   - error: An error if the request fails
func (bksdk *SDK) MyOrderHistory(sym string, page, limit, start, end int) (response.MyOrderHistory, error) {
	var respBody response.MyOrderHistory

	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOrderHistoryV3)
	queVal := url.Values{}
	queVal.Add("sym", sym)

	// Add optional parameters to the query string
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

// OrderInfo returns information regarding the specified order.
// It makes a GET request to the endpoint /api/v3/market/order-info.
// The function takes three parameters:
// - sym: the symbol of the order (e.g. btc_thb)
// - orderId: the ID of the order
// - side: the side of the order: buy or sell
// It returns an instance of response.OrderInfo and an error, if any.
func (bksdk *SDK) OrderInfo(sym, orderId, side string) (response.OrderInfo, error) {
	var respBody response.OrderInfo

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketOrderInfoV3)

	// Construct the query parameters
	queVal := url.Values{}
	queVal.Add("sym", sym)
	queVal.Add("id", orderId)
	queVal.Add("sd", side)

	// Make the GET request and get the response
	resp, body, errs := bksdk.authGet(targetUrl, queVal).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Unmarshal the response body into respBody
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Check if there is any error in the response
	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// OrderInfoByHash retrieves information about the specified order.
// It sends a GET request to the /api/v3/market/order-info endpoint.
// The order is identified by its hash.
//
// Parameters:
//   - hash: The hash of the order to retrieve information about.
//
// Returns:
//   - response.OrderInfo: The information about the order.
//   - error: An error if the request fails or if there is an issue with parsing the response.
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

// TradingCredit retrieves the trading credit balance.
func (bksdk *SDK) TradingCredit() (response.TradingCredit, error) {
	var respBody response.TradingCredit

	// Construct the target URL for the API endpoint
	targetUrl := bksdk.apiHost.JoinPath(api.UserTradingCreditsV3)

	// Make a POST request to the API endpoint
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Check if the response contains an error
	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the trading credit balance and no error
	return respBody, nil
}

// Limits checks deposit/withdraw limitations and usage.
func (bksdk *SDK) Limits() (response.Limits, error) {
	var respBody response.Limits

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.UserLimitsV3)

	// Send a POST request to the target URL
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Unmarshal the response body into the respBody struct
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Check if there is an error in the response
	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Wallet retrieves the user's available balances, including both available and reserved balances.
//
// Endpoint:  /api/v3/market/wallet
// Method: POST
// Parameters: N/A
//
// Returns:
// - response.Wallet: the user's wallet information
// - error: any error that occurred during the API call
func (bksdk *SDK) Wallet() (response.Wallet, error) {
	var respBody response.Wallet

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketWalletV3)

	// Make the API call
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Parse the API response
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Balance returns the balances information, including both available and reserved balances.
//
// Endpoint:  /api/v3/market/balances
// Method: POST
// Parameter: N/A
func (bksdk *SDK) Balance() (response.Balance, error) {
	var respBody response.Balance

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketBalancesV3)

	// Send the authenticated POST request
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Unmarshal the response body into the respBody struct
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// WsToken retrieves the token for websocket authentication.
//
// Endpoint:  /api/v3/market/wstoken
// Method: POST
// Parameter: N/A
//
// Returns:
// - response.WsToken: the response body containing the token.
// - error: if there was an error making the request or parsing the response.
func (bksdk *SDK) WsToken() (response.WsToken, error) {
	var respBody response.WsToken

	// Build the target URL for the API endpoint
	targetUrl := bksdk.apiHost.JoinPath(api.MarketWstokenV3)

	// Make the POST request to the API endpoint
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return respBody, errs[0]
	}

	// Parse the response body
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
