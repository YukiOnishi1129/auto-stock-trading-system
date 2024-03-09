package seed

import (
	"context"
	"database/sql"
)

var (
	err error
)

func CreateInitData(ctx context.Context, db *sql.DB) error {

	//ur := repository.NewUserRepository(db)
	//uu := usecase.NewUserUsecase(ur)
	//
	//uc := NewCreateUserDataSeed(db, uu)
	//if err = uc.CreateUserData(ctx); err != nil {
	//	return err
	//}
	return nil
}
