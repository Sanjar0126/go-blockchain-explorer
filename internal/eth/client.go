package eth

import (
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://mainnet.infura.io/v3/1fbaae955fe34d16b38d45dbd7ab0f8b")
	if err != nil {
		log.Fatal(err)
	}
}

func GetClient() *ethclient.Client {
	return client
}
