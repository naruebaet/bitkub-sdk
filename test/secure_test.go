package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/naruebaet/bitkub-sdk/bksdk"
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

	resp, _, _ := sdk.FiatWithdrawHistory(1, 10)
	res, _ := bksdk.PrettyStruct(resp)

	fmt.Println(res)
}

func TestFiatDepositHistory(t *testing.T) {
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")
	sdk := bksdk.New(apiKey, apiSecret)

	resp, _, _ := sdk.FiatDepositHistory(1, 10)
	res, _ := bksdk.PrettyStruct(resp)

	fmt.Println(res)
}
