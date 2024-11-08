package repositories

import (
	"database/sql"
	"office-inventory/models"
)

type ItemRepository struct {
	DB *sql.DB
}

// Mendapatkan semua item inventaris
func (repo *ItemRepository) GetAllItems() ([]models.Item, error) {
	rows, err := repo.DB.Query(`
		SELECT id, name, photo, price, purchase_date, usage_days, category_id 
		FROM items
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Photo, &item.Price, &item.PurchaseDate, &item.UsageDays, &item.CategoryID); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Membuat item baru di inventaris
func (repo *ItemRepository) CreateItem(item *models.Item) error {
	err := repo.DB.QueryRow(`
		INSERT INTO items (name, photo, price, purchase_date, usage_days, category_id) 
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
	`, item.Name, item.Photo, item.Price, item.PurchaseDate, item.UsageDays, item.CategoryID).Scan(&item.ID)
	return err
}

// Mendapatkan item berdasarkan ID
func (repo *ItemRepository) GetItemByID(id int) (*models.Item, error) {
	var item models.Item
	err := repo.DB.QueryRow(`
		SELECT id, name, photo, price, purchase_date, usage_days, category_id 
		FROM items WHERE id = $1
	`, id).Scan(&item.ID, &item.Name, &item.Photo, &item.Price, &item.PurchaseDate, &item.UsageDays, &item.CategoryID)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return &item, err
}

// Mengupdate item berdasarkan ID
func (repo *ItemRepository) UpdateItem(item *models.Item) error {
	_, err := repo.DB.Exec(`
		UPDATE items SET name = $1, photo = $2, price = $3, purchase_date = $4, usage_days = $5, category_id = $6 
		WHERE id = $7
	`, item.Name, item.Photo, item.Price, item.PurchaseDate, item.UsageDays, item.CategoryID, item.ID)
	return err
}

// Menghapus item berdasarkan ID
func (repo *ItemRepository) DeleteItem(id int) error {
	_, err := repo.DB.Exec("DELETE FROM items WHERE id = $1", id)
	return err
}

// Mendapatkan daftar item yang perlu diganti (lebih dari 100 hari penggunaan)
func (repo *ItemRepository) GetReplacementNeededItems() ([]models.Item, error) {
	rows, err := repo.DB.Query(`
		SELECT id, name, photo, price, purchase_date, usage_days, category_id 
		FROM items WHERE usage_days > 100
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []models.Item
	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.ID, &item.Name, &item.Photo, &item.Price, &item.PurchaseDate, &item.UsageDays, &item.CategoryID); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}

// Menghitung total investasi dengan depresiasi saldo menurun
func (repo *ItemRepository) GetTotalInvestment() (float64, error) {
	var totalInvestment float64
	rows, err := repo.DB.Query(`
		SELECT price, usage_days FROM items
	`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	for rows.Next() {
		var price float64
		var usageDays int
		if err := rows.Scan(&price, &usageDays); err != nil {
			return 0, err
		}

		// Menghitung depresiasi saldo menurun (contoh depresiasi 10%)
		depreciationRate := 0.10
		depreciatedValue := price
		for i := 0; i < usageDays/365; i++ {
			depreciatedValue -= depreciatedValue * depreciationRate
		}
		totalInvestment += depreciatedValue
	}

	return totalInvestment, nil
}

// Mendapatkan nilai depresiasi dan investasi dari item tertentu
func (repo *ItemRepository) GetItemInvestment(id int) (float64, error) {
	var price float64
	var usageDays int
	err := repo.DB.QueryRow("SELECT price, usage_days FROM items WHERE id = $1", id).Scan(&price, &usageDays)
	if err != nil {
		return 0, err
	}

	// Menghitung depresiasi saldo menurun (contoh depresiasi 10%)
	depreciationRate := 0.10
	depreciatedValue := price
	for i := 0; i < usageDays/365; i++ {
		depreciatedValue -= depreciatedValue * depreciationRate
	}

	return depreciatedValue, nil
}
