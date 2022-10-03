package repository

import (
	"errors"
	"surpreedz-backend/model"

	"gorm.io/gorm"
)

type CMSOrderRepository interface {
	GetAllOrder(table string, condition string) ([]model.Order, error)
}

type cmsOrderRepository struct {
	db *gorm.DB
}

func (c *cmsOrderRepository) GetAllOrder(table string, condition string) ([]model.Order, error) {
	var orders []model.Order
	result := c.db.Joins("inner join mst_order_status on mst_order.id = mst_order_status.order_id")
	result = result.Select("DISTINCT ON (mst_order_status.order_id) mst_order_status.order_id", "mst_order_status.status", "mst_order_status.date").Order("mst_order_status.date desc").Preload("OrderStatus").Where(table, condition).Find(&orders)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}
	return orders, nil
}

func NewCMSOrderRepository(db *gorm.DB) CMSOrderRepository {
	repo := new(cmsOrderRepository)
	repo.db = db
	return repo
}
