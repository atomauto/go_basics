package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Phone     string    `json:"phone"`
	BirthDate time.Time `json:"birthDate"`
}

func NewUser(name, login, password, phone string, birthDate time.Time) *User {
	return &User{uuid.New(), name, login, password, phone, birthDate}
}

func (u *User) Update(name, login, password, phone string, birthDate time.Time) {
	u.Name = name
	u.Login = login
	u.Password = password
	u.Phone = phone
	u.BirthDate = birthDate
}
