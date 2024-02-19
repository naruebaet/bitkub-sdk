package main

import (
	"fmt"

	"github.com/naruebaet/bitkub-sdk/bksdk"
)

func main() {
	// Create a new instance of the SDK with the provided API credentials.
	sdk := bksdk.New("xxx", "xxx")

	fmt.Println("===== GetSymbols =====")
	symbols, _ := sdk.GetSymbols()
	fmt.Println(symbols)

	fmt.Println("===== GetTicker =====")
	ticker, _ := sdk.GetTicker("")
	fmt.Println(ticker)
}
