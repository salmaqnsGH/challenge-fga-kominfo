package repositories

import (
	"latihan-jwt/models"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	Mock mock.Mock
}

func (r *MockProductRepository) FindByID(productID uint) (*models.Product, error) {
	arguments := r.Mock.Called(productID)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	product := arguments.Get(0).(*models.Product)

	return product, arguments.Error(1)
}

func (r *MockProductRepository) Create(product *models.Product) error {
	arguments := r.Mock.Called(product)

	return arguments.Error(0)
}

func (r *MockProductRepository) Update(product *models.Product) (*models.Product, error) {
	arguments := r.Mock.Called(product)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	product = arguments.Get(0).(*models.Product)

	return product, arguments.Error(1)
}

func (r *MockProductRepository) Delete(ID uint) error {
	arguments := r.Mock.Called(ID)

	return arguments.Error(0)
}

func (r *MockProductRepository) FindAll() ([]models.Product, error) {
	arguments := r.Mock.Called()

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	products := arguments.Get(0).([]models.Product)

	return products, arguments.Error(1)
}

func (r *MockProductRepository) FindAllByUserID(userID uint) ([]models.Product, error) {
	arguments := r.Mock.Called(userID)

	if arguments.Get(0) == nil {
		return nil, arguments.Error(1)
	}

	products := arguments.Get(0).([]models.Product)

	return products, arguments.Error(1)
}
