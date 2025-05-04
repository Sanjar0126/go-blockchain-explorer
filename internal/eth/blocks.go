package eth

import (
	"context"
	"go-blockchain-explorer/models"
	"math/big"
)

func GetLatestBlocks(count int) ([]models.BlockResponse, error) {
	var blocks []models.BlockResponse
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

		blocks = append(blocks, models.BlockResponse{
			Number:     block.NumberU64(),
			Hash:       block.Hash().Hex(),
			ParentHash: block.ParentHash().Hex(),
			Time:       block.Time(),
			TxCount:    len(block.Transactions()),
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
		Number:     block.NumberU64(),
		Hash:       block.Hash().Hex(),
		ParentHash: block.ParentHash().Hex(),
		Time:       block.Time(),
		TxCount:    len(block.Transactions()),
	}

	return resp, err
}
