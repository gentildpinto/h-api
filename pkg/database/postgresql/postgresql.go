package postgresql

import (
	"os"
	"strconv"
	"time"

	"github.com/gentildpinto/h-api/internal/config"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var maxDBConnections = 500

func New(cfg config.Postgresql) (*gorm.DB, error) {
	postgresConfig := postgres.Config{
		DSN:                  "host=" + cfg.Host + " user=" + cfg.User + " password=" + cfg.Password + " dbname=" + cfg.DBName + " port=" + cfg.Port + " sslmode=" + cfg.SSLMode + " TimeZone=Africa/Luanda",
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
