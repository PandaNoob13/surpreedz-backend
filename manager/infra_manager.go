package manager

import (
	"log"
	"os"
	"surpreedz-backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
}

type infra struct {
	db *gorm.DB
	//cfg config.Config
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	// infra := infra{
	// 	cfg: config,
	// 	db:  resource,
	// }
	return &infra{db: resource}
}

func initDbResource(dataSourceName string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println("connected")
	}

	env := os.Getenv("ENV")

	if env == "dev" {
		db = db.Debug()
	} else if env == "migration" {
		db = db.Debug()
		db.AutoMigrate(
		//1 &model.Account{},
		//1 &model.AccountDetail{},
		//4 &model.Feedback{},
		//3 &model.Order{},
		//4 &model.OrderRequest{},
		//4 &model.OrderStatus{},
		//1 &model.PhotoProfile{},
		//5 &model.Refund{},
		//2 &model.ServiceDetail{},
		//3 &model.ServicePrice{},
		//5 &model.VideoProfile{},
		//4 &model.VideoResult{},
		)
		if err != nil {
			return nil, err
		}
	}
	return db, nil
}
