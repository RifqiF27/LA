package orderrepository

import (
	"dashboard-ecommerce-team2/models"

	"github.com/stretchr/testify/mock"
)

type OrderRepositoryMock struct {
	mock.Mock
}

// GetAll implements OrderRepository.
func (o *OrderRepositoryMock) GetAll() ([]models.Order, error) {
	panic("unimplemented")
}

// GetByID implements OrderRepositoryMock.
func (o *OrderRepositoryMock) GetByID(id int) (*models.Order, error) {
	panic("unimplemented")
}

// UpdateStatus implements OrderRepositoryMock.
func (o *OrderRepositoryMock) UpdateStatus(id int, status string) error {
	panic("unimplemented")
}
