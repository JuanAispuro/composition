package main

import (
	"fmt"

	"github.com/JuanAispuro/composition/pkg/customer"
	"github.com/JuanAispuro/composition/pkg/invoice"
	"github.com/JuanAispuro/composition/pkg/invoiceitem"
)

func main() {
	i := invoice.New(
		"Mexico",
		"Culiacan",
		customer.New("Juan", "Calle Quinta Alta", "12345"),
		[]invoiceitem.Item{
			invoiceitem.New(1, "Curso Go", 12.34),
			invoiceitem.New(2, "Curso Dart", 54.23),
			invoiceitem.New(3, "Curso Flutter", 90.00),
		},
	)
	i.SetTotal()
	fmt.Printf("%+v", i)
}
