package repositories

import (
	"latihan-jwt/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProductByID(productID uint) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NeProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) GetProductByID(productID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}
