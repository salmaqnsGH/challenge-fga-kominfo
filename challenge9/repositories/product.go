package repositories

import (
	"latihan-jwt/models"

	"gorm.io/gorm"
)

type ProductRepository interface {
	Create(product *models.Product) error
	FindByID(productID uint) (*models.Product, error)
	Update(product *models.Product) (*models.Product, error)
	Delete(ID uint) error
	FindAll() ([]models.Product, error)
	FindAllByUserID(userID uint) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NeProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindByID(productID uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("User").First(&product, productID).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) (*models.Product, error) {
	err := r.db.Save(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) Delete(ID uint) error {
	var product models.Product
	if err := r.db.Where("id = ?", ID).First(&product).Delete(&product).Error; err != nil {
		return err
	}

	return nil
}

func (r *productRepository) FindAll() ([]models.Product, error) {
	var product []models.Product
	err := r.db.Find(&product).Error

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r *productRepository) FindAllByUserID(userID uint) ([]models.Product, error) {
	var products []models.Product
	if err := r.db.Where("user_id = ?", userID).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
