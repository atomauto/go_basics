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
