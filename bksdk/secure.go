package bksdk

import (
	"encoding/json"
	"errors"
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
func (bksdk *SDK) MyOpenOrder(sym string) ([]response.MyOpenOrderResult, error) {
	// Initialize an empty variable to store the response body
	var respBody response.MyOpenOrder

	// Join the target URL with the host and path
	targetUrl := bksdk.apiHost.JoinPath(api.MarketMyOpenOrderV3)

	// Initialize query values
	queryValues := url.Values{}
	queryValues.Add("sym", sym)

	// Make the authenticated GET request
	_, body, errs := bksdk.authGet(targetUrl, queryValues).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
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
func (bksdk *SDK) MyOrderHistory(sym string, page, limit, start, end int) ([]response.MyOrderHistoryResult, response.BKPaginate, error) {
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
	_, body, errs := bksdk.authGet(targetUrl, queVal).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and nil error
	return respBody.Result, respBody.Pagination, nil
}

// OrderInfo returns information regarding the specified order.
// It makes a GET request to the endpoint /api/v3/market/order-info.
// The function takes three parameters:
// - sym: the symbol of the order (e.g. btc_thb)
// - orderId: the ID of the order
// - side: the side of the order: buy or sell
// It returns an instance of response.OrderInfo and an error, if any.
func (bksdk *SDK) OrderInfo(sym, orderId, side string) (response.OrderInfoResult, error) {
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
	_, body, errs := bksdk.authGet(targetURL, queryValues).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Unmarshal the response body into respBody
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if there is any error in the response
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and nil error
	return respBody.Result, nil
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
func (bksdk *SDK) OrderInfoByHash(hash string) (response.OrderInfoResult, error) {
	var respBody response.OrderInfo

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketOrderInfoV3)

	// Set the query parameters
	queVal := url.Values{}
	queVal.Add("hash", hash)

	// Send the GET request to the target URL with the query parameters
	_, body, errs := bksdk.authGet(targetUrl, queVal).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the response body and store it in respBody
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if there is an error in the response body
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and nil error
	return respBody.Result, nil
}

// TradingCredit retrieves the trading credit balance.
func (bksdk *SDK) TradingCredit() (float64, error) {
	// Create a variable to store the response body
	var respBody response.TradingCredit

	// Construct the target URL for the API endpoint
	targetUrl := bksdk.apiHost.JoinPath(api.UserTradingCreditsV3)

	// Make a POST request to the API endpoint
	_, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Unmarshal the response body into the respBody variable
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the trading credit balance and no error
	return respBody.Result, nil
}

// Limits checks deposit/withdraw limitations and usage.
func (bksdk *SDK) Limits() (response.LimitsResult, error) {
	var respBody response.Limits

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.UserLimitsV3)

	// Send a POST request to the target URL
	_, body, errs := bksdk.authPost(targetURL, "").End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Unmarshal the response body into the respBody struct
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if there is an error in the response
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// Wallet retrieves the user's available balances, including both available and reserved balances.
// It makes a POST request to the /api/v3/market/wallet endpoint.
// It returns the user's wallet information and any error that occurred during the API call.
func (bksdk *SDK) Wallet() (response.WalletResult, error) {
	var respBody response.Wallet

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketWalletV3)

	// Make the API call
	_, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the API response
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// Balance returns the balances information, including both available and reserved balances.
//
// Endpoint:  /api/v3/market/balances
// Method: POST
// Parameter: N/A
func (bksdk *SDK) Balances() (response.BalanceResult, error) {
	// Initialize the response object
	var respBody response.Balances

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.MarketBalancesV3)

	// Send the authenticated POST request
	_, body, errs := bksdk.authPost(targetUrl, "").End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Unmarshal the response body into the respBody struct
	err := json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response object
	return respBody.Result, nil
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
func (bksdk *SDK) WsToken() (token string, err error) {
	// Create a variable to store the response body
	var respBody response.WsToken

	// Build the target URL for the API endpoint
	targetURL := bksdk.apiHost.JoinPath(api.MarketWstokenV3)

	// Make the POST request to the API endpoint
	_, body, errs := bksdk.authPost(targetURL, "").End()
	if errs != nil {
		return "", errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return "", err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return "", errors.New(bkerr.ErrorText(respBody.Error))
	}

	// Return the response body and no error
	return respBody.Result, nil
}

// InternalWithdraw makes a withdrawal to an internal address.
// The destination address does not need to be a trusted address.
// This API is not enabled by default. Only KYB users can request this feature by contacting us via support@bitkub.com.
func (bksdk *SDK) CryptoInternalWithdraw(currency string, address string, memo string, amount float64) (response.InternalWithdrawResult, error) {

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
		return respBody.Result, err
	}

	// Construct target URL
	targetUrl := bksdk.apiHost.JoinPath(api.CryptoInternalWithdrawV3)

	// Send authenticated POST request with request body
	_, body, errs := bksdk.authPost(targetUrl, string(reqBodyByte)).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Unmarshal response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// DepositHistory returns a list of crypto deposit history.
// It sends a POST request to the /api/v3/crypto/deposit-history endpoint.
// Parameters:
// - page: int, optional page number
// - limit: int, optional limit per page
// Returns:
// - response.DepositHistory: the deposit history response
// - error: any error that occurred during the request
func (bksdk *SDK) CryptoDepositHistory(page, limit int) ([]response.DepositHistoryResult, response.BKPaginate, error) {
	var respBody response.DepositHistory

	// Prepare the request body
	reqBody := map[string]interface{}{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	targetUrl := bksdk.apiHost.JoinPath(api.CryptoDepositHistoryV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetUrl, string(reqBodyByte)).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, respBody.Pagination, nil
}

// WithdrawHistory returns a list of crypto withdraw history.
// It sends a POST request to the /api/v3/crypto/withdraw-history endpoint.
// Parameters:
// - page: int, optional page number
// - limit: int, optional limit per page
// Returns:
// - response.WithdrawHistory: the withdraw history response
// - error: any error that occurred during the request
func (bksdk *SDK) CryptoWithdrawHistory(page, limit int) ([]response.WithdrawHistoryResult, response.BKPaginate, error) {
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
		return respBody.Result, respBody.Pagination, err
	}

	// Construct the target URL
	targetUrl := bksdk.apiHost.JoinPath(api.CryptoWithdrawHistoryV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetUrl, string(reqBodyByte)).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, respBody.Pagination, nil
}

// PlaceBid creates a buy order by sending a POST request to the /api/v3/market/place-bid endpoint.
// It takes the following parameters:
// - sym: string - The symbol you want to trade (e.g. btc_thb).
// - amt: float64 - Amount you want to spend with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok).
// - rat: float64 - Rate you want for the order with no trailing zero (e.g. 1000.00 is invalid, 1000 is ok).
// - typ: string - Order type: limit or market (for market order, please specify rat as 0).
// - client_id: string - Your id for reference (not required).
func (bksdk *SDK) PlaceBid(sym string, amt, rat float64, typ, client_id string) (response.PlaceBidResult, error) {
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
		return respBody.Result, err
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketPlaceBidV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(reqBodyByte)).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
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
func (bksdk *SDK) PlaceAsk(sym string, amt, rat float64, typ, client_id string) (response.PlaceAskResult, error) {
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
		return respBody.Result, err
	}

	// Convert the request body to JSON
	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.MarketPlaceAskV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(reqBodyByte)).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
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
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody, nil
}

