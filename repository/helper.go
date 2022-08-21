package repository

import (
	"fmt"
	"go-marketplace/entity"
	"go-marketplace/provider"
	"log"
)

func HelperGetLastIDRepository(tableName string, columnName string) int {
	db := provider.MysqlProvider()
	defer db.Close()

	result, err := db.Query(fmt.Sprintf("SELECT id FROM %s WHERE %s=(SELECT MAX(%s) FROM %s)", tableName, columnName, columnName, tableName))
	if err != nil {
		log.Fatalln("[responsitoy.ProductGetLastID-db.Query]", err.Error())
	}

	var lastId int

	for result.Next() {
		var product entity.Product
		err = result.Scan(&product.ID)
		if err != nil {
			panic(err.Error())
		}
		lastId += product.ID
	}
	defer result.Close()

	return lastId
}
