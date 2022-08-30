package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type PasswordRepository interface {
	FindByAccountId(id int) (model.Password, error)
}

type passwordRepository struct {
	db *gorm.DB
}

func (p *passwordRepository) FindByAccountId(id int) (model.Password, error) {
	var password model.Password
	result := p.db.Where("mst_password.account_id = ?", id).First(&password)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return password, err
		} else {
			return password, err
		}
	}
	return password, nil
}

func NewPasswordRepository(db *gorm.DB) PasswordRepository {
	repo := new(passwordRepository)
	repo.db = db
	return repo
}
