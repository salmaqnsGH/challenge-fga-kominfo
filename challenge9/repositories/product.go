package repositories

import (
	"latihan-jwt/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(productID uint) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NeProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) FindByID(productID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Update(product *models.Product) (*models.Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}