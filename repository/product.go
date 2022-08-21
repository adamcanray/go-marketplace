package repository

import (
	"fmt"
	"go-marketplace/entity"
	"go-marketplace/provider"
	"log"
)

func ProductAddRepository(name string, price int, stock int) error {
	db := provider.MysqlProvider()
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO product (name, price, stock) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatalln("[responsitoy.ProductAddRepository-db.Prepare]", err.Error())
	}
	_, errExec := stmt.Exec(name, price, stock)
	if errExec != nil {
		log.Fatalln("[responsitoy.ProductAddRepository-stmt.Exec]", err.Error())
	}
	defer stmt.Close()

	return err
}

func ProductGetListRepository() []entity.Product {
	db := provider.MysqlProvider()
	defer db.Close()

	result, err := db.Query("SELECT id, name, price, stock, created_at, updated_at FROM product")
	if err != nil {
		log.Fatalln("[responsitoy.ProductGetLastID-db.Query]", err.Error())
	}

	var list []entity.Product

	for result.Next() {
		var product entity.Product
		err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		list = append(list, product)
	}
	defer result.Close()

	return list
}

func ProductGetByIdRepository(id int) entity.Product {
	db := provider.MysqlProvider()
	defer db.Close()

	result, err := db.Query(fmt.Sprintf("SELECT id, name, price, stock, created_at, updated_at FROM product WHERE id = %d", id))
	if err != nil {
		log.Fatalln("[responsitoy.ProductGetById-db.Query]", err.Error())
	}

	var selectedProduct entity.Product

	for result.Next() {
		var product entity.Product
		err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		selectedProduct = product
	}
	defer result.Close()

	return selectedProduct
}

func ProductDeleteRepository(id int) error {
	db := provider.MysqlProvider()
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM product WHERE id = ?")
	if err != nil {
		log.Fatalln("[responsitoy.ProductDeleteRepository-db.Prepare]", err.Error())
	}
	_, errExec := stmt.Exec(id)
	if errExec != nil {
		log.Fatalln("[responsitoy.ProductDeleteRepository-stmt.Exec]", err.Error())
	}
	defer stmt.Close()

	return err
}

func ProductEditRepository(id int, name string, price int, stock int) error {
	db := provider.MysqlProvider()
	defer db.Close()

	stmt, err := db.Prepare("UPDATE product SET name = ?, price = ?, stock = ? WHERE id = ?")
	if err != nil {
		log.Fatalln("[responsitoy.ProductEditRepository-db.Prepare]", err.Error())
	}
	_, errExec := stmt.Exec(name, price, stock, id)
	if errExec != nil {
		log.Fatalln("[responsitoy.ProductEditRepository-stmt.Exec]", err.Error())
	}
	defer stmt.Close()

	return err
}
