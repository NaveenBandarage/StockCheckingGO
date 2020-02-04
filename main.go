package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Structure for the stuct.
type StockInfo struct {
	Symbol        string  `json:"symbol"`
	LastSalePrice float32 `json:"lastSalePrice"`
	LastUpdated   int64   `json:"lastUpdated"`
}

//Structure fot the stock info.
type Stocks []StockInfo

func main() {

	stockName := "&symbols=aapl"

	apiKey, err := ioutil.ReadFile("api.txt")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	base := "https://sandbox.iexapis.com/stable/tops?token="
	for i := 1; i <= 10; i++ {
		response, err := http.Get(base + string(apiKey) + stockName)

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(string(responseData))
		var stockVar Stocks

		jsonFile := json.Unmarshal([]byte(responseData), &stockVar)

		if jsonFile != nil {
			fmt.Println("error:", jsonFile)
		}

		fmt.Println(stockVar[0].Symbol, stockVar[0].LastSalePrice, stockVar[0].LastUpdated)
	}
}
