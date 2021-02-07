package controller

import (
	"log"
	"net/http"
	"strconv"

	"backend-shortcourse/go-webservice/form"
	"backend-shortcourse/go-webservice/model"

	"github.com/gin-gonic/gin"
)

type ProductController struct{}

// GetAllProducts get all products in products
func (pc ProductController) ReadAll(c *gin.Context) {
	productModel := model.ProductModel{}

	products, err := productModel.ReadAll()
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting all products.",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

// Get Product by product_id
func (pc ProductController) Read(c *gin.Context) {
	productModel := model.ProductModel{}

	productId, err := strconv.ParseUint(c.Query("product_id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product_id needs to be an integer.",
		})
		return
	}

	product, err := productModel.Read(productId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting product.",
		})
		return
	}

	c.JSON(http.StatusOK, product)
}

// Create a new product
func (pc ProductController) Add(c *gin.Context) {
	productModel := model.ProductModel{}

	var request form.ProductRequest
	err := c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	err = productModel.Add(request.ProductID,
		request.ProductName, request.ProductQuantity,
		request.ProductPrice, request.ProductTypeID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while adding product.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Add product successfully.",
	})
}

// Update product by product_id
func (pc ProductController) Update(c *gin.Context) {
	productModel := model.ProductModel{}

	productId, err := strconv.ParseUint(c.Query("product_id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product_id needs to be an integer.",
		})
		return
	}

	// check if there is product belong to this product_id
	_, err = productModel.Read(productId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No product belong to this product_id.",
		})
		return
	}

	var request form.UpdateProductRequest
	err = c.BindJSON(&request)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while binding to model.",
		})
		return
	}

	err = productModel.Update(productId,
		request.ProductName, request.ProductQuantity,
		request.ProductPrice, request.ProductTypeID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while updating product.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update product Successfully.",
	})
}

// Delete product by product_id
func (pc ProductController) Delete(c *gin.Context) {
	productModel := model.ProductModel{}

	productId, err := strconv.ParseUint(c.Query("product_id"), 10, 64)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Product_id needs to be an integer.",
		})
		return
	}

	// check if there is product belong to this product_id
	_, err = productModel.Read(productId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "No product belong to this product_id.",
		})
		return
	}

	err = productModel.Delete(productId)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while deleting product.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Delete product successfully.",
	})
}
