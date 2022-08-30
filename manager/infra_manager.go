package manager

import (
	"log"
	"surpreedz-backend/config"
	"surpreedz-backend/model"

	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Infra interface {
	SqlDb() *gorm.DB
	AzrClient() *azblob.ServiceClient
}

type infra struct {
	db         *gorm.DB
	azrService *azblob.ServiceClient
	cfg        config.Config
}

func (i *infra) SqlDb() *gorm.DB {
	return i.db
}

func (i *infra) AzrClient() *azblob.ServiceClient {
	return i.azrService
}

func NewInfra(config config.Config) Infra {
	resource, err := initDbResource(config.DataSourceName)
	if err != nil {
		log.Fatal(err.Error())
	}

	service, err := initAzureService(config.AzureConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	infra := infra{
		cfg:        config,
		azrService: service,
		db:         resource,
	}
	return &infra
}

func initAzureService(azureConfig config.AzureConfig) (*azblob.ServiceClient, error) {
	cred, err := azblob.NewSharedKeyCredential(azureConfig.AccountName, azureConfig.AccountKey)
	if err != nil {
		panic("Can't connect to surpreedz azure storage")
	}
	serviceClient, err := azblob.NewServiceClientWithSharedKey(azureConfig.ServiceUrl, cred, nil)
	return serviceClient, err
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
