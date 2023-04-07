package tests

import (
	"ms-product/database"
	"ms-product/model"
	"ms-product/service"
	"testing"
)

func TestMain(m *testing.M) {
	database.DbConnection()
	database.RunMigrations()
	m.Run()
}

func TestCreateProduct(t *testing.T) {
	product := model.Product{Name: "Name", Price: 10}
	createdProduct := service.CreateNewProduct(product)

	if createdProduct.ID == 0 {
		t.Error("Product should be created")
	}
}

func TestUpdateProduct(t *testing.T) {
	product := model.Product{Name: "Banana", Price: 10}
	createdProduct := service.CreateNewProduct(product)

	newProduct := model.Product{Name: "Morango", Price: 12}

	updatedProdct, rows := service.UpdateProduct(createdProduct.ID, newProduct)

	if rows == 0 {
		t.Error("Should update product")
	}

	if updatedProdct.Name != newProduct.Name {
		t.Error("Should update product name")
	}

	if updatedProdct.Price != newProduct.Price {
		t.Error("Should update product price")
	}
}
