package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/naruebaet/bitkub-sdk/bksdk"
	"github.com/naruebaet/bitkub-sdk/bksdk/bkerr"
	"github.com/naruebaet/bitkub-sdk/bksdk/response"
	"github.com/stretchr/testify/assert"
)

// TestGetBids is a unit test function for the GetBids function.
func TestGetBids(t *testing.T) {
	// Define the input arguments type.
	type args struct {
		symbol string // The symbol to get bids data.
		limit  int    // The limit of bids data to retrieve.
	}

	// Define the test cases.
	tests := []struct {
		name    string      // The name of the test case.
		args    args        // The input arguments.
		want    interface{} // The expected output.
		wantErr error       // The expected error.
	}{
		{
			name: "should return 5 bids data of BTC",
			args: args{"THB_BTC", 1},
			want: 5,
		},
		{
			name: "should return MarketBids type",
			args: args{"THB_BTC", 1},
			want: response.MarketBids{},
		},
		{
			name:    "should error when not found symbol",
			args:    args{"BTC_THB", 1},
			wantErr: fmt.Errorf(bkerr.ErrorText(bkerr.InvalidSymbol)),
		},
	}

	// Create a new instance of the SDK.
	sdk := bksdk.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	// Iterate over the test cases.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the GetBids function.
			got, err := sdk.GetBids(tt.args.symbol, tt.args.limit)

			// Perform assertions based on the test case name.
			switch tt.name {
			case "should return 5 bids data of BTC":
				assert.Equal(t, tt.want, len(got.Result[0]))
			case "should return MarketBids type":
				assert.IsType(t, tt.want, got)
			}

			// Assert that the error matches the expected error.
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
