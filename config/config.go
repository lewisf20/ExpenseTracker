package config

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func GetDatabase() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DBUSER")
	dbPass := os.Getenv("DBPASS")
	dbName := os.Getenv("DBNAME")

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		return nil, err
	}

	return db, nil
}
