package action

import (
	"strconv"
	"task_1/model"
	"time"
)

type NewCartAction struct {
	User *model.User

	Products *[]model.Product
	quantity []uint16
	prices   []float32
	added    []time.Time
}

type AddProductToCartAction struct {
	p        *model.Product
	price    float32
	quantity uint16
}

type ChangeQuantityAction struct {
	p        *model.Product
	quantity uint16
}

type DeleteProductFromCartAction struct {
	p *model.Product
}

func (a *NewCartAction) Update(u *model.User) string {
	c := model.NewCart(u)
	return "Created Cart with ID " + c.Id.String()
}

func (a *AddProductToCartAction) Update(c *model.Cart) string {
	c.AddProductToCart(a.p, a.price, a.quantity)
	return "Added Product with ID " + a.p.Id.String() + " to cart with ID " + c.Id.String()
}

func (a *ChangeQuantityAction) Update(c *model.Cart) string {
	c.ChangeQuantity(a.p, a.quantity)
	return "Changed Quantity to " + strconv.FormatUint(uint64(a.quantity), 10) + " of Product " + a.p.Id.String() + " in Cart " + c.Id.String()
}

func (a *DeleteProductFromCartAction) Update(c *model.Cart) string {
	c.DeleteProductFromCart(a.p)
	return "Removed Product " + a.p.Id.String() + " from Cart " + c.Id.String()
}
