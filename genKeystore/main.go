package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generate(password string) {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	passord := password
	account, err := key.NewAccount(passord)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(account.Address)
}

func main() {

	password := "password"
	generate(password)
	b, err := ioutil.ReadFile("./wallet/UTC--2022-09-22T19-07-47.626847000Z--e09a37bbc5527af01f10947f8bdf5d9065dd244c")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, password)

	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Private ", hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Public ", hexutil.Encode(pubData))

	fmt.Println("Add ", crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex())

}
