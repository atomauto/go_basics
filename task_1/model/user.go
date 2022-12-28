package model

import (
	"time"

	"github.com/google/uuid"
)

// TODO: we should replace discount with bonus systems with levels (adding some gamification would be nice)
type User struct {
	id      uuid.UUID
	name    string
	phone   string
	email   string
	address string

	comment       string
	discount      int8
	purchaseTotal float64

	created time.Time
}

// func New(name, phone, email, address, comment string) *User {
// 	return &User{uuid.New(), name, phone, email, address, comment, 0, 0, time.Now()}
// }

func NewUser(name, phone, email, address, comment string) *User {
	return &User{uuid.New(), name, phone, email, address, comment, 0, 0, time.Now()}
}

func (u *User) Update(name, phone, email, address, comment string) {
	u.name = name
	u.phone = phone
	u.email = email
	u.address = address
	u.comment = comment
}

func (u *User) Delete() {
	u = nil
}

func (u *User) AddPurchaseTotal(purchase float64) {
	u.purchaseTotal += purchase
}
