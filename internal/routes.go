package server

import (
	"github.com/gin-gonic/gin"
	"github.com/maslo-sh/daftcode-recruitment-task/server/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/rates", handlers.GetFiatExchangeRates)
	r.GET("/exchange", handlers.GetCryptoExchangeRates)
}
