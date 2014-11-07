package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/revel/revel"
	"os"
)

var (
	DB *gorm.DB
)

func init() {
	d := os.Getenv("RBSCDB")
	if len(d) < 1 {
		revel.WARN.Printf("Please provide database connection string in env var RBSCDB\n")
		return
	}
	db, err := gorm.Open("postgres", d)
	if err != nil {
		revel.WARN.Printf("Unable to connect to %s: %v\n", d, err)
		return
	}
	DB = &db
	DB.SingularTable(true)
}
