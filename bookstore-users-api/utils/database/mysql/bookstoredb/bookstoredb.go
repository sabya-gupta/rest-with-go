package bookstoredb

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func init() {
	var err error
	dataSrcName := "user:password@tcp(localhost)/bookstore"
	DBClient, err = sql.Open("mysql", dataSrcName)
	if err != nil {
		panic(err)
	}

	pingErr := DBClient.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	log.Println("DB is connected")
}
