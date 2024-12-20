package controllers

import (
	"net/http"

	"github.com/gesangwidigdo/store-management/initializers"
	"github.com/gesangwidigdo/store-management/models"
	"github.com/gesangwidigdo/store-management/utils"
	customresponse "github.com/gesangwidigdo/store-management/utils/customResponse"
	"github.com/gin-gonic/gin"
)

type productInput struct {
	Product_name string  `json:"product_name" binding:"required"`
	Price        float64 `json:"price" binding:"required,numeric"`
	Stock        int     `json:"stock"`
}

type updateStockInput struct {
	ID    uint `json:"product_id" binding:"required,numeric"`
	Stock int  `json:"stock" binding:"required,numeric"`
}

func CreateProduct(c *gin.Context) {
	var productData productInput

	// bind data
	if err := utils.BindData(&productData, c); !err {
		return
	}

	product := models.Product{
		Product_name: productData.Product_name,
		Price:        productData.Price,
		Stock:        productData.Stock,
	}

	if result := initializers.DB.Create(&product); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed create", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusCreated, "ok", "data", product, c)
}

func GetAllProduct(c *gin.Context) {
	var products []models.Product

	if result := initializers.DB.Find(&products); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed get data", "error", result.Error.Error(), c)
		return
	}

	var productResponses []customresponse.ProductResponse
	for _, product := range products {
		productResponses = append(productResponses, customresponse.ToProductResponse(product))
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", productResponses, c)
}

func GetProductByID(c *gin.Context) {
	id := c.Param("id")
	var product models.Product

	if result := initializers.DB.First(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed get data", "error", result.Error.Error(), c)
		return
	}

	productResponse := customresponse.ToProductResponse(product)

	utils.ReturnResponse(http.StatusOK, "ok", "data", productResponse, c)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var productData productInput

	if err := utils.BindData(&productData, c); !err {
		return
	}

	var product models.Product
	if result := initializers.DB.First(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", result.Error.Error(), c)
		return
	}

	// new data
	newData := models.Product{
		Product_name: productData.Product_name,
		Price:        productData.Price,
		Stock:        productData.Stock,
	}

	if result := initializers.DB.Model(&product).Updates(newData); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", result.Error.Error(), c)
		return
	}

	var updatedProduct models.Product
	if result := initializers.DB.First(&updatedProduct, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "new_data", updatedProduct, c)
}

func UpdateProductStock(c *gin.Context) {
	var productData updateStockInput

	if err := utils.BindData(&productData, c); !err {
		return
	}

	id := productData.ID

	var product models.Product
	if result := initializers.DB.First(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to get data", "error", result.Error.Error(), c)
		return
	}

	// update stock
	updateStock := models.Product{
		Product_name: product.Product_name,
		Price:        product.Price,
		Stock:        product.Stock + productData.Stock,
	}

	if err := initializers.DB.Model(&product).Updates(updateStock); err.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed to update stock", "error", err.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "data", updateStock, c)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Product

	if result := initializers.DB.Delete(&product, id); result.Error != nil {
		utils.ReturnResponse(http.StatusBadRequest, "failed delete data", "error", result.Error.Error(), c)
		return
	}

	utils.ReturnResponse(http.StatusOK, "ok", "", nil, c)
}
