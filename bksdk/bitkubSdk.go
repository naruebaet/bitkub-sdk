package bitkubsdk

import (
	"net/url"

	"github.com/parnurzeal/gorequest"
)

type Bitkubsdk struct {
	apiHost   *url.URL
	apiKey    string
	apiSecret string
	req       *gorequest.SuperAgent
}

func New(apiKey, apiSecret string) (*Bitkubsdk, error) {
	// init. gorequest super agent
	req := gorequest.New()

	// validate that is the url
	apiHostUrl, _ := url.Parse("https://api.bitkub.com")

	return &Bitkubsdk{
		apiHost:   apiHostUrl,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		req:       req,
	}, nil
}
