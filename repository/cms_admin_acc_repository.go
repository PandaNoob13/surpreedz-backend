package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type AdminAccRepository interface {
	FindByUsername(uname string) (model.Admin, error)
}

type adminAccRepository struct {
	db *gorm.DB
}

func (a *adminAccRepository) FindByUsername(uname string) (model.Admin, error) {
	var admin model.Admin
	result := a.db.Where("mst_admin.username = ?", uname).First(&admin)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return admin, err
		}
	}
	return admin, nil
}

func NewAdminAccRepository(db *gorm.DB) AdminAccRepository {
	repo := new(adminAccRepository)
	repo.db = db
	return repo
}
