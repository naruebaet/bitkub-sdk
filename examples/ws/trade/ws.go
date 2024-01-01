package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/naruebaet/bitkub-sdk/bksdk"
	"github.com/naruebaet/bitkub-sdk/bksdk/response"
)

func main() {
	// Create a channel to receive signals for interrupt and termination.
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Create a context with a timeout of 1 minute.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Create a new instance of the SDK with the provided API credentials.
	sdk := bksdk.New("xxx", "xxx")

	// Get the symbols from the SDK.
	symbols, err := sdk.GetSymbols()
	if err != nil {
		log.Fatal(err)
	}

	// Create the streamLine string from the symbols.
	var streams []string
	for _, sym := range symbols.Result {
		streams = append(streams, fmt.Sprintf(bksdk.WS_TRADE_STREAM, sym.Symbol))
	}
	streamLine := strings.Join(streams, ",")

	fmt.Println("ws starting...")

	// Create a channel for reading the WebSocket messages.
	reader := make(chan string)
	defer close(reader)

	// Create a WebSocket connection with the streamLine and reader.
	bksdk.CreateWsConnection(streamLine, reader, ctx)

	// Read messages from the reader channel until the context is done.
	for {
		select {
		case sig := <-signals:
			fmt.Println("-------------------")
			fmt.Println("Signal operated")
			fmt.Println(sig)
			cancel()
			fmt.Println("-------------------")
		default:
			raw := <-reader
			if raw == "Connection closed" {
				fmt.Println(raw)
				return
			}
			// Parse the message to ws ticker type.
			raws := strings.Split(raw, "\n")
			for _, rawv := range raws {
				var wsTrade response.WsTrade
				err = json.Unmarshal([]byte(rawv), &wsTrade)
				if err != nil {
					fmt.Println("---------ERR----------")
					fmt.Println(err)
					fmt.Println(raw)
					fmt.Println("---------ERR----------")
					continue
				}
				fmt.Printf("Trade: %s, Txn: %s, Rate: %f\n", wsTrade.Stream[17:], wsTrade.Txn, wsTrade.Rat)
			}
		}
	}
}
