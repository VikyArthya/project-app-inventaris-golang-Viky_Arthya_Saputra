package repositories

import (
	"database/sql"
	"office-inventory/models"
)

type CategoryRepository struct {
	DB *sql.DB
}

// Mendapatkan semua kategori
func (repo *CategoryRepository) GetAllCategories() ([]models.Category, error) {
	rows, err := repo.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// Membuat kategori baru
func (repo *CategoryRepository) CreateCategory(category *models.Category) error {
	err := repo.DB.QueryRow("INSERT INTO categories (name) VALUES ($1) RETURNING id", category.Name).Scan(&category.ID)
	return err
}

// Mendapatkan kategori berdasarkan ID
func (repo *CategoryRepository) GetCategoryByID(id int) (*models.Category, error) {
	var category models.Category
	err := repo.DB.QueryRow("SELECT id, name FROM categories WHERE id = $1", id).Scan(&category.ID, &category.Name)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &category, err
}

// Mengupdate kategori berdasarkan ID
func (repo *CategoryRepository) UpdateCategory(category *models.Category) error {
	_, err := repo.DB.Exec("UPDATE categories SET name = $1 WHERE id = $2", category.Name, category.ID)
	return err
}

// Menghapus kategori berdasarkan ID
func (repo *CategoryRepository) DeleteCategory(id int) error {
	_, err := repo.DB.Exec("DELETE FROM categories WHERE id = $1", id)
	return err
}
