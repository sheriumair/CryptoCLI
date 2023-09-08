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
	fmt.Println(pair)
	price, err := fetchPrice(pair)
	if err != nil {
		http.Error(w, "Failed to fetch price", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Current price of %s is %s\n", pair, price)
}

func fetchPrice(pair string) (string, error) {
	apiURL := fmt.Sprintf("https://api.binance.com/api/v3/avgPrice?symbol=%s", pair)
	//testURL := "https://www.google.com"
	fmt.Println("URL: %v %v", apiURL, pair)
	resp, err := http.Get(apiURL)
	if err != nil {
		fmt.Printf("Failed to make GET request: %s\n", err)
		return "", err
	}
	defer resp.Body.Close()
	fmt.Println("Status Code is %d", resp.StatusCode)
	fmt.Println("INSIDDE2")
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}
	fmt.Println("INSIDDE3")
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	fmt.Println("INSIDDE4")
	// Parse the JSON response into a PriceResponse struct
	var priceResponse PriceResponse
	if err := json.Unmarshal(body, &priceResponse); err != nil {
		return "", err
	}
	fmt.Println("INSIDDE5")
	// Extract the "price" field from the response
	price := priceResponse.Price

	return price, nil
}

type PriceResponse struct {
	Mins  int    `json:"mins"`
	Price string `json:"price"`
}
