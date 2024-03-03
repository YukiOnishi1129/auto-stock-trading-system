package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	"time"
)

func Init() (*sql.DB, error) {
	db, dbErr := connectDB()
	if dbErr != nil {
		fmt.Printf("Error connecting to DB\n")
		return nil, dbErr
	}
	return db, nil
}

func connectDB() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}
	c := mysql.Config{
		DBName:    os.Getenv("MYSQL_DATABASE"),
		User:      os.Getenv("MYSQL_USER"),
		Passwd:    os.Getenv("MYSQL_PASSWORD"),
		Net:       "tcp",
		Addr:      fmt.Sprintf("%s:%s", os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT")),
		ParseTime: true,
		Collation: "utf8mb4_unicode_ci",
		Loc:       jst,
	}

	db, dbErr := sql.Open("mysql", c.FormatDSN())
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}
