package model

import (
	"time"

	"github.com/google/uuid"
)

// TODO: we should replace discount with bonus systems with levels (adding some gamification would be nice)
type User struct {
	Id      uuid.UUID
	Name    string
	Phone   string
	Email   string
	Address string

	Comment       string
	Discount      int8
	PurchaseTotal float64

	Created time.Time
}

// func New(name, phone, email, address, comment string) *User {
// 	return &User{uuid.New(), name, phone, email, address, comment, 0, 0, time.Now()}
// }

func NewUser(name, phone, email, address, comment string) *User {
	return &User{uuid.New(), name, phone, email, address, comment, 0, 0, time.Now()}
}

func (u *User) Update(name, phone, email, address, comment string) {
	u.Name = name
	u.Phone = phone
	u.Email = email
	u.Address = address
	u.Comment = comment
}

func (u *User) Delete() {
	u = nil
}

func (u *User) AddPurchaseTotal(purchase float64) {
	u.PurchaseTotal += purchase
}
