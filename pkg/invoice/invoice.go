package invoice

import (
	"github.com/JuanAispuro/composition/pkg/customer"
	"github.com/JuanAispuro/composition/pkg/invoiceitem"
)

// Invoice estructura de la factura.
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
	//Con esto es de uno a muchos
	items []invoiceitem.Item
}

// Settotal is the setter of Invoice.total
func (i *Invoice) SetTotal() {
	for _, item := range i.items {
		i.total += item.Value()
	}
}

// Constructor
func New(country, city string, client customer.Customer, items []invoiceitem.Item) Invoice {
	return Invoice{
		country: country,
		city:    city,
		client:  client,
		items:   items,
	}
}
