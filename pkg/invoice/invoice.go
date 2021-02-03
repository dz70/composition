package invoice

// Invoice es la estructura de ciertos campos de la factura
type Invoice struct {
	country string
	city    string
	total   float64
	client  customer.Customer
}
