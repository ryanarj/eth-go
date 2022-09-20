package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var ganancheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganancheURL)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	block, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("blockNumber ", block)

	addrString := "0x222"

	addrHex := common.HexToAddress(addrString)

	balance, err := client.BalanceAt(context.Background(), addrHex, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance", balance)

}
