package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config1 struct {
	Db *gorm.DB
}

func (c *Config1) initDB() {
	dbHost := ""
	dbPort := ""
	dbUser := ""
	dbPassword := ""
	dbName := ""
	env := "dev"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	enigmaDb, err := db.DB()
	err = enigmaDb.Ping()
	if err != nil {
		panic(err)
	} else {
		log.Println("Connected...")
	}

	if env == "dev" {
		c.Db = db.Debug()
	} else if env == "migration" {
		c.Db = db.Debug()
		//err := c.Db.AutoMigrate(&model.Menu{}, &model.Table{}, &model.TransType{}, &model.Customer{}, &model.Discount{}, &model.MenuPrice{}, &model.Bill{}, &model.BillDetail{})

		if err != nil {
			return
		}
	} else {
		c.Db = db
	}
}

func (c *Config1) DbConn() *gorm.DB {
	return c.Db
}

func (c *Config1) DbClose() {
	a, err := c.Db.DB()
	if err != nil {
		panic(err)
	}
	a.Close()
}

func NewConfig1() Config1 {
	cfg := Config1{}
	cfg.initDB()
	return cfg
}
