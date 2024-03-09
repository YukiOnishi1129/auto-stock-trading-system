package seed

import (
	"context"
	"database/sql"
	"github.com/YukiOnishi1129/auto-stock-trading-system/batch-service/usecase"
)

type CreateUserDataSeedInterface interface {
	CreateUserData(db *sql.DB) error
}

type CreateUserDataSeed struct {
	db *sql.DB
	uu *usecase.UserUsecase
}

func NewCreateUserDataSeed(db *sql.DB, uu *usecase.UserUsecase) *CreateUserDataSeed {
	return &CreateUserDataSeed{db, uu}
}

func (cu *CreateUserDataSeed) CreateUserData(ctx context.Context) error {

	createUserDatas := []usecase.CreateUserDTO{
		{
			Name:  "taro",
			Email: "taro@gmail.com",
		},
		{
			Name:  "jiro",
			Email: "jiro@gmail.com",
		},
		{
			Name:  "hanako",
			Email: "hanako@gmail.com",
		},
	}

	for _, c := range createUserDatas {
		_, err := cu.uu.CreateUser(ctx, c)
		if err != nil {
			return err
		}
	}

	return nil
}
