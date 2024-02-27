package seed

import (
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type createUserData struct {
	ID       int64
	Name     string
	Email    string
	Password string
}

func CreateUserData(con *sql.DB) error {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	insertDatas := []createUserData{
		{
			ID:       1,
			Name:     "taro",
			Email:    "taro@gmail.com",
			Password: string(hashPassword),
		},
		{
			ID:       2,
			Name:     "jiro",
			Email:    "jiro@gmail.com",
			Password: string(hashPassword),
		},
		{
			ID:       3,
			Name:     "hanako",
			Email:    "hanako@gmail.com",
			Password: string(hashPassword),
		},
	}

	insSql := "INSERT INTO users (id, name, email, password) VALUES "
	for _, insertData := range insertDatas {
		insSql += fmt.Sprintf("('%v', '%v', '%v', '%v'),", insertData.ID, insertData.Name, insertData.Email, insertData.Password)
	}
	insSql = insSql[:len(insSql)-1]

	ins, err := con.Prepare(insSql)
	if err != nil {
		return err
	}
	defer func() {
		err := ins.Close()
		if err != nil {
			return
		}
	}()

	_, err = ins.Exec()
	if err != nil {
		return err
	}

	get, getErr := con.Prepare("SELECT MAX(id) FROM users;")
	if getErr != nil {
		return getErr
	}
	defer func() {
		err := get.Close()
		if err != nil {
			return
		}
	}()

	_, err = get.Exec()
	if err != nil {
		return err
	}
	fmt.Println("create data at users table")
	return nil
}
