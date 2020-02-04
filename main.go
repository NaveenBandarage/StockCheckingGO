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

//Structure for the stuct.
type StockInfo struct {
	Symbol        string  `json:"symbol"`
	LastSalePrice float32 `json:"lastSalePrice"`
	LastUpdated   int64   `json:"lastUpdated"`
}

//Structure fot the stock info.
type Stocks []StockInfo

func main() {
	j := 0
	stocks := [11]string{"AAPL", "TSLA", "GOOGL", "AMZN", "NFLX", "FB", "NIO", "BYD", "ACB", "AMD", "SPOT"}
	stockName := stocks[j]
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Print("Enter Text: ")
	// // Scans a line from Stdin(Console)
	// scanner.Scan()
	// // Holds the string that scanned
	// stockName := scanner.Text()
	fmt.Println("Checking for: ", stockName)
	//sleep time gotta go to work tomorrow will continue on this
	stockSymbol := "&symbols="

	apiKey, err := ioutil.ReadFile("api.txt")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	base := "https://cloud.iexapis.com/stable/tops?token="

	i := 0
	for {
		stockName := stocks[j]
		response, err := http.Get(base + string(apiKey) + stockSymbol + stockName)

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
		time.Sleep(time.Millisecond * 10)
		i++
		if i == 10 {
			// fmt.Println("Pausing Stock Checking atm. Do you want to check a new stock?")
			// fmt.Print("Enter Text (Y/N): ")
			// scanner.Scan()
			// // Holds the string that scanned
			// text := scanner.Text()
			// if text == "Y" {
			if j == 10 {
				j = 0
				fmt.Println("Exiting stock checker!")

				break
			} else {
				fmt.Println("Checking for: ", stocks[j+1])
				j++
			}
			// fmt.Println("Okay do you want to check a new stock?")
			// fmt.Print("Enter stock name: ")
			// scanner.Scan()
			// // Holds the string that scanned
			// text := scanner.Text()
			// stockName = text
			i = 0
			// } else {
			// 	fmt.Println("Okay do you want to keep going with", stockName, "?")
			// 	fmt.Print("Enter Text (Y/N): ")
			// 	scanner.Scan()
			// 	text := scanner.Text()
			// 	if text == "Y" {
			// 		fmt.Println("Okay we will keep moving!")
			// 		i = 0
			// 	} else {
			// 		fmt.Println("Exiting stock checker!")
			// 		break
			// 	}
		}

	}
}
