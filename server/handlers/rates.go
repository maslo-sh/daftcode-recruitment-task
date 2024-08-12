package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

const (
	ApiUrl        = "https://openexchangerates.org"
	ApiPath       = "/api/latest.json"
	ApiKeyEnvName = "OPENEXCHANGERATE_APP_ID"
)

type FiatExchangeRatesOrigin struct {
	Base  string             `json:"base"`
	Rates map[string]float64 `json:"rates"`
}

type FiatExchangeRate struct {
	From string  `json:"from"`
	To   string  `json:"to"`
	Rate float64 `json:"rate"`
}

// Function to fetch exchange rates
func getExchangeRates(currencies []string) ([]FiatExchangeRate, error) {
	origin, err := fetchOriginalRates()

	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}

	rates := calculateRatesFromOrigin(origin, currencies)

	return rates, nil
}

func fetchOriginalRates() (*FiatExchangeRatesOrigin, error) {
	url := fmt.Sprintf("%s%s?base=USD&app_id=%s", ApiUrl, ApiPath, os.Getenv(ApiKeyEnvName))
	//url := fmt.Sprintf("%s?base=%s&app_id=%s", ApiUrl, BaseCurrency, "7b71cb28026d416682badf33cae16d88")

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
	var exchangeRates FiatExchangeRatesOrigin
	if err := json.NewDecoder(resp.Body).Decode(&exchangeRates); err != nil {
		return nil, ErrJsonDecode
	}

	return &exchangeRates, nil
}

func calculateRatesFromOrigin(origin *FiatExchangeRatesOrigin, currencies []string) []FiatExchangeRate {
	var rates []FiatExchangeRate
	baseExchangeRates := make(map[string]float64)

	for _, curr := range currencies {
		baseExchangeRates[curr] = origin.Rates[curr]

	}

	for _, fromCurr := range currencies {
		for _, toCurr := range currencies {
			if fromCurr != toCurr && baseExchangeRates[toCurr]*baseExchangeRates[fromCurr] != 0 {
				rate := FiatExchangeRate{fromCurr, toCurr, baseExchangeRates[fromCurr] / baseExchangeRates[toCurr]}
				rates = append(rates, rate)
			}
		}
	}

	return rates
}

func GetFiatExchangeRates(c *gin.Context) {
	currencies := c.Query("currencies")

	if currencies == "" {
		fmt.Printf("failed to get rates: %v\n", ErrParam)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	currenciesArr := strings.Split(currencies, ",")

	if len(currenciesArr) == 1 {
		fmt.Printf("failed to get rates: %v\n", ErrParam)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rates, err := getExchangeRates(currenciesArr)
	if err != nil {
		fmt.Printf("failed to get rates: %v\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, &rates)
}
