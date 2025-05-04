package models

type BlockResponse struct {
	Number     uint64 `json:"number"`
	Hash       string `json:"hash"`
	ParentHash string `json:"parent_hash"`
	Time       uint64 `json:"timestamp"`
	TxCount    int    `json:"tx_count"`
}
