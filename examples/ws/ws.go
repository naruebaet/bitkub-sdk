package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/naruebaet/bitkub-sdk/bksdk"
	"github.com/naruebaet/bitkub-sdk/bksdk/response"
)

// main is the entry point of the program.
func main() {
	// Create a channel to receive signals for interrupt and termination.
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// Create a context with a timeout of 1 minute.
	ctx, cancel := context.WithCancel(context.Background())

	// Create a new instance of the SDK with the provided API credentials.
	sdk := bksdk.New("xxx", "xxx")

	// Get the symbols from the SDK.
	symbols, err := sdk.GetSymbols()
	if err != nil {
		panic(err)
	}

	// Create the streamLine string from the symbols.
	var streams []string
	for _, sym := range symbols.Result {
		streams = append(streams, fmt.Sprintf(bksdk.WS_TICKER_STREAM, sym.Symbol))
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
		case sig := <-sigs:
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

			// in this case we need to parse the message to ws ticker type
			var wsticker response.WsTicker
			err := json.Unmarshal([]byte(raw), &wsticker)
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Printf("Ticker : %s, Last: %f\n", wsticker.Stream, wsticker.Last)

		}
	}
}
