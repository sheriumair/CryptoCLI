package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

/*
*  Handles fetch endpoint functionality
 */

func HandleFetchPriceCommand(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	pair := parts[2]
	price, err := fetchPrice(pair)
	//Will throw an error if pair entered is incorrect
	if err != nil {
		http.Error(w, "Failed to fetch price as the pair entered is incorrect ", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Current price of %s is %s\n", pair, price)
}

// Pinging the Binance API
func fetchPrice(pair string) (string, error) {
	apiURL := fmt.Sprintf("https://api.binance.com/api/v3/avgPrice?symbol=%s", pair)
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Failed to make GET request: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var priceResponse PriceResponse
	if err := json.Unmarshal(body, &priceResponse); err != nil {
		return "", err
	}
	price := priceResponse.Price
	return price, nil
}

type PriceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}
