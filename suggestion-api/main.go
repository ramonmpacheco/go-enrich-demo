package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"github.com/pacheco/go-enrich-demo/suggestion-api/models"
)

var db *sql.DB

func main() {
	var err error
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		// Default for local testing if not running in compose
		connStr = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/suggestions", getSuggestions)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8082"
	}
	log.Printf("Starting suggestion-api on port %s", port)
	http.ListenAndServe(":"+port, r)
}

func getSuggestions(w http.ResponseWriter, r *http.Request) {
	customerCode := r.URL.Query().Get("customer_code")
	if customerCode == "" {
		http.Error(w, "customer_code is required", http.StatusBadRequest)
		return
	}

	rows, err := db.Query("SELECT restaurant_name, promotion_description FROM suggestions WHERE customer_code = $1", customerCode)
	if err != nil {
		log.Printf("Error querying DB: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var suggestions []models.Suggestion
	for rows.Next() {
		var s models.Suggestion
		if err := rows.Scan(&s.Name, &s.Promotion); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		suggestions = append(suggestions, s)
	}

	if suggestions == nil {
		suggestions = []models.Suggestion{}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(suggestions)
}
