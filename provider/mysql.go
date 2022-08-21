package provider

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlProvider() *sql.DB {
	// ?parseTime=true reference: https://stackoverflow.com/a/45040724/11587161
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1)/go-mysql?parseTime=true")
	if err != nil {
		log.Fatalln("[provider.MysqlProvider]", err.Error())
	}
	// defer db.Close()
	return db
}
