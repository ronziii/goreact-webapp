package forms

type ProductForm struct {
	Name  string  `binding:"required" json:"name"`
	Price float64 `binding:"required" json:"price"`
}
