package request

import (
	"errors"
	"strings"

	validate "github.com/go-playground/validator/v10"
)

var resolutions = map[string]string{
	"1":   "1",
	"5":   "5",
	"15":  "15",
	"60":  "60",
	"240": "240",
	"1D":  "1D",
}

// ValidateResolution checks if the given value is a valid resolution and returns the corresponding resolution value.
// If the value is not valid, it returns an error.
func ValidateResolution(value string) (string, error) {
	resolution := resolutions[strings.ToUpper(value)]
	if resolution == "" {
		return "", errors.New("Invalid resolution")
	}

	return resolution, nil
}

type PlaceBid struct {
	Symbol   string  `json:"sym"`
	Amount   float64 `json:"amt"`
	Rate     float64 `json:"rat"`
	Type     string  `json:"typ" validate:"oneof=limit market"`
	ClientID string  `json:"client_id"`
}

// validate type of bid
func (p *PlaceBid) Validate() error {
	validate := validate.New()
	return validate.Struct(p)
}

type PlaceAsk struct {
	Symbol   string  `json:"sym"`
	Amount   float64 `json:"amt"`
	Rate     float64 `json:"rat"`
	Type     string  `json:"typ" validate:"oneof=limit market"`
	ClientID string  `json:"client_id"`
}

func (p *PlaceAsk) Validate() error {
	validate := validate.New()
	return validate.Struct(p)
}

type CancelOrder struct {
	Symbol string `json:"sym"`
	ID     string `json:"id"`
	Side   string `json:"sd" oneof:"buy sell"`
	Hash   string `json:"hash"`
}

func (p *CancelOrder) Validate() error {
	validate := validate.New()
	return validate.Struct(p)
}
