package model

type CRUD interface {
	//New isn't working (((
	New() interface{}
	Update(interface{})
	Delete(interface{})
}
