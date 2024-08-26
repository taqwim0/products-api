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

func CreateFeatureToggle(w http.ResponseWriter, r *http.Request) {
	var featureToggle model.FeatureToggle
	json.NewDecoder(r.Body).Decode(&featureToggle)

	query := `INSERT INTO toggle_features_api (name, description, enabled) VALUES ($1, $2, $3) RETURNING id`
	err := utils.DB.QueryRow(query, featureToggle.Name, featureToggle.Description, featureToggle.Enabled).Scan(&featureToggle.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(featureToggle)
}

func GetFeatureToggles(w http.ResponseWriter, r *http.Request) {
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page := 1
	limit := 5

	var err error
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			http.Error(w, "Invalid page number", http.StatusBadRequest)
			return
		}
	}
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			http.Error(w, "Invalid limit number", http.StatusBadRequest)
			return
		}
	}

	offset := (page - 1) * limit

	rows, err := utils.DB.Query(`
		SELECT id, name, description, enabled FROM toggle_features_api
		ORDER BY id
		LIMIT $1 OFFSET $2
	`, limit, offset)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var featureToggles []model.FeatureToggle
	for rows.Next() {
		var toggle model.FeatureToggle
		err := rows.Scan(&toggle.ID, &toggle.Name, &toggle.Description, &toggle.Enabled)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		featureToggles = append(featureToggles, toggle)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var totalRecords int
	err = utils.DB.QueryRow("SELECT COUNT(*) FROM toggle_features_api").Scan(&totalRecords)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	totalPages := (totalRecords + limit - 1) / limit

	response := map[string]interface{}{
		"page":         page,
		"limit":        limit,
		"totalPages":   totalPages,
		"totalRecords": totalRecords,
		"data":         featureToggles,
		"totalData":    len(featureToggles),
	}

	json.NewEncoder(w).Encode(response)
}

func GetFeatureToggleByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var toggle model.FeatureToggle
	err := utils.DB.QueryRow("SELECT id, name, description, enabled FROM toggle_features_api WHERE id=$1", id).
		Scan(&toggle.ID, &toggle.Name, &toggle.Description, &toggle.Enabled)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "feature id not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(toggle)
}

func UpdateFeatureToggle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var featureToggle model.FeatureToggle
	json.NewDecoder(r.Body).Decode(&featureToggle)

	query := `UPDATE toggle_features_api SET name=$1, description=$2, enabled=$3 WHERE id=$4`
	_, err := utils.DB.Exec(query, featureToggle.Name, featureToggle.Description, featureToggle.Enabled, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(featureToggle)
}

func DeleteFeatureToggle(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	query := `DELETE FROM toggle_features_api WHERE id=$1`
	_, err := utils.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
