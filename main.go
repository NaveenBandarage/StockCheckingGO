package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

//StockInfo struct.
type StockInfo struct {
	Symbol        string  `json:"symbol"`
	LastSalePrice float32 `json:"lastSalePrice"`
	LastUpdated   int64   `json:"lastUpdated"`
}

//Stocks struck for the stock info.
type Stocks []StockInfo

func main() {
	//for incrementing the stocks array
	j := 0

	//array containing all of the stocks that i want to check.
	stocks := [11]string{"AAPL", "TSLA", "GOOGL", "AMZN", "NFLX", "FB", "NIO", "BYD", "ACB", "AMD", "SPOT"}
	stockName := stocks[j]

	fmt.Println("Checking for:", stockName)

	// Constructing the request url
	stockSymbol := "&symbols="
	base := "https://cloud.iexapis.com/stable/tops?token="

	apiKey, err := ioutil.ReadFile("api.txt")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	//for iterating through a stock.
	i := 0

	//infinte loop that goes through until a condition met that breaks the loop
	for {
		//
		// stockName := stocks[j] //getting the stockName from the stocks
		response, err := http.Get(base + string(apiKey) + stockSymbol + stockName)
		//

		//response, err := http.Get("https://cloud.iexapis.com/stable/tops?token=sk_1679b02c02f245278e2ddf3daaa49b4f")
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		defer response.Body.Close()

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
		//fmt.Println(string(responseData))
		fmt.Println(stockVar[0].Symbol, stockVar[0].LastSalePrice, stockVar[0].LastUpdated)
		time.Sleep(time.Millisecond * 1)
		i++
		if i == 10 {
			if j == 10 {
				j = 0
				fmt.Println("Exiting stock checker!")
				break
			} else {
				fmt.Println("Checking for: ", stocks[j+1])
				j++
			}
			i = 0
		}

	}
}
