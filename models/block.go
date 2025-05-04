package models

type BlockResponse struct {
    Number           uint64 `json:"number"`
    Hash             string `json:"hash"`
    ParentHash       string `json:"parent_hash"`
    Timestamp        uint64 `json:"timestamp"`
    Transactions     int    `json:"tx_count"`
    Miner            string `json:"miner"`
    GasUsed          uint64 `json:"gas_used"`
    GasLimit         uint64 `json:"gas_limit"`
    Size             uint64 `json:"size"`
    BaseFeePerGas    string `json:"base_fee_per_gas,omitempty"` 
    Difficulty       string `json:"difficulty"`
    TotalDifficulty  string `json:"total_difficulty,omitempty"`
}
