package bksdk

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

// generateSignature generates a signature pattern for Bitkub API.
// The signature is generated from the timestamp, request method, API path, query parameter,
// and JSON payload using HMAC SHA-256.
// The API Secret is used as the secret key for generating the HMAC variant of the JSON payload.
// The signature is returned in hex format.
// The user must attach the signature via the Request Header.
// To get a new timestamp in milliseconds, use the /api/v3/servertime endpoint. The old one is in seconds.
func (bksdk *SDK) generateSignature(timestamp, method, endpoint, bodyPayload string) (signature string) {
	// Get the secret key for HMAC SHA-256.
	secret := bksdk.apiSecret

	// Initialize the encryptor with the secret key.
	// Create a new HMAC by defining the hash type and the key.
	encryptor := hmac.New(sha256.New, []byte(secret))

	// Compute the HMAC.
	encryptor.Write([]byte(timestamp + method + endpoint + bodyPayload))
	dataHmac := encryptor.Sum(nil)

	// Convert the HMAC to hex format.
	hexHmac := hex.EncodeToString(dataHmac)

	return hexHmac
}

// authGet is a function that wraps a request before sending it to secure endpoints.
// It performs the following steps:
// 1. Retrieves the server time before generating the signature.
// 2. Generates the signature using the timestamp, method, endpoint, and payloads.
// 3. Sets the headers with the generated key and API key.
//
// Parameters:
// - targetUrl: The target URL for the request.
// - queryValues: The query values for the request.
//
// Returns:
// - A *gorequest.SuperAgent object representing the authenticated GET request.
func (bksdk *SDK) authGet(targetUrl *url.URL, queryValues url.Values) *gorequest.SuperAgent {
	// Retrieve the server time
	ts, _ := bksdk.GetServerTime()

	// Generate the signature
	sig := bksdk.generateSignature(ts, "GET", "/"+targetUrl.Path, "?"+queryValues.Encode())

	// Set the headers and query values of the request
	return bksdk.req.Get(targetUrl.String()).
		Set("X-BTK-TIMESTAMP", ts).
		Set("X-BTK-APIKEY", bksdk.apiKey).
		Set("X-BTK-SIGN", sig).
		Set("Content-Type", "application/json").Query(queryValues.Encode())
}

// authPost is a function that wraps a request before sending it to secure endpoints.
// It generates a signature using the server time, method, endpoint, and payload,
// and sets the necessary headers for authentication.
func (bksdk *SDK) authPost(targetUrl *url.URL, jsonPayload string) *gorequest.SuperAgent {
	// Step 1 - Get server time before generating the signature
	ts, _ := bksdk.GetServerTime()

	// Step 2 - Generate the signature with the timestamp, method, endpoint, and payloads
	sig := bksdk.generateSignature(ts, "POST", "/"+targetUrl.Path, jsonPayload)

	// Step 3 - Set the headers with the generated key and API key
	return bksdk.req.Post(targetUrl.String()).
		Set("X-BTK-TIMESTAMP", ts).
		Set("X-BTK-APIKEY", bksdk.apiKey).
		Set("X-BTK-SIGN", sig).
		Set("Content-Type", "application/json").
		Send(jsonPayload)
}

// PrettyStruct prints a pretty JSON representation of a struct.
// It takes in a `data` interface{} parameter and returns a string representation of the JSON.
// If there is an error during the marshaling process, it returns an empty string and the error.
func PrettyStruct(data interface{}) (string, error) {
	// Marshal the data into pretty JSON with 4 spaces indentation.
	val, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return "", err
	}
	return string(val), nil
}
