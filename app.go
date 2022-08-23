package main

import "surpreedz-backend/delivery"

func main() {
	delivery.Server().Run()

	// cfg := config.NewConfig1()
	// db := cfg.DbConn()
	// defer cfg.DbClose()

	//insert to refund
	// refundRepo := repository.NewRefundRepository(db)
	// refund := model.Refund{
	// 	OrderStatusId: 1,
	// 	Reason: "asd",
	// 	Date: time.Now(),
	// }
	// err := refundRepo.Create(&refund)
	// fmt.Println(err)

	//find refund
	// refundRepo := repository.NewRefundRepository(db)
	// refund := []model.Refund{}
	// refund, _ = refundRepo.FindAll()
	// fmt.Println(refund)

	//find refund by id
	// refundRepo := repository.NewRefundRepository(db)
	// refund := model.Refund{}
	// refund, _ = refundRepo.FindById(1)
	// fmt.Println(refund)

	//update refund by id
	// refundRepo := repository.NewRefundRepository(db)
	// refundExist := model.Refund{
	// 	ID: 1,
	// }
	// _ = refundRepo.UpdateByID(&refundExist, map[string]interface{}{
	// 	"date": time.Now(),
	// })

	//delete refund by id
	// refundRepo := repository.NewRefundRepository(db)
	// refundExist := model.Refund{
	// 	ID: 1,
	// }
	// _ = refundRepo.Delete(&refundExist)

	//find orderstatus by id
	// orderStatusRepo := repository.NewOrderStatusRepository(db)
	// orderStatus := model.OrderStatus{}
	// orderStatus, _ = orderStatusRepo.FindById(1)
	// fmt.Println(orderStatus)

	//find orderstatus
	// orderStatusRepo := repository.NewOrderStatusRepository(db)
	// orderStatus := []model.OrderStatus{}
	// orderStatus, _ = orderStatusRepo.FindAll()
	// fmt.Println(orderStatus)

	//delete orderstatus by id
	// orderStatusRepo := repository.NewOrderStatusRepository(db)
	// orderStatusExist := model.OrderStatus{
	// 	ID: 1,
	// }
	// _ = orderStatusRepo.Delete(&orderStatusExist)

	//find order
	// orderRepo := repository.NewOrderRepository(db)
	// order := []model.Order{}
	// order, _ = orderRepo.FindAll()
	// fmt.Println(order)

	//find order by id
	// orderRepo := repository.NewOrderRepository(db)
	// order := model.Order{}
	// order, _ = orderRepo.FindById(1)
	// fmt.Println(order)

}
