package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Product struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ProductTag struct {
	ID        int
	Name      string
	ProductID int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type ProductTags []*ProductTag

type ProductProductTag struct {
	ID           int
	ProductTagID int
	ProductID    int
	ProductTag   *ProductTag
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type ProductProductTags []*ProductProductTag

func main() {
	connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v ", "postgres", "postgres", "postgres", "localhost", 5432)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&ProductProductTag{})
}
