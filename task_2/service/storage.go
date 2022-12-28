package service

import (
	"task_2/model"
	"time"

	"github.com/google/uuid"
)

type MemoryStorage struct {
	Users map[uuid.UUID]model.User
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{make(map[uuid.UUID]model.User)}
}

func (storage *MemoryStorage) CreateUser(name, login, password, phone string, birthDate time.Time) string {
	u := model.NewUser(name, login, password, phone, birthDate)
	storage.Users[u.Id] = *u
	return u.Id.String()
}

func (storage *MemoryStorage) UpdateUser(id uuid.UUID, name, login, password, phone string, birthDate time.Time) string {
	u := storage.Users[id]
	u.Update(name, login, password, phone, birthDate)
	return u.Id.String()
}
