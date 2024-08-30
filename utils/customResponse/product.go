package customresponse

import "github.com/gesangwidigdo/store-management/models"

type ProductResponse struct {
	ID           uint    `json:"id"`
	Product_name string  `json:"product_name"`
	Price        float64 `json:"price"`
	Stock        int     `json:"stock"`
}

func ToProductResponse(product models.Product) ProductResponse {
	return ProductResponse{
		ID:           product.ID,
		Product_name: product.Product_name,
		Price:        product.Price,
		Stock:        product.Stock,
	}
}
