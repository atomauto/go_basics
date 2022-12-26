package dummy

import (
	"log"
	"math/rand"
	"strconv"
	"task_1/model"

	"github.com/ddosify/go-faker/faker"
)

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

	for i := 0; i < d.users; i++ {
		model.NewUser(faker.RandomUsername(), faker.RandomPhoneNumber(), faker.RandomEmail(), faker.RandomAddressCountry()+" "+faker.RandomAddressCity()+" "+faker.RandomAddressStreetAddress(), faker.RandomWords())
	}
	for i := 0; i < d.products; i++ {
		price, err := strconv.ParseFloat(faker.RandomPrice(), 32)
		if err != nil {
			log.Println("Error in converting random price to float")
		}
		minimumPrice := price / 2
		model.NewProduct(faker.RandomProductName(), float32(price), float32(minimumPrice), uint16(rand.Intn(2)+1), uint8(rand.Intn(4)+1), rand.Float32(), rand.Float32())
	}
	// for i := 0; i < d.carts; i++ {
	// 	model.NewCart()
	// }
	return d
}
