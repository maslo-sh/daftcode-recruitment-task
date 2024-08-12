package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UsdExchangeRate struct {
	DecimalPlaces int
	Rate          float64
}

var cryptoExchangeRates = map[string]UsdExchangeRate{
	"BEER":  {18, 0.00002461},
	"FLOKI": {18, 0.0001428},
	"GATE":  {18, 6.87},
	"USDT":  {6, 0.999},
	"WBTC":  {8, 57037.22},
}

type CryptoExchangeRate struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

func GetCryptoExchangeRates(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	amount := c.Query("amount")

	if from == "" || to == "" || amount == "" {
		log.Printf("failed to calculate exchange: %v\n", ErrParam)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !cryptoExists(from) || !cryptoExists(to) {
		log.Printf("failed to calculate exchange: %v\n", ErrNoSuchCrypto)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	toEntry := cryptoExchangeRates[to]
	amountValue, err := strconv.ParseFloat(amount, toEntry.DecimalPlaces)

	if err != nil {
		log.Printf("failed to parse amount parameter: %v\n", ErrWrongFloatFormat)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rate := calculateCryptoExchangeRate(from, to, amountValue)

	c.JSON(http.StatusOK, &rate)
}

func calculateCryptoExchangeRate(from, to string, amount float64) CryptoExchangeRate {
	return CryptoExchangeRate{
		from,
		to,
		cryptoExchangeRates[from].Rate / cryptoExchangeRates[to].Rate * amount}
}

func cryptoExists(crypto string) bool {
	_, ok := cryptoExchangeRates[crypto]
	return ok
}
