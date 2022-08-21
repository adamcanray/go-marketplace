package main

import (
	"go-marketplace/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/admin/product", handler.ProductGetHandler)
	mux.HandleFunc("/admin/product/add", handler.ProductAddHandler)
	mux.HandleFunc("/admin/product/add/process", handler.ProductAddProcessHandler)
	mux.HandleFunc("/admin/product/delete/process", handler.ProductDeleteProcessHandler)
	mux.HandleFunc("/admin/product/edit", handler.ProductEditHandler)
	mux.HandleFunc("/admin/product/edit/process", handler.ProductEditProcessHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting web on port 8090")

	err := http.ListenAndServe(":8090", mux)
	log.Fatal(err)
}
