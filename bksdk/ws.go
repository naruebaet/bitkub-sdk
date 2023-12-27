package bksdk

import (
	"context"
	"fmt"

	"github.com/gorilla/websocket"
)

const WS_HOST = "wss://api.bitkub.com/websocket-api/"
const WS_TICKER_STREAM = "market.ticker.%s"
const WS_TRADE_STREAM = "market.trade.%s"

// CreateWsConnection creates a websocket connection.
//
// Parameters:
// - streamName: The name of the stream.
// - reader: A channel used to read messages from the websocket.
// - ctx: The context object for managing the connection's lifecycle.
func CreateWsConnection(streamName string, reader chan string, ctx context.Context) {

	// Create stream name
	streamName = WS_HOST + streamName

	// Create a new websocket connection
	conn, _, err := websocket.DefaultDialer.Dial(streamName, nil)
	if err != nil {
		reader <- fmt.Sprintf("Failed to connect to websocket: %s", err.Error())
	}

	// Start a goroutine to read messages from the websocket
	go func() {
		defer conn.Close()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				reader <- fmt.Sprintf("Error reading message from websocket: %s", err.Error())
			}

			// Send the message to the channel and continue reading
			select {
			case <-ctx.Done():
				reader <- "Connection closed"
				return
			default:
				reader <- string(message)
			}
		}
	}()
}
