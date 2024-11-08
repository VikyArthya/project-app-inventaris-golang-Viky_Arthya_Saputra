package main

import (
	"log"
	"net/http"
	"office-inventory/config"
	"office-inventory/handlers"
	"office-inventory/repositories"
	"office-inventory/routes"
)

func main() {
	db := config.ConnectDB()
	defer db.Close()

	categoryRepo := &repositories.CategoryRepository{DB: db}
	itemRepo := &repositories.ItemRepository{DB: db}

	categoryHandler := &handlers.CategoryHandler{Repo: categoryRepo}
	itemHandler := &handlers.ItemHandler{Repo: itemRepo}

	r := routes.Routes(categoryHandler, itemHandler)

	log.Println("Server berjalan di http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
