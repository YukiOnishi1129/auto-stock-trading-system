package seed

import "database/sql"

var (
	err error
)

func CreateInitData(con *sql.DB) error {
	if err = CreateUserData(con); err != nil {
		return err
	}
	return nil
}
