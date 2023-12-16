package bitkubsdk

import (
	"encoding/json"
	"net/http"

	"github.com/naruebaet/bitkub-sdk/bksdk/api"
	"github.com/naruebaet/bitkub-sdk/bksdk/response"
)

// Endpoint : /api/status
// Method : GET
func (bksdk *Bitkubsdk) GetStatus() (response.Status, error) {
	// init response body
	var respBody response.Status

	url := bksdk.apiHost.JoinPath(api.Status)

	resp, body, errs := bksdk.req.Get(url.String()).End()
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
func (bksdk *Bitkubsdk) GetServerTime() (string, error) {

	url := bksdk.apiHost.JoinPath(api.ServertimeV3)

	resp, timestamp, errs := bksdk.req.Get(url.String()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		return "0", errs[0]
	}
	return timestamp, nil
}
