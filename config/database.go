package config

import(
	// "fmt"
	pg "github.com/go-pg/pg/v10"
)

const(

	DB_HOST = "localhost"
	DB_USER = "admin"
	DB_PASSWORD = "12345"
	DB_NAME = "myapp"
	DB_PORT = ":5432"
)

func DBConnection() *pg.DB {
	db := pg.Connect(&pg.Options{
	   	Addr:     DB_PORT,
	   	User:     DB_USER,
	   	Password: DB_PASSWORD,
	   	Database: DB_NAME,
	})

	return db
}