package main

import (
	UpbitTicker "go-example/src/test"
)

const (
	upbitWebSocketURL = "wss://api.upbit.com/websocket/v1"
)

func main() {
	UpbitTicker.InitGetTickerMessage()
}
