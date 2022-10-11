package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func generate(password string) string {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	passord := password
	account, err := key.NewAccount(passord)
	if err != nil {
		log.Fatal(err)
	}
	return account.URL.Path
}

func main() {

	password := "password"
	path := generate(password)

	b, err := ioutil.ReadFile(path)
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
