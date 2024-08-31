package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
	"github.com/gesangwidigdo/store-management/utils"
	customresponse "github.com/gesangwidigdo/store-management/utils/customResponse"
	"github.com/gin-gonic/gin"
)

type ProductTransactionInput struct {
	Transaction_id uint `json:"transaction_id" binding:"required,numeric"`
	Product_id     uint `json:"product_id" binding:"required,numeric"`
	Quantity       int  `json:"quantity" binding:"required,numeric"`
}

type UpdateQtyInput struct {
	Quantity int `json:"quantity" binding:"required,numeric"`
}

func CreateProductTransaction(c *gin.Context) {
	var PTData ProductTransactionInput

	// bind data
	if err := utils.BindData(&PTData, c); !err {
		return
	}

	// Get Product Data
	var product models.Product
	if err := initializers.DB.First(&product, PTData.Product_id).Error; err != nil {
		utils.ReturnResponse(http.StatusBadRequest, "retrieve product data failed", "error", "data not found", c)
		return
	}

	// Check if product's stock decent or not
	if product.Stock < PTData.Quantity {
		utils.ReturnResponse(http.StatusBadRequest, "failed add data", "error", "not enough stock", c)
		return
	}

	// get transaction data
	var transaction models.Transaction
	if err := initializers.DB.First(&transaction, PTData.Transaction_id); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "retrieve transaction data failed", "error", "data not found", c)
		return
	}

	// calculate total
	totalPrice := product.Price * float64(PTData.Quantity)

	productTransaction := models.ProductTransaction{
		Transaction_id: PTData.Transaction_id,
		Product_id:     PTData.Product_id,
		Quantity:       PTData.Quantity,
		Total:          totalPrice,
	}

	if result := initializers.DB.Create(&productTransaction); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed create", "error", result.Error.Error(), c)
		return
	}

	// decrease product's stock
	if err := initializers.DB.Model(&models.Product{}).Where("id = ?", PTData.Product_id).Update("stock", product.Stock-PTData.Quantity); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", err.Error.Error(), c)
		return
	}

	// update grand total
	if err := initializers.DB.Model(&models.Transaction{}).Where("id = ?", PTData.Transaction_id).Update("grand_total", transaction.Grand_total+totalPrice); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update transaction data", "error", err.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", productTransaction, c)
}

func GetProductTransactionByTransactionID(c *gin.Context) {
	transactionID := c.Param("id")
	var pts []models.ProductTransaction

	if err := initializers.DB.Model(&models.ProductTransaction{}).Where("transaction_id = ?", transactionID).Find(&pts); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", err.Error.Error(), c)
		return
	}

	var ptResponse []customresponse.ProductTransactionResponse
	for _, ptData := range pts {
		ptResponse = append(ptResponse, customresponse.ToProductTransactionResponse(ptData))
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", ptResponse, c)
}

func UpdateQuantity(c *gin.Context) {
	var newQty UpdateQtyInput
	// bind new qty
	if err := utils.BindData(&newQty, c); !err {
		return
	}
	
	transaction_id := c.Param("transaction_id")
	product_id := c.Param("product_id")

	var ptData models.ProductTransaction

	// find data
	if result := initializers.DB.Where("transaction_id = ? AND product_id = ?", transaction_id, product_id).First(&ptData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", result.Error.Error(), c)
		return
	}
	
	// update data
	ptData.Quantity = newQty.Quantity
	if result := initializers.DB.Save(&ptData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.UPDATE_FAILED, "error", result.Error.Error(), c)
		return
	}
	
	utils.ReturnResponse(http.StatusOK, "ok", "-", "", c)
}

func DeleteProductTransaction(c *gin.Context) {
	transaction_id := c.Param("transaction_id")
	product_id := c.Param("product_id")

	var ptData models.ProductTransaction
	var transaction models.Transaction

	// Fetch transaction
	if result := initializers.DB.First(&transaction, transaction_id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", result.Error.Error(), c)
		return
	}

	if result := initializers.DB.Where("transaction_id = ? AND product_id = ?", transaction_id, product_id).First(&ptData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.GET_FAILED, "error", result.Error.Error(), c)
		return
	}

	if err := initializers.DB.Model(models.Transaction{}).Where("id = ?", transaction_id).Update("grand_total", transaction.Grand_total - ptData.Total); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update transaction data", "error", err.Error.Error(), c)
		return
	}

	if result := initializers.DB.Model(&models.ProductTransaction{}).Where("transaction_id = ? AND product_id = ?", transaction_id, product_id).Delete(&ptData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, utils.DELETE_FAILED, "error", result.Error.Error(), c)
		return
	}


	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}