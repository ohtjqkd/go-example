package main

import (
	"fmt"
	UpbitTicker "go-example/src/test"
	"net/http"
	"sync"
)

const (
	upbitWebSocketURL = "wss://api.upbit.com/websocket/v1"
)

func main() {
	var wait sync.WaitGroup
	UpbitTicker.InitGetTickerMessage(&wait)
	resp, _ := http.Get("https://core.finexblock.com")
	buf := []byte{}
	resp.Body.Read(buf)
	fmt.Println(buf)
	wait.Wait()
}
