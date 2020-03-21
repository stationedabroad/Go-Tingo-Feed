package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type EndOfDayPrices struct {
	Date   	  	time.Time `json:"date"`
	Open      	float64   `json:"open"`
	High      	float64   `json:"high"`
	Low       	float64   `json:"low"`
	Close     	float64   `json:"close"`
	Volume    	int64     `json:"volume"`
	AdjOpen   	float64   `json:"adjOpen"`
	AdjHigh   	float64   `json:"adjHigh"`
	AdjLow      float64   `json:"adjLow"`
	AdjClose  	float64   `json:"adjClose"`
	AdjVolume 	int64     `json:"adjVolume"`
	DivCash     float64   `json:"divCash"`
	SplitFactor float64   `json:"splitFactor"`	
}

func main() {

	url := "https://api.tiingo.com/tiingo/daily/aapl/prices?startDate=2020-03-02&token=%s"
	method := "GET"
	token := os.Getenv("TIINGO_TK")
	req_url := fmt.Sprintf(url, token)
	
	client := &http.Client{}
	req, err := http.NewRequest(method, req_url, nil)
	if err != nil {
		fmt.Printf("new request error: %v\n", err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("http get error: %v\n", err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body), "\n")
	var prices []EndOfDayPrices
	err = json.Unmarshal(body, &prices)
	if err != nil {
		fmt.Printf("unmarshalling error: %v\n", err)
	}
	fmt.Println(prices)
}