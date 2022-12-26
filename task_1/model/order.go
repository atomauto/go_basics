package model

import (
	"time"

	"github.com/google/uuid"
)

// TODO: implement better status support as struct, suggest colors
// TODO: implement shipping calculation and limits of weight/volume according to delivery type
type Order struct {
	id   uuid.UUID
	User *User

	Products *[]Product
	quantity []uint16
	prices   []float32

	totalWeight float32
	totalVolume float32
	totalSum    float64

	shippingPrice float32
	deliveryType  int8 //0-Retail Store,1-Self-Delivery,2-Our Shipping,3-Marketplace Shipping,4-CDEK,5-Boxberry,
	//6-Yandex Logistic,7-Ozon Logistic,8-
	trackNumber string

	paymentMethod int8 //0-Cash,1-Card,2-Bank account,3-Qiwi,4-Other payment aggregator,-1-Deferred Payment

	status int8 //0-created,1-received,2-approved,3-batching,4-dispatched from warehouse,5-shipped from warehouse,
	//6-in delivery,7-delivered,8-closed,-1-canceled,-2-unsuccesful delivery
	created time.Time
}

// TODO: totalWeight and further params aren't handled now
func NewOrder(User *User, Products *[]Product, quantity []uint16, prices []float32) *Order {
	return &Order{uuid.New(), User, Products, quantity, prices, 0, 0, 0, 0, 0, "", 0, 0, time.Now()}
}
