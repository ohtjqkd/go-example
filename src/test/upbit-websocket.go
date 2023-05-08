package upbit_websocket

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"

	upbit "github.com/pronist/upbit/client"
)

type Data struct{
	Type string `json:"type"`
	Code string `json:"code"`
	OpeningPrice float32 `json:"opening_price"`
	ClosePrice float32 `json:"high_price"`
	AskBid string `json:"ask_bid"`
}

func MakeWebsocket(orderSymbol string) error {
	websocket, err := upbit.NewWebsocketClient("ticker", []string{orderSymbol}, true, true)	
	if err != nil {return nil}
	j, err := json.Marshal(websocket.Data)
	websocket.Ws.WriteMessage(1, j)
	GetTickerMessage(websocket)
	return nil
}

func GetTickerMessage(ws *upbit.WebsocketClient) {
	var data Data
	fmt.Println("here")
	for {
		_, p, err := ws.Ws.ReadMessage()
		if err != nil {
			fmt.Println(err)
		}
		json.Unmarshal(p, &data)
		fmt.Printf("data: %v\n", data)
		time.Sleep(1)
	}
}

func InitGetTickerMessage() {
	tickers := []string{"KRW-BTC", "KRW-ADA"}
	var wait sync.WaitGroup
	wait.Add(len(tickers))
	for _, tick := range tickers {
		fmt.Println("before go routine")
		go func(ticker string) {
			defer wait.Done()
			MakeWebsocket(ticker)
		} (tick)
	}
	wait.Wait()
}
