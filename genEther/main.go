package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://kovan.infura.io/v3/a197fe7fa0684b3b8ad84bf01fd2da89"
)

func main() {

	// "609ce3501f28c39d81f78c3063b6ba5d2b790333"
	// "bd97a16556e42f2214a95ca036a6f85cf64ee904"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	a1 := common.HexToAddress("609ce3501f28c39d81f78c3063b6ba5d2b790333")
	a2 := common.HexToAddress("bd97a16556e42f2214a95ca036a6f85cf64ee904")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balnace 1: ", b1)
	fmt.Println("Balnace 2: ", b2)

}
