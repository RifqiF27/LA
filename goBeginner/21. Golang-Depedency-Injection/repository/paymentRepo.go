package repository

import (
	"book-store/collections"
	"context"
	"database/sql"
)

type PaymentMethodRepository interface {
	Create(ctx context.Context, method *collections.PaymentMethod) error
	GetAll(ctx context.Context) ([]collections.PaymentMethod, error)
	GetByID(ctx context.Context, id int) (*collections.PaymentMethod, error)
	Update(ctx context.Context, method *collections.PaymentMethod) error
	Delete(ctx context.Context, id int) error
}

type paymentMethodRepository struct {
	db *sql.DB
}

func NewPaymentMethodRepository(db *sql.DB) PaymentMethodRepository {
	return &paymentMethodRepository{db}
}

func (r *paymentMethodRepository) Create(ctx context.Context, method *collections.PaymentMethod) error {
	query := `
        INSERT INTO "Payment_Methods" (name, photo_url, is_active, created_at, updated_at)
        VALUES ($1, $2, $3, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `
	err := r.db.QueryRowContext(ctx, query, method.Name, method.PhotoURL, method.IsActive).
		Scan(&method.ID, &method.CreatedAt, &method.UpdatedAt)
	return err
}

func (r *paymentMethodRepository) GetAll(ctx context.Context) ([]collections.PaymentMethod, error) {
	return nil, nil

}

func (r *paymentMethodRepository) GetByID(ctx context.Context, id int) (*collections.PaymentMethod, error) {
	return nil, nil
}

func (r *paymentMethodRepository) Update(ctx context.Context, method *collections.PaymentMethod) error {
	return nil
}

func (r *paymentMethodRepository) Delete(ctx context.Context, id int) error {
	return nil
}
