package main

import (
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
		})
	})

	handlers.AddInRouter(r)

	r.Run()
}
