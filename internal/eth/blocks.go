package eth

import (
	"context"
	"go-blockchain-explorer/models"
	"math/big"
)

func GetLatestBlocks(count int) ([]*models.BlockResponse, error) {
	var blocks []*models.BlockResponse
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	current := header.Number
	for i := 0; i < count; i++ {
		block, err := client.BlockByNumber(context.Background(), current)
		if err != nil {
			return nil, err
		}

		blocks = append(blocks, &models.BlockResponse{
			Number:       block.NumberU64(),
			Hash:         block.Hash().Hex(),
			ParentHash:   block.ParentHash().Hex(),
			Timestamp:    block.Time(),
			Transactions: len(block.Transactions()),
			Miner:        block.Coinbase().Hex(),
			GasUsed:      block.GasUsed(),
			GasLimit:     block.GasLimit(),
			Size:         uint64(block.Size()),

			Difficulty:      block.Difficulty().String(),
			TotalDifficulty: block.Difficulty().String(),
			BaseFeePerGas: func() string {
				if block.BaseFee() != nil {
					return block.BaseFee().String()
				}
				return ""
			}(),
		})

		current = new(big.Int).Sub(current, big.NewInt(1))
	}

	return blocks, nil
}

func GetBlockByNumber(num string) (*models.BlockResponse, error) {
	n, ok := new(big.Int).SetString(num, 10)
	if !ok {
		return nil, nil
	}

	block, err := client.BlockByNumber(context.Background(), n)
	if err != nil {
		return nil, err
	}

	resp := &models.BlockResponse{
		Number:       block.NumberU64(),
		Hash:         block.Hash().Hex(),
		ParentHash:   block.ParentHash().Hex(),
		Timestamp:    block.Time(),
		Transactions: len(block.Transactions()),
		Miner:        block.Coinbase().Hex(),
		GasUsed:      block.GasUsed(),
		GasLimit:     block.GasLimit(),
		Size:         uint64(block.Size()),

		Difficulty:      block.Difficulty().String(),
		TotalDifficulty: block.Difficulty().String(),
		BaseFeePerGas: func() string {
			if block.BaseFee() != nil {
				return block.BaseFee().String()
			}
			return ""
		}(),
	}

	return resp, err
}
