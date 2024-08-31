package customresponse

import "github.com/gesangwidigdo/store-management/models"

type ProductTransactionResponse struct {
	Transaction_id uint    `json:"transaction_id"`
	Product_id     uint    `json:"product_id"`
	Quantity       int     `json:"quantity"`
	Total          float64 `json:"total"`
}

func ToProductTransactionResponse(ptData models.ProductTransaction) ProductTransactionResponse {
	return ProductTransactionResponse{
		Transaction_id: ptData.Transaction_id,
		Product_id:     ptData.Product_id,
		Quantity:       ptData.Quantity,
		Total:          ptData.Total,
	}
}
