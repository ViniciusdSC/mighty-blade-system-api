package dbconnection

import (
	"fmt"
	"strconv"

	"github.com/ViniciusdSC/mighty-blade-system-api/internal/infrastructure/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	DBConnection interface {
		CreateConnection() (GormDB, error)
	}

	dbConnection struct {
		dbConfig config.DBConfig
	}
)

func NewDatabaseConnection(dbConfig config.DBConfig) DBConnection {
	return dbConnection{
		dbConfig,
	}
}

func (dbc dbConnection) CreateConnection() (GormDB, error) {
	sslMode := "enabled"

	if !dbc.dbConfig.GetSSLMode() {
		sslMode = "disable"
	}

	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbc.dbConfig.GetHost(),
		dbc.dbConfig.GetUser(),
		dbc.dbConfig.GetPassword(),
		dbc.dbConfig.GetDatabaseName(),
		strconv.Itoa(dbc.dbConfig.GetPort()),
		sslMode,
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
