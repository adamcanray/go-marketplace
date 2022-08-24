package main

import (
	"go-marketplace/handler"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	// load .env file when MODE is not development|staging|sandbox|production
	if os.Getenv("MODE") != "development" ||
		os.Getenv("MODE") != "staging" ||
		os.Getenv("MODE") != "sandbox" ||
		os.Getenv("MODE") != "production" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file", err)
		}
	}

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
