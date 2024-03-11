package repository

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/auto-stock-trading-system/batch-service/database/entity"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserRepositoryInterface interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
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

func (ur *UserRepository) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, qErr := entity.FindUser(ctx, ur.db, id)
	if qErr != nil {
		return nil, qErr
	}
	return user, nil
}

func (ur *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	err := user.Insert(ctx, ur.db, boil.Infer())
	if err != nil {
		return err
	}
	return nil
}
