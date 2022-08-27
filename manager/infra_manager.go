package manager

import (
	"log"
	"surpreedz-backend/config"
	"surpreedz-backend/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db  *gorm.DB
	cfg config.Config
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	infra := infra{
		cfg: config,
		db:  resource,
	}
	return &infra
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("connected")
	}

	env := "dev"

	if env == "dev" {
		db = db.Debug()
	} else if env == "migration" {
		db = db.Debug()
		db.AutoMigrate(
			&model.Account{},
			&model.AccountDetail{},
			&model.PhotoProfile{},
			&model.ServiceDetail{},
			&model.ServicePrice{},
			&model.Order{},
			&model.OrderRequest{},
			&model.OrderStatus{},
			&model.Feedback{},
			&model.VideoResult{},
			&model.Refund{},
			&model.VideoProfile{},
			&model.Password{},
		)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
