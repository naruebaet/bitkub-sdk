package bksdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"

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
func (bksdk *SDK) generateSignature(timestamp, method, endpoint, bodyPayload string) (signature string) {
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
func (bksdk *SDK) authGet(targetUrl *url.URL, queryValues url.Values) *gorequest.SuperAgent {
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
func (bksdk *SDK) authPost(targetUrl *url.URL, jsonPayload string) *gorequest.SuperAgent {
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

// PrettyStruct : print pretty json struct
func PrettyStruct(data interface{}) (string, error) {
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
