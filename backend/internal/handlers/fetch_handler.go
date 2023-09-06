package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func HandleFetchPriceCommand(w http.ResponseWriter, r *http.Request) {

	parts := strings.Split(r.URL.Path, "/")
	if len(parts) != 3 || parts[1] != "fetch-price" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	pair := parts[2]
	price, err := fetchPrice(pair)
	if err != nil {
		http.Error(w, "Failed to fetch price", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Current price of %s is %s\n", pair, price)
}

func fetchPrice(pair string) (string, error) {
	apiURL := fmt.Sprintf("https://api.binance.com/api/v3/avgPrice?symbol=%s", pair)
	resp, err := http.Get(apiURL)
	if err != nil {
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

	// Parse the JSON response into a PriceResponse struct
	var priceResponse PriceResponse
	if err := json.Unmarshal(body, &priceResponse); err != nil {
		return "", err
	}

	// Extract the "price" field from the response
	price := priceResponse.Price

	return price, nil
}

type PriceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}
