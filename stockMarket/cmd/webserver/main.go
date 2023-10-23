package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"stockMarket/internal/web"
)

func main() {
	r := gin.Default()

	web.RegisterRoutes(r)

	port := 8080
	r.Run(fmt.Sprintf(":%d", port)) // listen and serve on 0.0.0.0:8080
}
