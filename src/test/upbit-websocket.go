package upbit_websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
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

type Ticker struct {
	Market string `json:"market"`
	KoreanName string `json:"korean_name"`
	EnglishName string `json:"english_name"`
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

func InitGetTickerMessage(wait *sync.WaitGroup) {
	var client upbit.Client
	client.Client = http.DefaultClient
	client.AccessKey = ""
	client.SecretKey = ""
	resp, _ := client.Get("https://api.upbit.com/v1/market/all")
	fmt.Printf("resp: %v\n", resp)
	p := []byte{}
	n, _ := resp.Body.Read(p)
	fmt.Printf("n: %v\n", n)
	var tickers = []Ticker{}
	json.Unmarshal(p, tickers)
	wait.Add(len(tickers))
	for _, tick := range tickers {
		fmt.Println("before go routine")
		go func(ticker string) {
			defer wait.Done()
			MakeWebsocket(ticker)
		} (tick.Market)
	}
}
