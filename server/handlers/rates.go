package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	ApiUrl = "https://openexchangerates.org/api/latest.json"
)

type ExchangeRatesOrigin struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type ExchangeRate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

// Function to fetch exchange rates
func getExchangeRates(currencies []string) ([]ExchangeRate, error) {
	origin, err := fetchOriginalRates()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	rates := calculateRatesFromOrigin(origin, currencies)

	return rates, nil
}

func fetchOriginalRates() (*ExchangeRatesOrigin, error) {
	//url := fmt.Sprintf("%s?app_id=%s", ApiUrl, os.Getenv("API_KEY"))
	url := fmt.Sprintf("%s?app_id=%s", ApiUrl, "7b71cb28026d416682badf33cae16d88")

	// Make the HTTP request
	resp, err := http.Get(url)
	if err != nil {
		return nil, ErrApi
	}
	defer resp.Body.Close()

	// Check if the response status is OK
	if resp.StatusCode != http.StatusOK {
		return nil, ErrStatusCode
	}

	// Deserialize the JSON response into the ExchangeRates struct
	var exchangeRates ExchangeRatesOrigin
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRates); err != nil {
		return nil, ErrJsonDecode
	}

	return &exchangeRates, nil
}

func calculateRatesFromOrigin(origin *ExchangeRatesOrigin, currencies []string) []ExchangeRate {
	var rates []ExchangeRate
	for _, curr := range currencies {
		rate := ExchangeRate{origin.Base, curr, origin.Rates[curr]}
		rates = append(rates, rate)
	}

	return rates
}

func GetRates(c *gin.Context) {
	currencies := c.Query("currencies")

	if currencies == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	currenciesArr := strings.Split(currencies, ",")

	if len(currenciesArr) == 1 {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rates, err := getExchangeRates(currenciesArr)
	if err != nil {
		fmt.Printf("failed to get rates: %v", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, &rates)
}
