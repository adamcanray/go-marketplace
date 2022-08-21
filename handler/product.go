package handler

import (
	"fmt"
	"go-marketplace/repository"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func ProductAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		tmpl, err := template.ParseFiles(path.Join("views/admin", "product-add.html"), path.Join("views/admin", "layout.html"))
		if err != nil {
			log.Println("[handler.ProductAddHandler-template.ParseFiles]", err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			log.Println("[handler.ProductAddHandler-tmpl.Execute]", err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		log.Println("[handler.ProductAddHandler-method=GET]")
		return
	}

	log.Println("[handler.ProductAddHandler]")
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	return
}

func ProductAddProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		name := r.Form.Get("name")
		price, err := strconv.Atoi(r.Form.Get("price"))
		if err != nil {
			log.Println("[handler.ProductAddProcessHandler-strconv.Atoi#price]", err)
			return
		}
		stock, err := strconv.Atoi(r.Form.Get("stock"))
		if err != nil {
			log.Println("[handler.ProductAddProcessHandler-strconv.Atoi#stock]", err)
			return
		}

		err = repository.ProductAddRepository(name, price, stock)
		if err != nil {
			log.Fatalln("[handler.ProductAddProcessHandler-repository.ProductAddRepository]", err.Error())
		}

		lastId := repository.HelperGetLastIDRepository("product", "id")

		log.Println("[handler.ProductAddProcessHandler-method=POST#redirect]")
		http.Redirect(w, r, fmt.Sprintf("/admin/product?last_product_id=%d", lastId), http.StatusSeeOther)
		return
	}

	log.Println("[handler.ProductAddProcessHandler-method=GET#detail]")
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	return
}

func ProductGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		page := r.URL.Query().Get("page")
		contentPerPage := r.URL.Query().Get("content_per_page")
		orderBy := r.URL.Query().Get("order_by")
		order := r.URL.Query().Get("order")

		idNumb, err := strconv.Atoi(id)

		if err != nil || idNumb < 1 {
			tmpl, err := template.ParseFiles(path.Join("views/admin", "product.html"), path.Join("views/admin", "layout.html"))
			if err != nil {
				log.Println("[handler.ProductGetHandler-template.ParseFiles#list]", err)
				http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
				return
			}

			pageNumb, err := strconv.Atoi(page)
			contentPerPageNumb, err := strconv.Atoi(contentPerPage)

			data := repository.ProductGetListRepository(orderBy, order, pageNumb, contentPerPageNumb)
			err = tmpl.Execute(w, data)
			if err != nil {
				log.Println("[handler.ProductGetHandler-tmpl.Execute#list]", err)
				http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
				return
			}

			log.Println("[handler.ProductGetHandler-method=GET#list]")
			return
		} else if err == nil || idNumb >= 1 {
			tmpl, err := template.ParseFiles(path.Join("views/admin", "product-detail.html"), path.Join("views/admin", "layout.html"))
			if err != nil {
				log.Println("[handler.ProductGetHandler-template.ParseFiles#detail]", err)
				http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
				return
			}

			data := repository.ProductGetByIdRepository(idNumb)

			err = tmpl.Execute(w, data)
			if err != nil {
				log.Println("[handler.ProductGetHandler-tmpl.Execute#detail]", err)
				http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
				return
			}

			log.Println("[handler.ProductGetHandler-method=GET#detail]")
			return
		}

		log.Println("[handler.ProductGetHandler-method=GET]")
		http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
		return
	}

	log.Println("[handler.ProductGetHandler]")
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	return
}

func ProductDeleteProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		idNumb, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			log.Fatalln("[handler.ProductDeleteProcessHandler-strconv.Atoi#id]", err.Error())
			return
		}

		err = repository.ProductDeleteRepository(idNumb)
		if err != nil {
			log.Fatalln("[handler.ProductDeleteProcessHandler-repository.ProductDeleteRepository]", err.Error())
		}

		log.Println("[handler.ProductDeleteProcessHandler-method=GET]")
		http.Redirect(w, r, "/admin/product", http.StatusSeeOther)
		return
	}

	log.Println("[handler.ProductDeleteProcessHandler]")
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	return
}

func ProductEditHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := r.URL.Query().Get("id")
		idNumb, err := strconv.Atoi(id)
		if err != nil {
			log.Fatalln("[handler.ProductEditHandler-strconv.Atoi#id]", err.Error())
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles(path.Join("views/admin", "product-edit.html"), path.Join("views/admin", "layout.html"))
		if err != nil {
			log.Println("[handler.ProductEditHandler-template.ParseFiles]", err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		product := repository.ProductGetByIdRepository(idNumb)
		if err != nil {
			log.Fatalln("[handler.ProductEditHandler-repository.ProductGetByIdRepository]", err.Error())
		}

		err = tmpl.Execute(w, product)
		if err != nil {
			log.Println("[handler.ProductEditHandler-tmpl.Execute]", err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		log.Println("[handler.ProductEditHandler-method=GET]")
		return
	}

	log.Println("[handler.ProductEditHandler]")
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	return
}

func ProductEditProcessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error is happening, keep calm.", http.StatusInternalServerError)
			return
		}

		id, err := strconv.Atoi(r.Form.Get("id"))
		if err != nil {
			log.Println("[handler.ProductEditProcessHandler-strconv.Atoi#id]", err)
			return
		}
		name := r.Form.Get("name")
		price, err := strconv.Atoi(r.Form.Get("price"))
		if err != nil {
			log.Println("[handler.ProductEditProcessHandler-strconv.Atoi#price]", err)
			return
		}
		stock, err := strconv.Atoi(r.Form.Get("stock"))
		if err != nil {
			log.Println("[handler.ProductEditProcessHandler-strconv.Atoi#stock]", err)
			return
		}

		err = repository.ProductEditRepository(id, name, price, stock)
		if err != nil {
			log.Fatalln("[handler.ProductEditProcessHandler-repository.ProductEditRepository]", err.Error())
		}

		log.Println("[handler.ProductEditProcessHandler-method=POST]")
		http.Redirect(w, r, "/admin/product", http.StatusSeeOther)
		return
	}

	log.Println("[handler.ProductEditProcessHandler]")
	http.Error(w, "Error is happening, keep calm.", http.StatusBadRequest)
	return
}
