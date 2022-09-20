package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	pvk, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	private_data := crypto.FromECDSA(pvk)
	fmt.Println("Private: ", hexutil.Encode(private_data))

	public_data := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println("Public: ", hexutil.Encode(public_data))

	fmt.Println("Address: ", crypto.PubkeyToAddress(pvk.PublicKey).Hex())

}
