package database

import (
	"fmt"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"

	"github.com/kiem-toan/cmd/audit-server/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db  *gorm.DB
	abs string
}

func New(d config.Config) *Database {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:              time.Second,   // Slow SQL threshold
			LogLevel:                   logger.Info, // Log level
			IgnoreRecordNotFoundError: true,           // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)


	c := d.Databases.Postgres
	connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", c.Database, c.Username, c.Password, c.Host, c.Port, c.SSLMode)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return &Database{Db: db}
}
