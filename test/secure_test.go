package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/naruebaet/bitkub-sdk/bksdk"
	"github.com/naruebaet/bitkub-sdk/bksdk/bkerr"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func TestTradingCredit(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	sdk := bksdk.New(apiKey, apiSecret)
	resp, _ := sdk.TradingCredit()

	res, _ := bksdk.PrettyStruct(resp)
	fmt.Println(res)
}

func TestFiatWithdrawHistory(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	sdk := bksdk.New(apiKey, apiSecret)

	resp, _ := sdk.FiatWithdrawHistory(1, 10)
	res, _ := bksdk.PrettyStruct(resp)

	if resp.Error != 0 {
		fmt.Println(bkerr.ErrorText(resp.Error))
	}

	fmt.Println(res)
}
