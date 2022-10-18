package automator

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func Auto() {
	client, err := ethclient.Dial("https://decoded.wtf")
	if err != nil {
		log.Fatal(err)
	}

	block, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Latest block->", block)
}
