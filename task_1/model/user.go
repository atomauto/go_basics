package model

import (
	"time"

	"github.com/google/uuid"
)

// TODO: we should replace discount with bonus systems with levels (adding some gamification would be nice)
type User struct {
	id    uuid.UUID
	name  string
	phone string
	email string

	comment       string
	discount      int8
	purchaseTotal float64

	created time.Time
}

func New(name, phone, email string) *User {
	return &User{uuid.New(), name, phone, email, "", 0, 0, time.Now()}
}
