package handlers

import (
	"encoding/json"
	"net/http"
	"office-inventory/models"
	"office-inventory/repositories"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ItemHandler struct {
	Repo *repositories.ItemRepository
}

// GET /api/items - Mendapatkan semua item inventaris
func (h *ItemHandler) GetAllItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.Repo.GetAllItems()
	if err != nil {
		http.Error(w, "Unable to retrieve items", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// POST /api/items - Menambahkan item baru ke inventaris
func (h *ItemHandler) CreateItem(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := h.Repo.CreateItem(&item); err != nil {
		http.Error(w, "Unable to create item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(item)
}

// GET /api/items/{id} - Mendapatkan detail item berdasarkan id
func (h *ItemHandler) GetItemByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	item, err := h.Repo.GetItemByID(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

// PUT /api/items/{id} - Mengedit item berdasarkan id
func (h *ItemHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var item models.Item
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	item.ID = id

	if err := h.Repo.UpdateItem(&item); err != nil {
		http.Error(w, "Unable to update item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(item)
}

// DELETE /api/items/{id} - Menghapus item berdasarkan id
func (h *ItemHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	if err := h.Repo.DeleteItem(id); err != nil {
		http.Error(w, "Unable to delete item", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET /api/items/replacement-needed - Mendapatkan daftar item yang perlu diganti
func (h *ItemHandler) GetReplacementNeededItems(w http.ResponseWriter, r *http.Request) {
	items, err := h.Repo.GetReplacementNeededItems()
	if err != nil {
		http.Error(w, "Unable to retrieve replacement-needed items", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GET /api/items/investment - Mendapatkan nilai investasi total dengan depresiasi
func (h *ItemHandler) GetTotalInvestment(w http.ResponseWriter, r *http.Request) {
	totalInvestment, err := h.Repo.GetTotalInvestment()
	if err != nil {
		http.Error(w, "Unable to retrieve total investment", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(totalInvestment)
}

// GET /api/items/investment/{id} - Mendapatkan nilai depresiasi dan investasi untuk item tertentu
func (h *ItemHandler) GetItemInvestment(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	investment, err := h.Repo.GetItemInvestment(id)
	if err != nil {
		http.Error(w, "Item not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(investment)
}
