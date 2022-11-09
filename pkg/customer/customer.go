package customer

//structure del costumer
type Customer struct {
	name   string
	addres string
	phone  string
}

//Returns new customer
func New(name, addres, phone string) Customer {
	return Customer{name, addres, phone}
}
