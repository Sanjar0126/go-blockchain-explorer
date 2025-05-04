package models

type TransactionResponse struct {
	Hash  string `json:"hash"`
	From  string `json:"from"`
	To    string `json:"to"`
	Value string `json:"value"`
	Gas   uint64 `json:"gas"`
}

type TransactionListResponse struct {
	Items []TransactionResponse `json:"items"`
	Count int                   `json:"count"`
	Page  int                   `json:"page"`
}
