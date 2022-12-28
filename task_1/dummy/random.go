package dummy

import (
	"log"
	"math/rand"
	"strconv"
	"task_1/model"

	"github.com/ddosify/go-faker/faker"
)

const MAX_QUANTITY_IN_CART = 100

type Dummy struct {
	users                  int
	products               int
	carts                  int
	orders                 int
	generateDataWithErrors bool
}

func New(users, products, carts, orders int, generateDataWithErrors bool) *Dummy {
	return &Dummy{users, products, carts, orders, generateDataWithErrors}
}
func Generate(d *Dummy) *Dummy {
	faker := faker.NewFaker()
	var u *[]model.User
	var p *[]model.Product
	var c *[]model.Cart
	var o *[]model.Order
	for i := 0; i < d.users; i++ {
		u[i] = model.NewUser(faker.RandomUsername(), faker.RandomPhoneNumber(), faker.RandomEmail(), faker.RandomAddressCountry()+" "+faker.RandomAddressCity()+" "+faker.RandomAddressStreetAddress(), faker.RandomWords())
	}
	for i := 0; i < d.products; i++ {
		price, err := strconv.ParseFloat(faker.RandomPrice(), 32)
		if err != nil {
			log.Println("Error in converting random price to float")
		}
		minimumPrice := price / 2
		p[i] = model.NewProduct(faker.RandomProductName(), float32(price), float32(minimumPrice), uint16(rand.Intn(2)+1), uint8(rand.Intn(4)+1), rand.Float32(), rand.Float32())
	}
	for i := 0; i < d.carts; i++ {
		c[i] = model.NewCart(u[rand.Intn(users+1)])
		for j := 0; j < rand.Intn(d.products/100+1); j++ {
			c.AddProductToCart(p[rand.Intn(d.products)+1], c.User.discount*p.baseRetailPrice, rand.Intn(MAX_QUANTITY_IN_CART)+1)
		}
	}
	for i := 0; i < d.orders; i++ {
		cartIndex := rand.Intn(d.carts) + 1
		o[i] = model.NewOrder(c[cartIndex].User, c[cartIndex].Products, c[cartIndex].quantity, c[cartIndex].prices, c[cartIndex].totalWeight, c[cartIndex].totalVolume, c[cartIndex.totalSum])
	}
	return d
}
