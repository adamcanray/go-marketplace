package repository

import (
	"fmt"
	"go-marketplace/entity"
	"go-marketplace/model"
	"go-marketplace/provider"
	"log"
	"math"
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

func ProductGetListRepository(order_by string, order string, page int, content_per_page int) model.ListProduct {
	db := provider.MysqlProvider()
	defer db.Close()

	// set default value
	if order_by == "" {
		order_by = "id"
	}
	if order == "" {
		order = "ASC"
	}
	if page == 0 {
		page = 1
	}
	if content_per_page == 0 {
		content_per_page = 5
	}

	page = page - 1

	offset := content_per_page * page

	result, err := db.Query(fmt.Sprintf("SELECT id, name, price, stock, created_at, updated_at FROM product ORDER BY %s %s LIMIT %d,%d", order_by, order, offset, content_per_page))
	if err != nil {
		log.Fatalln("[responsitoy.ProductGetLastID-db.Query#list]", err.Error())
	}

	var countRows int
	err = db.QueryRow("SELECT COUNT(*) FROM product").Scan(&countRows)
	if err != nil {
		log.Fatalln("[responsitoy.ProductGetLastID-db.Query#count-rows]", err.Error())
	}

	pages := math.Ceil(float64(countRows) / float64(content_per_page))

	var list model.ListProduct
	var products []entity.Product

	for result.Next() {
		var product entity.Product
		err = result.Scan(&product.ID, &product.Name, &product.Price, &product.Stock, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			panic(err.Error())
		}
		products = append(products, product)
	}
	defer result.Close()

	list.Data.Products = products
	list.Meta.Page = page + 1
	list.Meta.ContentPerPage = content_per_page
	list.Meta.MaxPage = int(pages)
	list.Meta.Total = countRows

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
