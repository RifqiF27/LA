package repository

import (
	"be-golang-chapter-36-implem/model"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	return db, mock
}

func TestCustomerRepository_GetByCondition_GORM(t *testing.T) {
	db, mock := setupTestDB(t)
	defer func() { _ = mock.ExpectationsWereMet() }()

	customerRepo := NewCustomerRepository(db, nil)

	t.Run("successfully get customer by email", func(t *testing.T) {
		customer := model.Customer{
			Email: "johndoe@example.com",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).
			AddRow(1, "John Doe", "johndoe@example.com", "123456789", "password123")

		mock.ExpectQuery(`SELECT id, name, email, phone, password FROM customers WHERE 1=1 AND email = ?`).
			WithArgs(customer.Email). // GORM secara default mengisi LIMIT dengan 1
			WillReturnRows(rows)

		result, err := customerRepo.GetByCondition(customer)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "johndoe@example.com", result.Email)
	})

	t.Run("successfully get customer by phone", func(t *testing.T) {
		customer := model.Customer{
			Phone: "123456789",
		}

		rows := sqlmock.NewRows([]string{"id", "name", "email", "phone", "password"}).
			AddRow(1, "John Doe", "johndoe@example.com", "123456789", "password123")

		mock.ExpectQuery(`SELECT id, name, email, phone, password FROM customers WHERE 1=1 AND phone = ?`).
			WithArgs(customer.Phone).
			WillReturnRows(rows)

		result, err := customerRepo.GetByCondition(customer)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, "John Doe", result.Name)
		assert.Equal(t, "123456789", result.Phone)
	})

	t.Run("customer not found", func(t *testing.T) {
		customer := model.Customer{
			Email: "unknown@example.com",
		}

		mock.ExpectQuery(`SELECT id, name, email, phone, password FROM customers WHERE 1=1 AND email = ?`).
			WithArgs(customer.Email).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := customerRepo.GetByCondition(customer)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, gorm.ErrRecordNotFound, err)
	})
}

func BenchmarkGetAll(b *testing.B) {
	// Mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		b.Fatalf("Failed to create mock DB: %s", err)
	}
	defer db.Close()

	// Mock data
	mockRows := sqlmock.NewRows([]string{"name", "email", "phone", "password"}).
		AddRow("John Doe", "john@example.com", "123456789", "hashedPassword").
		AddRow("Jane Doe", "jane@example.com", "987654321", "hashedPassword")

	// Tambahkan ekspektasi untuk setiap iterasi benchmark
	mock.MatchExpectationsInOrder(false) // Tidak peduli urutan ekspektasi

	// Repository setup
	logger := zap.NewNop()
	repo := NewCustomerRepository(db, logger)

	for i := 0; i < b.N; i++ {
		mock.ExpectQuery("SELECT name, email, phone, password FROM customers").
			WillReturnRows(mockRows)

		customers, err := repo.GetAll()
		if err != nil {
			b.Errorf("Unexpected error: %v", err)
		}

		if customers == nil {
			b.Errorf("Expected value not to be nil")
		}
	}
}


func TestCustomerRepository_Create(t *testing.T) {
	db, mock := setupTestDB(t)
	defer db.Close()

	repo := NewCustomerRepository(db, zap.NewNop())

	t.Run("successfully create customer", func(t *testing.T) {
		customer := &model.Customer{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Phone:    "123456789",
			Password: "password123",
		}

		// Mock ekspektasi query
		mock.ExpectQuery(`INSERT INTO customers`).
			WithArgs(customer.Name, customer.Email, customer.Phone, customer.Password).
			WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

		err := repo.Create(customer)

		assert.NoError(t, err)
		assert.Equal(t, 1, customer.ID)
	})

	t.Run("failed to create customer due to query error", func(t *testing.T) {
		customer := &model.Customer{
			Name:     "John Doe",
			Email:    "johndoe@example.com",
			Phone:    "123456789",
			Password: "password123",
		}

		// Mock ekspektasi query error
		mock.ExpectQuery(`INSERT INTO customers`).
			WithArgs(customer.Name, customer.Email, customer.Phone, customer.Password).
			WillReturnError(sql.ErrConnDone)

		err := repo.Create(customer)

		assert.Error(t, err)
		assert.Equal(t, sql.ErrConnDone, err)
	})
}


func TestCustomerRepository_GetAll(t *testing.T) {
	db, mock := setupTestDB(t)
	defer db.Close()

	repo := NewCustomerRepository(db, zap.NewNop())

	t.Run("successfully get all customers", func(t *testing.T) {
		mockRows := sqlmock.NewRows([]string{"name", "email", "phone", "password"}).
			AddRow("John Doe", "johndoe@example.com", "123456789", "password123").
			AddRow("Jane Doe", "janedoe@example.com", "987654321", "password456")

		mock.ExpectQuery(`SELECT name, email, phone, password FROM customers`).
			WillReturnRows(mockRows)

		customers, err := repo.GetAll()

		assert.NoError(t, err)
		assert.NotNil(t, customers)
		assert.Len(t, *customers, 2)
		assert.Equal(t, "John Doe", (*customers)[0].Name)
	})

	t.Run("failed to get all customers due to query error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT name, email, phone, password FROM customers`).
			WillReturnError(sql.ErrConnDone)

		customers, err := repo.GetAll()

		assert.Error(t, err)
		assert.Nil(t, customers)
		assert.Equal(t, sql.ErrConnDone, err)
	})
}
