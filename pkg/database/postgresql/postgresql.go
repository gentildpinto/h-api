package postgresql

import (
	"os"
	"strconv"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var maxDBConnections = 500

func New(user, password, host, port, dbName, sslMode string) (*gorm.DB, error) {
	postgresConfig := postgres.Config{
		DSN:                  "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbName + " port=" + port + " sslmode=" + sslMode + " TimeZone=Africa/Luanda",
		PreferSimpleProtocol: true,
	}

	database, err := gorm.Open(postgres.New(postgresConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   logger.Default.LogMode(logger.Info),
		CreateBatchSize:                          1000,
		AllowGlobalUpdate:                        true,
	})

	if err != nil {
		return nil, err
	}

	if connections, err := strconv.Atoi(os.Getenv("MAX_DB_CONNECTIONS")); err == nil {
		maxDBConnections = connections
	}

	maxOpenConnections := int(float64(maxDBConnections) * 0.75)
	maxIdleConnections := int(maxDBConnections) - maxOpenConnections

	sqlDB, err := database.DB()

	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(maxOpenConnections)
	sqlDB.SetMaxIdleConns(maxIdleConnections)
	sqlDB.SetConnMaxLifetime(time.Hour * 8)

	return database, nil
}
