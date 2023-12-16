package bitkubsdk

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/naruebaet/bitkub-sdk/api"
	"github.com/naruebaet/bitkub-sdk/response"
)

// Endpoint : /api/status
// Method : GET
func (bksdk *Bitkubsdk) GetStatus() (response.Status, error) {
	var dataResp response.Status

	url := bksdk.apiHost.JoinPath(api.Status)

	resp, body, errs := bksdk.req.Get(url.String()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		bksdk.log.Error(errs[0].Error())
		return dataResp, errors.New(api.Status + " Internal server error!")
	}

	err := json.Unmarshal([]byte(body), &dataResp)
	if err != nil {
		bksdk.log.Error(api.Status + " Can't unmarshal the response body!")
		return dataResp, errors.New(api.Status + " Internal server error!")
	}

	return dataResp, nil
}

func (bksdk *Bitkubsdk) GetServerTime() (string, error) {

	url := bksdk.apiHost.JoinPath(api.ServertimeV3)

	resp, timestamp, errs := bksdk.req.Get(url.String()).End()
	if errs != nil && resp.StatusCode != http.StatusOK {
		bksdk.log.Error(errs[0].Error())
		return "0", errors.New(api.Status + " Internal server error!")
	}

	return timestamp, nil
}
