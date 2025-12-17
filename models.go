package main

// Construct for Transaction
type Transaction struct {
	Currency	string
	Date        string
	Type        string
	Amount      string
	Price       string
	Fee		  	string
	Paid		string
}

// Construct for PaginationData
type PaginationData struct {
	Transactions []Transaction
	CurrentPage  int
	TotalPages   int
	PrevPage	 int
	NextPage	 int
	TotalItems	 int
	ActiveFilter string
}
