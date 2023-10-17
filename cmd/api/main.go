package main

import (
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/handlers"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/api/services"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/config"
	dbconnection "github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/db-connection"
	"github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/validation"
	"github.com/gin-gonic/gin"
)

func main() {
	config := config.NewDBConfig()
	dbConn := dbconnection.NewDatabaseConnection(config)

	conn, err := dbConn.CreateConnection()

	if err != nil {
		panic(err)
	}

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": true,
		})
	})

	v := validation.NewValidator()
	s := services.NewAppServices(v, conn)
	h := handlers.NewAppHandler(*s)

	h.IH.Route(r)

	r.Run()
}
