package bksdk

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/naruebaet/bitkub-sdk/bksdk/api"
	"github.com/naruebaet/bitkub-sdk/bksdk/bkerr"
	"github.com/naruebaet/bitkub-sdk/bksdk/request"
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
	// Initialize an empty variable to store the response body
	var respBody response.MyOpenOrder

	// Join the target URL with the host and path
	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOpenOrderV3)

	// Initialize query values
	queryValues := url.Values{}
	queryValues.Add("sym", sym)

	// Make the authenticated GET request
	resp, body, errs := bksdk.authGet(targetUrl, queryValues).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if there was an error in the API response
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
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
	// Initialize the response body
	var respBody response.MyOrderHistory

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOrderHistoryV3)

	// Create the query string parameters
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

	// Make the authenticated GET request
	resp, body, errs := bksdk.authGet(targetUrl, queVal).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Return the response body and nil error
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
	// Initialize the response body
	var respBody response.OrderInfo

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketOrderInfoV3)

	// Construct the query parameters
	queryValues := url.Values{}
	queryValues.Add("sym", sym)
	queryValues.Add("id", orderId)
	queryValues.Add("sd", side)

	// Make the GET request and get the response
	resp, body, errs := bksdk.authGet(targetURL, queryValues).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
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

	// Return the response body and nil error
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

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketOrderInfoV3)

	// Set the query parameters
	queVal := url.Values{}
	queVal.Add("hash", hash)

	// Send the GET request to the target URL with the query parameters
	resp, body, errs := bksdk.authGet(targetUrl, queVal).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if the response status code is not OK
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body and store it in respBody
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Check if there is an error in the response body
	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and nil error
	return respBody, nil
}

// TradingCredit retrieves the trading credit balance.
func (bksdk *SDK) TradingCredit() (response.TradingCredit, error) {
	// Create a variable to store the response body
	var respBody response.TradingCredit

	// Construct the target URL for the API endpoint
	targetUrl := bksdk.apiHost.JoinPath(api.UserTradingCreditsV3)

	// Make a POST request to the API endpoint
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
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
	targetURL := bksdk.apiHost.JoinPath(api.UserLimitsV3)

	// Send a POST request to the target URL
	resp, body, errs := bksdk.authPost(targetURL, "").End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
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
// It makes a POST request to the /api/v3/market/wallet endpoint.
// It returns the user's wallet information and any error that occurred during the API call.
func (bksdk *SDK) Wallet() (response.Wallet, error) {
	var respBody response.Wallet

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketWalletV3)

	// Make the API call
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
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
	// Initialize the response object
	var respBody response.Balance

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketBalancesV3)

	// Send the authenticated POST request
	resp, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if the request was successful
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Unmarshal the response body into the respBody struct
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Return the response object
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
	// Create a variable to store the response body
	var respBody response.WsToken

	// Build the target URL for the API endpoint
	targetURL := bksdk.apiHost.JoinPath(api.MarketWstokenV3)

	// Make the POST request to the API endpoint
	resp, body, errs := bksdk.authPost(targetURL, "").End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check if the response status code is not OK
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Return the response body and no error
	return respBody, nil
}

// InternalWithdraw makes a withdrawal to an internal address.
// The destination address does not need to be a trusted address.
// This API is not enabled by default. Only KYB users can request this feature by contacting us via support@bitkub.com.
func (bksdk *SDK) InternalWithdraw(currency string, address string, memo string, amount float64) (response.InternalWithdraw, error) {

	var respBody response.InternalWithdraw

	// Create request body
	reqBody := map[string]interface{}{
		"cur": currency,
		"amt": amount,
		"adr": address,
		"mem": memo,
	}

	// Marshal request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return response.InternalWithdraw{}, err
	}

	// Construct target URL
	targetUrl := bksdk.apiHost.JoinPath(api.CryptoInternalWithdrawV3)

	// Send authenticated POST request with request body
	resp, body, errs := bksdk.authPost(targetUrl, string(reqBodyByte)).End()
	if errs != nil {
		return response.InternalWithdraw{}, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Unmarshal response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return response.InternalWithdraw{}, err
	}

	return respBody, nil
}

