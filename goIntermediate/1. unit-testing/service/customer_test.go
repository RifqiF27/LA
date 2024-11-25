package service

import (
	"be-golang-chapter-36-implem/helper"
	"be-golang-chapter-36-implem/model"
	"be-golang-chapter-36-implem/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	// logger := zap.NewNop()

	mockRepo := repository.CustomerRepositoryMock{}
	allRepo := repository.AllRepository{
		CustomerRep: &mockRepo,
	}

	customerService := NewCustomerService(allRepo, nil)

	// Mock data
	validCustomer := model.Customer{
		Email:    "test@example.com",
		Password: "validpassword",
	}

	hashedPassword, _ := helper.HashPassword("validpassword")
	storedCustomer := &model.Customer{
		Email:    "test@example.com",
		Password: hashedPassword,
	}

	hashedPasswordFailed, _ := helper.HashPassword("1231313")
	storedCustomerFailed := &model.Customer{
		Email:    "test@example.com",
		Password: hashedPasswordFailed,
	}

	t.Run("invalid email", func(t *testing.T) {
		// Setup mock
		// mockRepo.On("GetByCondition", validCustomer).Return(nil, errors.New("not found"))
		mockRepo.On("GetByCondition", validCustomer).Once().Return(nil, errors.New("not found"))
		// Test
		result, err := customerService.Login(validCustomer)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid email", err.Error())

		mockRepo.AssertExpectations(t)
	})

	t.Run("successful login", func(t *testing.T) {
		// mockRepo.ExpectedCalls = nil // untuk reset mock setup sebelumnya
		// Setup mock
		// mockRepo.On("GetByCondition", validCustomer).Return(storedCustomer, nil)
		mockRepo.On("GetByCondition", validCustomer).Once().Return(storedCustomer, nil) // untuk digunakan sekali

		// Test
		result, err := customerService.Login(validCustomer)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, validCustomer.Email, result.Email)

		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid password", func(t *testing.T) {
		// mockRepo.ExpectedCalls = nil // untuk reset mock setup sebelumnya
		// Setup mock
		// mockRepo.On("GetByCondition", validCustomer).Return(storedCustomerFailed, nil)
		mockRepo.On("GetByCondition", validCustomer).Once().Return(storedCustomerFailed, nil) // untuk digunakan sekali

		// Test
		result, err := customerService.Login(validCustomer)
		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "invalid password", err.Error())

		mockRepo.AssertExpectations(t)
	})
}

func TestCreate(t *testing.T) {
	// Mock repository
	mockRepo := repository.CustomerRepositoryMock{}
	allRepo := repository.AllRepository{
		CustomerRep: &mockRepo,
	}

	customerService := NewCustomerService(allRepo, nil)

	t.Run("successful create", func(t *testing.T) {
		newCustomer := &model.Customer{
			Name:     "John Doe",
			Email:    "john@example.com",
			Phone:    "123456789",
			Password: "plaintextpassword",
		}

		mockRepo.On("Create", mock.Anything).Once().Return(nil)

		err := customerService.Create(newCustomer)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to hash password", func(t *testing.T) {
		newCustomer := &model.Customer{
			Name:     "John Doe",
			Email:    "john@example.com",
			Phone:    "123456789",
			Password: "",
		}

		err := customerService.Create(newCustomer)

		assert.Error(t, err)
		assert.Equal(t, "failed to hash password", err.Error())
	})
}

func TestGetAll(t *testing.T) {
	// Mock repository
	mockRepo := repository.CustomerRepositoryMock{}
	allRepo := repository.AllRepository{
		CustomerRep: &mockRepo,
	}

	customerService := NewCustomerService(allRepo, nil)

	t.Run("successfully retrieve all customers", func(t *testing.T) {
		customers := []model.Customer{
			{
				Name:     "John Doe",
				Email:    "john@example.com",
				Phone:    "123456789",
				Password: "hashedPassword",
			},
			{
				Name:     "Jane Doe",
				Email:    "jane@example.com",
				Phone:    "987654321",
				Password: "hashedPassword",
			},
		}

		mockRepo.On("GetAll").Once().Return(&customers, nil)

		result, err := customerService.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, len(customers), len(*result))
		mockRepo.AssertExpectations(t)
	})

	t.Run("failed to retrieve customers", func(t *testing.T) {
		mockRepo.On("GetAll").Once().Return(nil, errors.New("database error"))

		result, err := customerService.GetAll()

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "database error", err.Error())
		mockRepo.AssertExpectations(t)
	})
}

