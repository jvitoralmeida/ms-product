package main

import (
	"github.com/gin-gonic/gin"
	"ms-product/api/controllers"
	"ms-product/config"
	"ms-product/database"
	"os"
)

func init() {
	config.LoadEnvs()
	database.DbConnection()
	database.RunMigrations()
}

func main() {
	r := gin.Default()

	r.GET("/product/:product_id", controllers.GetProductById)
	r.GET("/products", controllers.GetAllProducts)
	r.POST("/product", controllers.CreateNewProduct)
	r.DELETE("/product/:product_id", controllers.DeleteProductById)
	r.PUT("/product/:product_id", controllers.UpdateProductById)

	r.Run(os.Getenv("SERVER_PORT"))
}
