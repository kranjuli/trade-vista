package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

func readCSV(filePath string) ([]Transaction, error) {
	// Get absolute file path
	path, err := filepath.Abs(filePath)
	if err != nil {
		return nil, err
	}
	// 1, open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %v", err)
	}
	// 2, ensure the file is closed after reading
	defer file.Close()

	// 3, create a new CSV reader
	reader := csv.NewReader(file)

	// 4, read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	// 5, parse records into Transaction structs
	var transactions []Transaction
	for i := 1; i < len(records); i++ {
		if len((records[i])) < 5 {
			continue
		}
		transactions = append(transactions, Transaction{
			Currency: records[i][4],
			Date:    records[i][1] + " / " + records[i][2],
			Type:    records[i][3],
			Amount:  records[i][5],
			Price:   records[i][7],
			Fee:  	 records[i][11],
			Paid:    records[i][9],
		})
	}
	// 6, return the slice of transactions
	return transactions, nil
}

// paginator logic
func paginate(transactions []Transaction, page, pageSize int) PaginationData {
	// Calculate pagination
	totalPages := (len(transactions) + pageSize - 1) / pageSize
	if page < 1 {
		page = 1
	}

	if page > totalPages {
		page = totalPages
	}

	startIndex := (page - 1) * pageSize
	endIndex := startIndex + pageSize

	if endIndex > len(transactions) {
		endIndex = len(transactions)
	}

	paginatedTransactions := transactions[startIndex:endIndex]

	return PaginationData{
		Transactions: paginatedTransactions,
		CurrentPage:  page,
		TotalPages:   totalPages,
		PrevPage:	 page - 1,
		NextPage:	 page + 1,
		TotalItems:	 len(transactions),
	}
}