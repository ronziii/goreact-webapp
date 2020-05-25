package entities

type Product struct {
	SequentialIdentifier
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Timestamps
}

type ProductList struct {
	Products []*Product `json:"products"`
}
