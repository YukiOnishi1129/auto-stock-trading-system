package usecase

import (
	"context"
	"github.com/YukiOnishi1129/auto-stock-trading-system/batch-service/database/entity"
	"github.com/YukiOnishi1129/auto-stock-trading-system/batch-service/infrastructure/mysql/repository"
	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
)

type UserUsecaseInterface interface {
	GetUsers(ctx context.Context) ([]*entity.User, error)
	GetUserByID(ctx context.Context, id string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}

type UserUsecase struct {
	ur *repository.UserRepository
}

func NewUserUsecase(ur *repository.UserRepository) *UserUsecase {
	return &UserUsecase{ur}
}

// GetUsers is a usecase to get all users
func (uu *UserUsecase) GetUsers(ctx context.Context) ([]*entity.User, error) {
	users, rErr := uu.ur.GetUsers(ctx)
	if rErr != nil {
		return nil, rErr
	}

	return users, nil
}

func (uu *UserUsecase) GetUserByID(ctx context.Context, id string) (*entity.User, error) {
	user, rErr := uu.ur.GetUserByID(ctx, id)
	if rErr != nil {
		return nil, rErr
	}
	return user, nil
}

type CreateUserDTO struct {
	Email string
	Name  string
	Image null.String
}

func (uu *UserUsecase) CreateUser(ctx context.Context, createUserDTO CreateUserDTO) (*entity.User, error) {
	uuID, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	user := &entity.User{
		ID:    uuID.String(),
		Email: null.StringFrom(createUserDTO.Email),
		Name:  null.StringFrom(createUserDTO.Name),
		Image: createUserDTO.Image,
	}

	err = uu.ur.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res, err := uu.ur.GetUserByID(ctx, uuID.String())
	if err != nil {
		return nil, err
	}

	return res, nil
}
