package model

import (
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	id   uuid.UUID
	User *User

	Products *[]Product
	quantity []uint16
	prices   []float32
	added    []time.Time

	totalWeight float32
	totalVolume float32
	totalSum    float64

	created time.Time
}

/*Checks when you add a product:
- price with user discount >= Product.minimumPrice
- quantity >= Product.minQuantity AND quantity%Product.multiplicity==0
*/
