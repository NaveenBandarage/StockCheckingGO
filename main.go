package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Struct for testing response
type Response struct {
	Symbol        string  `json:"symbol"`
	LastSalePrice float32 `json:"lastSalePrice"`
} //my brain is so gone bruh

type Post struct {
	Symbol        string  `json:"symbol"`
	LastSalePrice float32 `json:"lastSalePrice"`
}

type Posts []Post

func main() {
	stockName := "&symbols=aapl"

	apiKey, err := ioutil.ReadFile("api.txt")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	base := "https://sandbox.iexapis.com/stable/tops?token="
	response, err := http.Get(base + string(apiKey) + stockName)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))
	var posts Posts

	jsonFile := json.Unmarshal([]byte(responseData), &posts)

	if jsonFile != nil {
		fmt.Println("error:", jsonFile)
	}
	fmt.Println(posts[0].Symbol)

}
