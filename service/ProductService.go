package service

import (
	"gorm.io/gorm"
	"ms-product/database"
	"ms-product/model"
)

func CreateNewProduct(product model.Product) model.Product {
	database.DB.Create(&product)
	return product
}

func FindAllProduct() []model.Product {
	var products []model.Product
	database.DB.Find(&products)

	return products
}

func FindProductById(id int) (model.Product, int64) {

	var product model.Product
	result := database.DB.First(&product, id)

	return product, result.RowsAffected
}

func DeleteProductById(id int) int64 {
	result := database.DB.Delete(&model.Product{}, id)
	return result.RowsAffected
}

func UpdateProduct(id int, product model.Product) (model.Product, int64) {

	p, rows := FindProductById(id)

	if rows == 0 {
		return model.Product{}, 0
	}

	result := database.DB.Model(&p).Updates(model.Product{Name: product.Name, Price: product.Price})

	return p, result.RowsAffected
}

func UpdateProductWithTx(id int, newProduct model.Product) model.Product {
	var product model.Product
	err := database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&product, id).Error; err != nil {
			return err
		}

		if product.ID != 0 {

			if err := tx.Model(&product).Updates(model.Product{Name: newProduct.Name, Price: newProduct.Price}).Error; err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		panic(err)
	}
	return product
}
