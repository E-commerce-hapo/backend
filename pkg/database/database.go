package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/E-commerce-hapo/backend/core/config"
	log2 "github.com/E-commerce-hapo/backend/pkg/log"
)

type Database struct {
	Db *gorm.DB
}

func New(d config.Config) *Database {
	c := d.Databases.PostgresDB
	connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", c.Database, c.Username, c.Password, c.Host, c.Port, c.SSLMode)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 newLogger(),
	})
	if err != nil {
		log2.Panic(err, nil, nil)
	}
	fmt.Println("Connect database successful !")
	return &Database{Db: db}
}

func newLogger() logger.Interface {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	return logger
}
