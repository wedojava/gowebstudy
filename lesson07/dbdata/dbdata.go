package dbdata

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlDeamonCode() (*sql.DB, error) {
	db, err := sql.Open("mysql", "go_web:go_web@tcp(127.0.0.1:3306)/go_web")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db, nil
}
