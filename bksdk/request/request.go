package request

import (
	"errors"
	"strings"
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
