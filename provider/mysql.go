package provider

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func MysqlProvider() *sql.DB {
	// References:
	// - ?parseTime=true -- https://stackoverflow.com/a/45040724/11587161
	// - docker.for.mac.localhost -- https://stackoverflow.com/a/52504428/11587161
	db, err := sql.Open("mysql", "root:root@tcp(docker.for.mac.localhost:3306)/go-mysql?parseTime=true")
	if err != nil {
		log.Fatalln("[provider.MysqlProvider]", err.Error())
	}
	// defer db.Close()
	return db
}
