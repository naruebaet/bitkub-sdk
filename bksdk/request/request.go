package request

import (
	"errors"
	"strings"
)

var resl = map[string]string{
	"1":   "1",
	"5":   "5",
	"15":  "15",
	"60":  "60",
	"240": "240",
	"1D":  "1D",
}

func ValidateResolution(val string) (string, error) {
	// find val
	res := resl[strings.ToUpper(val)]
	if res == "" {
		return "", errors.New("Invalid resolutions!")
	}

	return res, nil
}
