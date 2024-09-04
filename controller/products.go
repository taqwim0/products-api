package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"toggle-features-api/model"
	"toggle-features-api/utils"

	"github.com/gorilla/mux"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	query := `INSERT INTO products (name, description, price, variety, rating, stock) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	err := utils.DB.QueryRow(query, product.Name, product.Description, product.Price, product.Variety, product.Rating, product.Stock).Scan(&product.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, err := utils.DB.Query(`
		SELECT id, name, description, price, variety, rating, stock FROM products
	`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var product model.Product
		err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Variety, &product.Rating, &product.Stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(products) < 1 {
		// Hardcoded products if database empty (testing purposes)
		mockProduct := []model.Product{
			{
				ID:          1,
				Name:        "Laptop",
				Description: "Laptop for school",
				Price:       12000000,
				Variety:     "Electronics",
				Rating:      4.8,
				Stock:       10,
			},
			{
				ID:          2,
				Name:        "Smartphone",
				Description: "Smartphone with updated features",
				Price:       10000000,
				Variety:     "Electronics",
				Rating:      4.5,
				Stock:       10,
			},
		}
		json.NewEncoder(w).Encode(mockProduct)
	}

	json.NewEncoder(w).Encode(products)
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var product model.Product
	err := utils.DB.QueryRow("SELECT id, name, description, price, variety, rating, stock FROM products WHERE id=$1", id).
		Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Variety, &product.Rating, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "product id not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var product model.Product
	json.NewDecoder(r.Body).Decode(&product)

	query := `UPDATE products SET name=$1, description=$2, price=$3, variety=$4, rating=$5, stock=$6 WHERE id=$7`
	_, err := utils.DB.Exec(query, product.Name, product.Description, product.Price, product.Variety, product.Rating, product.Stock, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	query := `DELETE FROM products WHERE id=$1`
	_, err := utils.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
