package repository

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/auto-stock-trading-system/user-service/database/entity"
)

type UserRepositoryInterface interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users, qErr := entity.Users().All(ctx, ur.db)
	if qErr != nil {
		return nil, qErr
	}
	return users, nil
}
