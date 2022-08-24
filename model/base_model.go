package model

import (
	"database/sql"
	"time"
)

type Base_model struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}
