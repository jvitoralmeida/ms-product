package model

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
}
