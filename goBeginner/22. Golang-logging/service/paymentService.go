package service

import (
	"book-store/collections"
	"book-store/repository"
	"context"
	"errors"
)

type PaymentMethodService interface {
	CreatePaymentMethod(ctx context.Context, method *collections.PaymentMethod) error
	ListPaymentMethods(ctx context.Context) ([]collections.PaymentMethod, error)
	GetPaymentMethod(ctx context.Context, id int) (*collections.PaymentMethod, error)
	UpdatePaymentMethod(ctx context.Context, method *collections.PaymentMethod) error
	DeletePaymentMethod(ctx context.Context, id int) error
}

type paymentMethodService struct {
	repo repository.PaymentMethodRepository
}

func NewPaymentMethodService(repo repository.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{repo}
}

func (s *paymentMethodService) CreatePaymentMethod(ctx context.Context, method *collections.PaymentMethod) error {
	if method.Name == "" {
		return errors.New("name is required")
	}
	if method.PhotoURL == "" {
		return errors.New("name is required")
	}
	return s.repo.Create(ctx, method)
}

func (s *paymentMethodService) ListPaymentMethods(ctx context.Context) ([]collections.PaymentMethod, error) {
	return s.repo.GetAll(ctx)
}

func (s *paymentMethodService) GetPaymentMethod(ctx context.Context, id int) (*collections.PaymentMethod, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *paymentMethodService) UpdatePaymentMethod(ctx context.Context, method *collections.PaymentMethod) error {
	return s.repo.Update(ctx, method)
}

func (s *paymentMethodService) DeletePaymentMethod(ctx context.Context, id int) error {
	return s.repo.Delete(ctx, id)
}
