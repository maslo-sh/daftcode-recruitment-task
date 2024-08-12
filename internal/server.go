package internal

import (
	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	RegisterRoutes(r)
	r.Run()
}
