package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}
}

func GetCryptoPrice(cryptoCode string) (float64, error) {
	LoadEnv()
	apiURL := os.Getenv("COINMARKETCAP_API_URL")
	apiKey := os.Getenv("COINMARKETCAP_API_KEY")
	url := fmt.Sprintf("%s?symbol=%s&convert=USD&CMC_PRO_API_KEY=%s", apiURL, cryptoCode, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			// Handle error
		}
	}(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("CoinMarketCap API request failed with status code: %d", resp.StatusCode)
	}

	var response map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, err
	}

	data, ok := response["data"].([]interface{})
	if !ok || len(data) == 0 {
		return 0, fmt.Errorf("invalid response format from CoinMarketCap API")
	}

	priceInfo, ok := data[0].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid response format from CoinMarketCap API")
	}

	quote, ok := priceInfo["quote"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid response format from CoinMarketCap API")
	}

	usd, ok := quote["USD"].(map[string]interface{})
	if !ok {
		return 0, fmt.Errorf("invalid response format from CoinMarketCap API")
	}

	price, ok := usd["price"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid response format from CoinMarketCap API")
	}

	return price, nil
}
