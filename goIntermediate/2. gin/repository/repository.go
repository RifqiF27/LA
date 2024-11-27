package repository

import (
	"database/sql"

	"go.uber.org/zap"
)

type AllRepository struct {
	ShippingRepo ShippingRepoInterface
}

func NewAllRepository(db *sql.DB, log *zap.Logger) AllRepository {
	return AllRepository{
		ShippingRepo: NewShippingRepository(db, log),
	}
}
