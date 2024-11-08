package customresponse

import (
	"time"

	"github.com/gesangwidigdo/store-management/models"
)

type TransactionResponse struct {
	ID               uint                    `json:"id"`
	Grand_total      float64                 `json:"grand_total"`
	Transaction_time time.Time               `json:"transaction_time"`
	Status					 bool                    `json:"status"`
	Employee         ForeignEmployeeResponse `json:"employee_data"`
}

func ToTransactionResponse(transaction models.Transaction) TransactionResponse {
	return TransactionResponse{
		ID:               transaction.ID,
		Grand_total:      transaction.Grand_total,
		Transaction_time: transaction.Transaction_time,
		Status:           transaction.Status,
		Employee:         ForeignEmployeeResponseData(transaction.Employee),
	}
}
