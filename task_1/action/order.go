package action

import "task_1/model"

type NewOrderAction struct {
	User *model.User

	Products *[]model.Product
	quantity []uint16
	prices   []float32

	totalWeight float32
	totalVolume float32
	totalSum    float64
}

type UpdateOrderAction struct {
	User *model.User

	Products *[]model.Product
	quantity []uint16
	prices   []float32
}

func (a *NewOrderAction) Update() string {
	o := model.NewOrder(a.User, a.Products, a.quantity, a.prices, a.totalWeight, a.totalVolume, a.totalSum)
	return "Created Order with ID " + o.Id.String()
}

func (a *UpdateOrderAction) Update(o *model.Order) string {
	o.Update(a.User, a.Products, a.quantity, a.prices)
	return "Changed Order with ID " + u.Id.String()
}
