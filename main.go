package bitkubsdk

import (
	"log/slog"
	"net/url"
	"strings"

	"github.com/parnurzeal/gorequest"
)

type Bitkubsdk struct {
	apiHost   *url.URL
	apiKey    string
	apiSecret string
	req       *gorequest.SuperAgent
	log       *slog.Logger
}

func New(apiHost, apiKey, apiSecret string) (*Bitkubsdk, error) {
	// init. gorequest super agent
	req := gorequest.New()

	// trim the url
	apiHost = strings.Trim(apiHost, " ")

	// validate that is the url
	apiHostUrl, _ := url.Parse(apiHost)

	return &Bitkubsdk{
		apiHost:   apiHostUrl,
		apiKey:    apiKey,
		apiSecret: apiSecret,
		req:       req,
	}, nil
}

func (bksdk *Bitkubsdk) AddSlog(logger *slog.Logger) {
	bksdk.log = logger
}
