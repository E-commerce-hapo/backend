package database

import (
	"fmt"
	"github.com/kiem-toan/cmd/audit-server/config"
	"github.com/kiem-toan/infrastructure/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Db  *gorm.DB
	abs string
}

func New(d config.Config) *Database {
	c := d.Databases.Postgres
	connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", c.Database, c.Username, c.Password, c.Host, c.Port, c.SSLMode)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:logger.New(),
	})
	if err != nil {
		panic(err)
	}
	return &Database{Db: db}
}
