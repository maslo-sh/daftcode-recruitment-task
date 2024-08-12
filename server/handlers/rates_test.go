package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetFiatExchangeRatesHandler(t *testing.T) {
	defer gock.Off()

	r := gin.Default()
	r.GET("/rates", GetFiatExchangeRates)

	w := httptest.NewRecorder()

	gock.New(ApiUrl).
		Get(ApiPath).
		Reply(200).
		JSON(FiatExchangeRatesOrigin{Base: "USD", Rates: map[string]float64{
			"PLN": 4.0,
			"GBP": 1.5,
			"CHF": 1.0,
		}})

	expected := "[{\"from\":\"PLN\",\"to\":\"GBP\",\"rate\":2.6666666666666665},{\"from\":\"PLN\",\"to\":\"CHF\",\"rate\":4},{\"from\":\"GBP\",\"to\":\"PLN\",\"rate\":0.375},{\"from\":\"GBP\",\"to\":\"CHF\",\"rate\":1.5},{\"from\":\"CHF\",\"to\":\"PLN\",\"rate\":0.25},{\"from\":\"CHF\",\"to\":\"GBP\",\"rate\":0.6666666666666666}]"

	req, _ := http.NewRequest("GET", "/rates?currencies=PLN,GBP,CHF", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 200)

	actual := w.Body.String()

	assert.Equal(t, expected, actual)
}

func TestGetFiatExchangeRatesHandlerWrongParametersFailure(t *testing.T) {
	r := gin.Default()
	r.GET("/rates", GetFiatExchangeRates)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/rates?currencies=PLN", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
}

func TestGetFiatExchangeRatesHandlerEmptyParametersFailure(t *testing.T) {
	r := gin.Default()
	r.GET("/rates", GetFiatExchangeRates)

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/rates?currencies=", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, w.Code, 400)
}
