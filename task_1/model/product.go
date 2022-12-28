package model

import (
	"time"

	"github.com/google/uuid"
)

// TODO: add fake discounts system to products
type Product struct {
	Id uuid.UUID

	Name            string
	baseRetailPrice float32
	minimumPrice    float32

	minimumQuantity uint16
	multiplicity    uint8

	weight float32
	volume float32

	created time.Time
}

func NewProduct(name string, baseRetailPrice, minimumPrice float32, minimumQuantity uint16, multiplicity uint8, weight, volume float32) *Product {
	return &Product{uuid.New(), name, baseRetailPrice, minimumPrice, minimumQuantity, multiplicity, weight, volume, time.Now()}
}

func New(name string, baseRetailPrice, minimumPrice float32, minimumQuantity uint16, multiplicity uint8, weight, volume float32) *Product {
	return &Product{uuid.New(), name, baseRetailPrice, minimumPrice, minimumQuantity, multiplicity, weight, volume, time.Now()}
}

func (p *Product) Update(name string, baseRetailPrice, minimumPrice float32, minimumQuantity uint16, multiplicity uint8, weight, volume float32) {
	p.Name = name
	p.baseRetailPrice = baseRetailPrice
	p.minimumPrice = minimumPrice
	p.minimumQuantity = minimumQuantity
	p.multiplicity = multiplicity
	p.weight = weight
	p.volume = volume
}

func (p *Product) Delete() {
	p = nil
	//We must check all carts for having this product
}
