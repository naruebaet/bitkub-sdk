package test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/naruebaet/bitkub-sdk/bksdk"
)

func TestTradingCredit(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sdk := bksdk.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	resp, _ := sdk.TradingCredit()

	res, _ := bksdk.PrettyStruct(resp)

	fmt.Println(res)
}

// test FiatWithdrawHistory
func TestFiatWithdrawHistory(t *testing.T) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sdk := bksdk.New(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	resp, _ := sdk.FiatWithdrawHistory(1, 10)

	res, _ := bksdk.PrettyStruct(resp)

	fmt.Println(res)
}
