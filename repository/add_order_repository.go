package repository

import (
	"surpreedz-backend/model"
	"surpreedz-backend/model/dto"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddOrderRepository interface {
	AddOrder(newOrder *dto.OrderDto) (string, error)
}

type addOrderRepository struct {
	db *gorm.DB
}

func (a *addOrderRepository) AddOrder(newOrder *dto.OrderDto) (string, error) {
	tx := a.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	//create order
	toOrder := &model.Order{
		ID:              uuid.New().String(),
		BuyerId:         newOrder.BuyerId,
		ServiceDetailId: newOrder.ServiceDetailId,
		DueDate:         newOrder.DueDate,
	}
	if err := tx.Create(toOrder).Error; err != nil {
		tx.Rollback()
		return toOrder.ID, err
	}

	//find order by buyer id
	// var order model.Order
	// result := tx.Where("mst_order.buyer_id = ?", toOrder.BuyerId).Last(&order)
	// if err := result.Error; err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		tx.Rollback()
	// 		return err
	// 	} else {
	// 		tx.Rollback()
	// 		return err
	// 	}
	// }

	//create order req
	toOrderRequest := &model.OrderRequest{
		OrderId:       toOrder.ID,
		Occasion:      newOrder.Occasion,
		RecipientName: newOrder.RecipientName,
		Message:       newOrder.Message,
		Description:   newOrder.RecipientDescription,
	}

	if err := tx.Create(toOrderRequest).Error; err != nil {
		tx.Rollback()
		return toOrder.ID, err
	}

	//create order status
	toOrderStatus := &model.OrderStatus{
		OrderId: toOrder.ID,
		Status:  "Waiting for confirmation", // On Progress, Accept or Reject dari seller (case sensitive)
		Date:    time.Now(),
	}

	if err := tx.Create(toOrderStatus).Error; err != nil {
		tx.Rollback()
		return toOrder.ID, err
	}

	//create payment status
	toPaymentStatus := &model.PaymentStatus{
		OrderId:       toOrder.ID,
		StatusPayment: "unpaid",
		PaymentType:   "-",
	}

	if err := tx.Create(toPaymentStatus).Error; err != nil {
		tx.Rollback()
		return toOrder.ID, err
	}

	return toOrder.ID, tx.Commit().Error
}

func NewAddOrderRepository(db *gorm.DB) AddOrderRepository {
	repo := new(addOrderRepository)
	repo.db = db
	return repo
}
