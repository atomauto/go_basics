package model

import (
	"time"

	"github.com/google/uuid"
)

// TODO: add fake discounts system to products
type Product struct {
	id uuid.UUID

	name            string
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

func (p *Product) UpdateProduct(name string, baseRetailPrice, minimumPrice float32, minimumQuantity uint16, multiplicity uint8, weight, volume float32) {
	p.name = name
	p.baseRetailPrice = baseRetailPrice
	p.minimumPrice = minimumPrice
	p.minimumQuantity = minimumQuantity
	p.multiplicity = multiplicity
	p.weight = weight
	p.volume = volume
}

func (p *Product) RemoveProduct() {
	//How to free memory
	//and check all carts for having this product
}
