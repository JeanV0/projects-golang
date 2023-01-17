package db

import (
	"fmt"
	dotenv "marketplace/config/services/Dotenv"

	dbx "github.com/go-ozzo/ozzo-dbx"
	_ "github.com/lib/pq"
)

var DbInstance *dbx.DB

func GetConnectionDB() error {

	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",

		dotenv.MyEnvironmentApp.Database_user,
		dotenv.MyEnvironmentApp.Database_password,
		dotenv.MyEnvironmentApp.Database_host,
		dotenv.MyEnvironmentApp.Database_port,
		dotenv.MyEnvironmentApp.Database_DbName)

	db, err := dbx.Open("postgres", dsn)

	DbInstance = db

	if err != nil {
		return err
	}

	return nil
}
