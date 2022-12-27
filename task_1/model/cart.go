package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	id   uuid.UUID
	User *User

	Products *[]Product
	quantity []uint16
	prices   []float32
	added    []time.Time

	totalWeight float32
	totalVolume float32
	totalSum    float64

	created time.Time
}

func NewCart(User *User) *Cart {
	return &Cart{uuid.New(), User, nil, nil, nil, nil, 0, 0, 0, time.Now()}
}

func (c *Cart) addProductToCart(p *Product, price float32, quantity uint16) {
	if p.minimumPrice >= price {
		price = p.minimumPrice
		fmt.Printf("Sorry, the minimal price for product is %f", p.minimumPrice)
	}
	if quantity < p.minimumQuantity {
		quantity = p.minimumQuantity
		fmt.Printf("Sorry, the minimal amount to purchase is %v", quantity)
	}
	if quantity%uint16(p.multiplicity) != 0 {
		quantity = uint16(p.multiplicity)
		fmt.Printf("Sorry, quantity should be multiple to %v", p.multiplicity)
	}
	c.Products = append(c.Products, p)
	c.quantity = append(c.quantity, quantity)
	c.prices = append(c.prices, price)
	c.added = append(c.added, time.Now())

	c.totalWeight += p.weight*float32(quantity)
	c.totalVolume += p.volume*float32(quantity)
	c.totalSum += float64(price)*float64(quantity)
}

func (c *Cart) changeQuantity(p *Product, quantity uint16) {
	if quantity < p.minimumQuantity {
		quantity = p.minimumQuantity
		fmt.Printf("Sorry, the minimal amount to purchase is %v", quantity)
		return
	}
	if quantity%uint16(p.multiplicity) != 0 {
		quantity = uint16(p.multiplicity)
		fmt.Printf("Sorry, quantity should be multiple to %v", p.multiplicity)
		return
	}
	c.quantity[?] = quantity
	c.totalWeight += p.weight*float32(quantity)
	c.totalVolume += p.volume*float32(quantity)
	c.totalSum += float64(c.prices[?])*float64(quantity)
}

func remove(s []int, i int) []int {
    s[i] = s[len(s)-1]
    return s[:len(s)-1]
}


func (c *Cart) deleteProductFromCart(p *Product) {
	c.Products[?] = c.Products[len(c.Products)-1]
	c.Products = c.Products[:len(c.Products)-1]
	c.quantity[?] = c.quantity[len(c.quantity)-1]
	c.quantity = c.quantity[:len(c.quantity)-1]
	c.prices[?] = c.prices[len(c.prices)-1]
	c.prices = c.prices[:len(c.prices)-1]

	c.totalWeight -= p.weight*float32(p.quantity[?])
	c.totalVolume -= p.volume*float32(p.quantity[?])
	c.totalSum -= float64(prices[?])*float64(p.quantity[?])
	//Maybe, we should have created new structure cartItem, because in current realisation we cannot identify cart item properly 
}