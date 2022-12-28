package service

import "context"

type UpdateService struct {
	UpdatesChannel chan string
}
type Action interface {
	Update() string
}

func NewUpdateService(ctx context.Context) *UpdateService {
	return &UpdateService{make(chan string)}
}
func (s *UpdateService) Update(a Action) {
	s.UpdatesChannel <- a.Update()
}