// DepositHistory returns a list of crypto deposit history.
// It sends a POST request to the /api/v3/crypto/deposit-history endpoint.
// Parameters:
// - page: int, optional page number
// - limit: int, optional limit per page
// Returns:
// - response.DepositHistory: the deposit history response
// - error: any error that occurred during the request
func (bksdk *SDK) DepositHistory(page, limit int) (response.DepositHistory, error) {
	var respBody response.DepositHistory

	// Prepare the request body
	reqBody := map[string]interface{}{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	targetUrl := bksdk.apiHost.JoinPath(api.CryptoDepositHistoryV3)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetUrl, string(reqBodyByte)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// WithdrawHistory returns a list of crypto withdraw history.
// It sends a POST request to the /api/v3/crypto/withdraw-history endpoint.
// Parameters:
// - page: int, optional page number
// - limit: int, optional limit per page
// Returns:
// - response.WithdrawHistory: the withdraw history response
// - error: any error that occurred during the request
func (bksdk *SDK) WithdrawHistory(page, limit int) (response.WithdrawHistory, error) {
	// Initialize the response variable
	var respBody response.WithdrawHistory

	// Prepare the request body
	reqBody := map[string]interface{}{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.CryptoWithdrawHistoryV3)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetUrl, string(reqBodyByte)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// PlaceBid creates a buy order by sending a POST request to the /api/v3/market/place-bid endpoint.
// It takes the following parameters:
// - sym: string - The symbol you want to trade (e.g. btc_thb).
// - amt: float64 - Amount you want to spend with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok).
// - rat: float64 - Rate you want for the order with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok).
// - typ: string - Order type: limit or market (for market order, please specify rat as 0).
// - client_id: string - Your id for reference (not required).
func (bksdk *SDK) PlaceBid(sym string, amt, rat float64, typ, client_id string) (response.PlaceBid, error) {
	// Initialize the response variable
	var respBody response.PlaceBid

	// Create the request body
	reqBody := request.PlaceBid{
		Symbol:   sym,
		Amount:   amt,
		Rate:     rat,
		Type:     typ,
		ClientID: client_id,
	}

	// Validate the request body
	err := reqBody.Validate()
	if err != nil {
		return respBody, err
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketPlaceBidV3)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetURL, string(reqBodyByte)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// PlaceAsk creates a sell order.
// Endpoint: /api/v3/market/place-ask
// Method: POST
// Parameters:
// - sym: string - The symbol you want to trade (e.g. btc_thb).
// - amt: float64 - Amount you want to spend with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok).
// - rat: float64 - Rate you want for the order with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok).
// - typ: string - Order type: limit or market (for market order, please specify rat as 0).
// - client_id: string - Your id for reference (not required).
// It returns the response body and an error (if any).
func (bksdk *SDK) PlaceAsk(sym string, amt, rat float64, typ, client_id string) (response.PlaceAsk, error) {
	// Initialize the response variable
	var respBody response.PlaceAsk

	// Create the request body
	reqBody := request.PlaceAsk{
		Symbol:   sym,
		Amount:   amt,
		Rate:     rat,
		Type:     typ,
		ClientID: client_id,
	}

	// Validate the request body
	err := reqBody.Validate()
	if err != nil {
		return respBody, err
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketPlaceAskV3)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetURL, string(reqBodyByte)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint: /api/v3/market/cancel-order
// Method: POST
// Desc: Cancel an open order
// Parameters:
// - sym string The symbol. Please note that the current endpoint requires the symbol thb_btc. However, it will be changed to btc_thb soon and you will need to update the configurations accordingly for uninterrupted API functionality.
// - id string Order id you wish to cancel
// - sd string Order side: buy or sell
// - hash string Cancel an order with order hash (optional). You don't need to specify sym, id, and sd when you specify order hash.
func (bksdk *SDK) CancelOrder(sym, id, sd, hash string) (response.CancelOrder, error) {
	// Initialize the response variable
	var respBody response.CancelOrder

	// Create the request body
	reqBody := request.CancelOrder{
		Symbol: sym,
		ID:     id,
		Side:   sd,
		Hash:   hash,
	}

	// Validate the request body
	err := reqBody.Validate()
	if err != nil {
		return respBody, err
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketCancelOrderV3)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Addresses is a function that lists all crypto addresses.
// It sends a POST request to the /api/v3/crypto/addresses endpoint
// with optional pagination parameters (page and limit).
// It returns a response containing the crypto addresses and an error, if any.
func (bksdk *SDK) Addresses(page, limit int) (response.CryptoAddresses, error) {
	// Initialize the response variable
	var respBody response.CryptoAddresses

	// Create the request body
	reqBody := map[string]int{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.CryptoAddresses)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint: /api/v3/crypto/generate-address
// Method: POST
// Desc: Generate a new crypto address (will replace existing adresses; previous address can still be used to received funds)
// Parameters:
// - sym string Symbol (e.g. THB_BTC, THB_ETH, etc.)
func (bksdk *SDK) GenerateAddress(symbol string) (response.CryptoGenerateAddress, error) {
	// Initialize the response variable
	var respBody response.CryptoGenerateAddress

	// symbol to upper case

	// create the request body
	reqBody := map[string]string{
		"sym": strings.ToUpper(symbol),
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.CryptoGenerateAddress)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}

// Endpoint: /api/v3/crypto/withdraw
// Method: POST
// Desc: Make a withdrawal to a trusted address.
// Parameters:
// - cur string Currency for withdrawal (e.g. BTC, ETH)
// - amt float Amount you want to withdraw
// - adr string Address to which you want to withdraw
// - mem string (Optional) Memo or destination tag to which you want to withdraw
// - net string Cryptocurrency network to withdraw
// No default value of this field. Please find the available network from the link as follows. https://www.bitkub.com/fee/cryptocurrency
// For example ETH refers to ERC-20.
// For request on ERC-20, please assign the net value as ETH.
// For request on BEP-20, please assign the net value as BSC.
// For request on KAP-20, please assign the net value as BKC.
func (bksdk *SDK) Withdraw(currency string, address string, memo string, amount float64, network string) (response.CryptoWithdraw, error) {
	// Initialize the response variable
	var respBody response.CryptoWithdraw

	// Create the request body
	reqBody := map[string]interface{}{
		"cur": currency,
		"amt": amount,
		"adr": address,
		"mem": memo,
		"net": network,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.CryptoWithdraw)

	// Send the authenticated POST request with the request body
	resp, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return respBody, errors.New(body)
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	return respBody, nil
}