// Addresses is a function that lists all crypto addresses.
// It sends a POST request to the /api/v3/crypto/addresses endpoint
// with optional pagination parameters (page and limit).
// It returns a response containing the crypto addresses and an error, if any.
func (bksdk *SDK) CryptoAddresses(page, limit int) ([]response.CryptoAddressesResult, response.BKPaginate, error) {
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
		return respBody.Result, respBody.Pagination, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.CryptoAddressesV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, respBody.Pagination, nil
}

// GenerateAddress generates a new crypto address for the given symbol.
// The previous address can still be used to receive funds.
// It makes a POST request to the /api/v3/crypto/generate-address endpoint.
//
// Parameters:
// - symbol: The symbol for which to generate the address (e.g. THB_BTC, THB_ETH, etc.)
//
// Returns:
// - response: The generated address and other relevant information.
// - error: An error if the request fails or if there is an issue parsing the response.
func (bksdk *SDK) CryptoGenerateAddress(symbol string) ([]response.CryptoGenerateAddressResult, error) {
	// Initialize the response variable
	var respBody response.CryptoGenerateAddress

	// Convert the symbol to upper case
	symbol = strings.ToUpper(symbol)

	// Create the request body
	reqBody := map[string]string{
		"sym": symbol,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.CryptoGenerateAddressV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// Withdraw makes a withdrawal to a trusted address.
// Endpoint: /api/v3/crypto/withdraw
// Method: POST
// Parameters:
//   - currency: Currency for withdrawal (e.g. BTC, ETH)
//   - amount: Amount you want to withdraw
//   - address: Address to which you want to withdraw
//   - memo: (Optional) Memo or destination tag to which you want to withdraw
//   - network: Cryptocurrency network to withdraw
//     No default value of this field. Please find the available network from the link as follows. https://www.bitkub.com/fee/cryptocurrency
//     For example ETH refers to ERC-20.
//     For request on ERC-20, please assign the network value as ETH.
//     For request on BEP-20, please assign the network value as BSC.
//     For request on KAP-20, please assign the network value as BKC.
//
// Returns:
// - response: The response body with the withdrawal details
// - error: An error if the withdrawal request fails
func (bksdk *SDK) CryptoWithdraw(currency string, address string, memo string, amount float64, network string) (response.CryptoWithdrawResult, error) {
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
	if err != nil {
		return respBody.Result, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.CryptoWithdrawV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// Accounts lists all approved bank accounts.
//
// Parameters:
// - page: int, Page number (optional)
// - limit: int, Limit number of results (optional)
//
// Returns:
// - response.FiatAccounts: List of approved bank accounts
// - error: Error if any occurs
func (bksdk *SDK) FiatAccounts(page int, limit int) ([]response.FiatAccountsResult, response.BKPaginate, error) {
	// Initialize the response variable
	var respBody response.FiatAccounts

	// Create the request body
	reqBody := map[string]interface{}{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.FiatAccountsV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, respBody.Pagination, nil
}

// FiatWithdraw makes a withdrawal to an approved bank account.
// It sends a POST request to the /api/v3/fiat/withdraw endpoint.
// Parameters:
// - id: string, the bank account id
// - amt: float64, the amount to withdraw
// Returns:
// - response.FiatWithdraw: the response body
// - error: any error that occurred during the request
func (bksdk *SDK) FiatWithdraw(id string, amt float64) (response.FiatWithdrawResult, error) {
	// Initialize the response variable
	var respBody response.FiatWithdraw

	// Create the request body
	reqBody := map[string]interface{}{
		"id":  id,
		"amt": amt,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.FiatWithdrawV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, nil
}

// Endpoint: /api/v3/fiat/deposit-history
// Method: POST
// Desc: List fiat deposit history
// Parameters:
// - p int Page (optional)
// - lmt int Limit (optional)
// Function: FiatDepositHistory retrieves the fiat deposit history using the specified page and limit.
// It returns a response object containing the deposit history and an error if any.
func (bksdk *SDK) FiatDepositHistory(page, limit int) ([]response.FiatDepositHistoryResult, response.BKPaginate, error) {
	// Initialize the response variable
	var respBody response.FiatDepositHistory

	// Create the request body
	reqBody := map[string]interface{}{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.FiatDepositHistoryV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, respBody.Pagination, nil
}

// FiatWithdrawHistory retrieves the list of fiat withdrawal history.
//
// Parameters:
// - page: Page number (optional)
// - limit: Limit number (optional)
//
// Returns:
// - response.FiatWithdrawHistory: The response body containing the fiat withdrawal history
// - error: Any error that occurred during the API request
func (bksdk *SDK) FiatWithdrawHistory(page, limit int) ([]response.FiatWithdrawHistoryResult, response.BKPaginate, error) {
	// Initialize the response variable
	var respBody response.FiatWithdrawHistory

	// Create the request body
	reqBody := map[string]interface{}{
		"p":   page,
		"lmt": limit,
	}

	// Convert the request body to JSON
	jsonReqBody, err := json.Marshal(reqBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Construct the target URL
	targetURL := bksdk.apiHost.JoinPath(api.FiatWithdrawHistoryV3)

	// Send the authenticated POST request with the request body
	_, body, errs := bksdk.authPost(targetURL, string(jsonReqBody)).End()
	if errs != nil {
		return respBody.Result, respBody.Pagination, errs[0]
	}

	// Parse the response body
	err = json.Unmarshal([]byte(body), &respBody)
	if err != nil {
		return respBody.Result, respBody.Pagination, err
	}

	// Check if the response body contains an error
	if respBody.Error != 0 {
		return respBody.Result, respBody.Pagination, errors.New(bkerr.ErrorText(respBody.Error))
	}

	return respBody.Result, respBody.Pagination, nil
}
