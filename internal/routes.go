package internal

import (
	"github.com/gin-gonic/gin"
	handlers "github.com/maslo-sh/daftcode-recruitment-task/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/rates", handlers.GetFiatExchangeRates)
	r.GET("/exchange", handlers.GetCryptoExchangeRates)
}
