package action

import "task_1/model"

type NewUserAction struct {
	name    string
	phone   string
	email   string
	address string

	comment string
}

type UpdateUserAction struct {
	name    string
	phone   string
	email   string
	address string

	comment string
}

func (a *NewUserAction) Update() string {
	u := model.NewUser(a.name, a.phone, a.email, a.address, a.comment)
	return "Created User " + u.Name + " with ID " + u.Id.String()
}

func (a *UpdateUserAction) Update(u *model.User) string {
	u.Update(a.name, a.phone, a.email, a.address, a.comment)
	return "Changed User data, current Name " + u.Name + ", ID " + u.Id.String()
}
