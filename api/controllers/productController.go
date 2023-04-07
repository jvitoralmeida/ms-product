package controllers

import (
	"github.com/gin-gonic/gin"
	payload "ms-product/api/model"
	"ms-product/model"
	"ms-product/service"
	"strconv"
)

func CreateNewProduct(c *gin.Context) {
	var product payload.ProductRequest
	if err := c.BindJSON(&product); err != nil {
		panic(err)
	}

	createdProduct := service.CreateNewProduct(model.Product{Name: product.Name, Price: product.Price})
	c.IndentedJSON(201, createdProduct)
}

func GetAllProducts(c *gin.Context) {
	products := service.FindAllProduct()
	c.IndentedJSON(200, products)
}

func GetProductById(c *gin.Context) {
	stringProductId := c.Param("product_id")
	productId, err := strconv.Atoi(stringProductId)
	if err != nil {
		panic(err)
	}

	product, rows := service.FindProductById(productId)

	if rows == 0 {
		c.IndentedJSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.IndentedJSON(200, product)
}

func DeleteProductById(c *gin.Context) {
	productId := c.Param("product_id")
	productIdInt, err := strconv.Atoi(productId)

	if err != nil {
		panic(err)
	}

	rows := service.DeleteProductById(productIdInt)

	if rows == 0 {
		c.IndentedJSON(404, gin.H{"message": "Product not found"})
		return
	}
	c.Status(204)
}

func UpdateProductById(c *gin.Context) {
	productId, err := strconv.Atoi(c.Param("product_id"))

	if err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	var product model.Product
	if err := c.BindJSON(&product); err != nil {
		c.IndentedJSON(400, gin.H{"message": err.Error()})
		return
	}

	product, rows := service.UpdateProduct(productId, product)

	if rows == 0 {
		c.IndentedJSON(404, gin.H{"message": "Product not found"})
		return
	}

	c.IndentedJSON(200, product)
}
