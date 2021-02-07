package form

type Product struct {
	ProductID       uint64  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductQuantity uint64  `json:"product_quantity"`
	ProductPrice    float64 `json:"product_price"`
}

type ProductRequest struct {
	ProductID       uint64  `json:"product_id"`
	ProductName     string  `json:"product_name"`
	ProductQuantity uint64  `json:"product_quantity"`
	ProductPrice    float64 `json:"product_price"`
	ProductTypeID   uint64  `json:"product_product_type_id"`
}

type UpdateProductRequest struct {
	ProductName     string  `json:"product_name"`
	ProductQuantity uint64  `json:"product_quantity"`
	ProductPrice    float64 `json:"product_price"`
	ProductTypeID   uint64  `json:"product_product_type_id"`
}
