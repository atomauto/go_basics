package dummy

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
func generate(d *Dummy) *Dummy {
	for i := 0; i < d.users; i++ {

	}
	return d
}
