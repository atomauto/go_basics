package action

import "task_1/model"

type NewProductAction struct {
	name            string
	baseRetailPrice float32
	minimumPrice    float32

	minimumQuantity uint16
	multiplicity    uint8

	weight float32
	volume float32
}

type UpdateProductAction struct {
	name            string
	baseRetailPrice float32
	minimumPrice    float32

	minimumQuantity uint16
	multiplicity    uint8

	weight float32
	volume float32
}

func (a *NewProductAction) Update() string {
	p := model.NewProduct(a.name, a.baseRetailPrice, a.minimumPrice, a.minimumQuantity, a.multiplicity, a.weight, a.volume)
	return "Created Product " + p.Name + " with ID " + p.Id.String()
}

func (a *UpdateProductAction) Update(p *model.Product) string {
	p.Update(a.name, a.baseRetailPrice, a.minimumPrice, a.minimumQuantity, a.multiplicity, a.weight, a.volume)
	return "Changed Product data, current Name " + p.Name + ", ID " + p.Id.String()
}
