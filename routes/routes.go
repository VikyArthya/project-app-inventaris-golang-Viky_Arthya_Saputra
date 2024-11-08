package routes

import (
	"office-inventory/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes(categoryHandler *handlers.CategoryHandler, itemHandler *handlers.ItemHandler) *chi.Mux {
	r := chi.NewRouter()

	// Middleware
	r.Use(middleware.Logger)    // Logging
	r.Use(middleware.Recoverer) // Recover from panics
	r.Use(middleware.AllowContentType("application/json"))

	// Kategori Barang Routes
	r.Route("/api/categories", func(r chi.Router) {
		r.Get("/", categoryHandler.GetAllCategories)      // Mendapatkan semua kategori barang
		r.Post("/", categoryHandler.CreateCategory)       // Menambahkan kategori baru
		r.Get("/{id}", categoryHandler.GetCategoryByID)   // Mendapatkan detail kategori berdasarkan ID
		r.Put("/{id}", categoryHandler.UpdateCategory)    // Mengedit kategori berdasarkan ID
		r.Delete("/{id}", categoryHandler.DeleteCategory) // Menghapus kategori berdasarkan ID
	})

	// Manajemen Barang Inventaris Routes
	r.Route("/api/items", func(r chi.Router) {
		r.Get("/", itemHandler.GetAllItems)                                 // Mendapatkan semua barang inventaris
		r.Post("/", itemHandler.CreateItem)                                 // Menambahkan barang baru ke dalam inventaris
		r.Get("/{id}", itemHandler.GetItemByID)                             // Mendapatkan detail barang inventaris berdasarkan ID
		r.Put("/{id}", itemHandler.UpdateItem)                              // Mengedit barang inventaris berdasarkan ID
		r.Delete("/{id}", itemHandler.DeleteItem)                           // Menghapus barang inventaris berdasarkan ID
		r.Get("/replacement-needed", itemHandler.GetReplacementNeededItems) // Barang yang perlu diganti
		r.Get("/investment", itemHandler.GetTotalInvestment)                // Mendapatkan total investasi setelah depresiasi
		r.Get("/investment/{id}", itemHandler.GetItemInvestment)            // Mendapatkan nilai investasi barang tertentu berdasarkan ID
	})

	return r
}
