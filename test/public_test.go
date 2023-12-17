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

// GetBids real request test.
func TestGetBids(t *testing.T) {
	// test arguments
	type args struct {
		sym string
		lmt int
	}

	// test case definition
	tests := []struct {
		id      int
		name    string
		args    args
		want    any
		wantErr error
	}{
		{
			id:   1,
			name: "should return 5 bids data of BTC",
			args: args{"THB_BTC", 1},
			want: 5,
		},
		{
			id:   2,
			name: "should return MarketBids type",
			args: args{"THB_BTC", 1},
			want: response.MarketBids{},
		},
		{
			id:      3,
			name:    "should error when not found symbol",
			args:    args{"BTC_THB", 1},
			wantErr: fmt.Errorf(bkerr.ErrorText(bkerr.InvalidSymbol)),
		},
	}

	// init sdk for test
	sdk := bksdk.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sdk.GetBids(tt.args.sym, tt.args.lmt)

			if tt.id == 1 {
				assert.Equal(t, tt.want, len(got.Result[0]))
			} else if tt.id == 2 {
				assert.IsType(t, tt.want, got)
			}

			assert.Equal(t, tt.wantErr, err)
		})
	}

}
