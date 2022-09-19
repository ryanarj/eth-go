package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var gananchURL = "http://localhost:8545"
var infuraURL = ""

func main() {
	client, err := ethclient.DialContext(context.Background(), gananchURL)
	if err != nil {
		log.Fatalf("Error client: %v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil)

	if err != nil {
		log.Fatalf("Error block: %v", err)
	}

	fmt.Println("block number", block.Number())

	addr := ""
	address := common.HexToAddress(addr)

	bal, err := client.BalanceAt(context.Background(), address, nil)

	if err != nil {
		log.Fatalf("Error bal: %v", err)
	}

	fmt.Println("the balc", bal)

	floatBal := new(big.Float)
	floatBal.SetString(bal.String())

	fmt.Println(floatBal)

	value := new(big.Float).Quo(floatBal, big.NewFloat(math.Pow10(18)))

	fmt.Println("Ether balance", value)
}
