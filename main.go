package main

import (
	"html/template"
	"net/http"
	"fmt"
	"path/filepath"
	"bytes"
	"log"
)


func main() {
	// Set up HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Ignore favicon requests
		if r.URL.Path == "/favicon.ico" {
        return
    	}
		csvPath := filepath.Join("transactions", "2025.csv")
		transactions, err := readCSV(csvPath)
		if err != nil {
			http.Error(w, "Unable to read CSV file", http.StatusInternalServerError)
			return
		}
		// Filter
		filterType := r.URL.Query().Get("type")

		if filterType != "" {
			var filtered []Transaction
			for _, tx := range transactions {
				if tx.Type == filterType {
					filtered = append(filtered, tx)
				}
			}
			transactions = filtered
		} else {
			transactions = transactions
		}

		// Pagination logic
		page := 1
		fmt.Sscanf(r.URL.Query().Get("page"), "%d", &page)
		data := paginate(transactions, page, 15)
		data.ActiveFilter = filterType

		// Define HTML template
		tmplPath := filepath.Join("templates", "trading.html")
		tmpl, err := template.ParseFiles(tmplPath)
		if err != nil {
			log.Printf("ParseFiles error: %v", err)
			http.Error(w, "Unable to load trading.html", http.StatusInternalServerError)
			return
		}
		
		// Render template with transactions data
		var buf bytes.Buffer
		err = tmpl.Execute(&buf, data)

		if err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, "Unable to render trading.html", http.StatusInternalServerError)
			return
		}
		buf.WriteTo(w)
	})

	// Start the server
	fmt.Println("Server is running at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
